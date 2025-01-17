package storehelper

import (
	"errors"
	"strings"
	"sync"
	"time"

	"github.com/cenk/backoff"
	"vcs.taiyouxi.net/platform/planx/util/dynamodb"
	"vcs.taiyouxi.net/platform/planx/util/logs"
)

type StoreDynamoDB struct {
	db        *dynamodb.DynamoDB
	region    string
	db_name   string
	accessKey string
	secretKey string

	count_all int
	count_err int

	time_log time.Time

	format string
	seq    string
}

func (s *StoreDynamoDB) LogInfo() {
	n := time.Now()
	in := n.Sub(s.time_log)

	logs.Info("put info %d / %d by %d",
		s.count_err, s.count_all, in.Seconds())
	s.count_all = 0
	s.count_err = 0
	s.time_log = n
}

func NewStoreDynamoDB(region, db_name, accessKey, secretKey string, format, seq string) *StoreDynamoDB {
	db := &dynamodb.DynamoDB{}

	return &StoreDynamoDB{
		db,
		region,
		db_name,
		accessKey,
		secretKey,
		0, 0, time.Now(),
		format, seq,
	}
}

func (s *StoreDynamoDB) Clone() (IStore, error) {
	db := &dynamodb.DynamoDB{}
	n := &StoreDynamoDB{
		db,
		s.region,
		s.db_name,
		s.accessKey,
		s.secretKey,
		0, 0, time.Now(),
		s.format, s.seq,
	}

	if err := s.Open(); err != nil {
		return nil, err
	}
	return n, nil
}

func (s *StoreDynamoDB) Open() error {
	err := s.db.Connect(
		s.region,
		s.accessKey,
		s.secretKey,
		"")
	if err != nil {
		return err
	}

	return s.db.InitTable()
}

func (s *StoreDynamoDB) Close() error {
	return nil
}

//TODO YZH: batch write optmization of DynamoDB
func (s *StoreDynamoDB) Put(key string, val []byte, rh ReadHandler) error {
	if s.count_all >= 100 {
		s.LogInfo()
	}

	err := s.db.SetByHash(s.db_name, key, string(val))

	s.count_all++
	if err == nil {
		return nil
	}
	s.count_err++

	//logs.Trace("err : %v", err.Error())

	b := backoff.NewExponentialBackOff()
	b.InitialInterval = 100 * time.Millisecond
	b.MaxElapsedTime = 1 * time.Minute
	// 以下取默认值
	// DefaultMaxInterval         = 60 * time.Second
	// DefaultMaxElapsedTime      = 15 * time.Minute
	ticker := backoff.NewTicker(b)
	defer ticker.Stop()

	for _ = range ticker.C {
		err = s.db.SetByHash(s.db_name, key, string(val))
		if err == nil {
			break
		}
		//logs.Warn("re put %s %v", key, k)
	}
	return err
}

func (s *StoreDynamoDB) Get(key string) ([]byte, error) {
	res, err := s.db.GetByHash(s.db_name, key)
	if err != nil {
		return []byte{}, err
	}
	logs.Trace("get dynamodb %v", res)

	r, ok := res.(string)
	if !ok {
		logs.Trace("res is no []byte")
		return []byte{}, errors.New("res is no []byte")
	}
	return []byte(r), err
}

func (s *StoreDynamoDB) Del(key string) error {
	err := s.db.DelKey(s.db_name, key)
	return err
}

func (s *StoreDynamoDB) DelKR(key, r string) error {
	err := s.db.DelKeyRange(s.db_name, key, r)
	return err
}

func (s *StoreDynamoDB) StoreKey(key string) string {
	now_time := time.Now()
	if s.format == "" {
		return key
	}
	return now_time.Format(s.format) + s.seq + key
}

func (s *StoreDynamoDB) RedisKey(key_in_store string) (string, bool) {
	//TODO YZH
	if s.seq == "" {
		return key_in_store, true
	}

	d := strings.Split(key_in_store, s.seq)
	if len(d) < 2 {
		logs.Error("key_in_store err : %s in %s", key_in_store, s.seq)
		return "", false
	}
	return d[len(d)-1], true
}

type KeyScanHander func(idx int, key, data string) error

func (s *StoreDynamoDB) Scan(hander KeyScanHander, scan_len, worker_num int64) error {
	return s.ScanKV(hander, scan_len, worker_num, true)
}
func (s *StoreDynamoDB) ScanKV(hander KeyScanHander, scan_len, worker_num int64, needV bool) error {
	b := backoff.NewExponentialBackOff()
	b.InitialInterval = 500 * time.Millisecond
	b.MaxElapsedTime = 5 * time.Minute

	//return s.db.Scan(s.db_name, scan_len, f, b)

	res_channel := make(chan dynamodb.DynamoKV)

	var wg_all sync.WaitGroup
	var wg sync.WaitGroup
	wg_all.Add(1)

	var i int64 = 0
	for ; i < worker_num; i++ {
		wg.Add(1)
		// 小心闭包的捕获变量的机制
		go func(idx int64) {
			//logs.Warn("scanner start %d", idx)
			defer wg.Done()
			err := s.db.ParallelScan(s.db_name,
				scan_len, int64(idx), int64(worker_num),
				res_channel, b)
			if err != nil {
				logs.Error("ParallelScan err: %v", err.Error())
			}
		}(i)
	}

	go func() {
		defer wg_all.Done()
		//logs.Warn("scan start")
		for {
			s, ok := <-res_channel
			//logs.Trace("new %v", s)
			if !ok {
				logs.Warn("RestoreDynamoDB queue close")
				return
			}
			k_s, kok := s.K.(string)
			v_s, vok := s.V.(string)

			if !kok {
				logs.Error("K type Err")
				continue
			}
			if needV && !vok {
				logs.Error("V type Err")
				continue
			}
			if !needV && vok {
				v_s = ""
			}
			err := hander(0, k_s, v_s)
			if err != nil {
				logs.Error("hander %s err: %s", k_s, err.Error())
			}
		}
	}()

	wg.Wait()
	close(res_channel)
	wg_all.Wait()

	return nil
}
