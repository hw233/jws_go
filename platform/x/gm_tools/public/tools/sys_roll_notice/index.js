// Generated by CoffeeScript 1.12.7
(function() {
  var AccountKeyInput, Api, App, Checkbox, CheckboxGroup, NewSendModal, NoticeTable, React, ReactLoadingState, State, antd;

  antd = require('antd');

  Api = require('../api/api_ajax');

  React = require('react');

  AccountKeyInput = require('../../common/account_input');

  State = require('../../common/state');

  ReactLoadingState = State.ReactLoadingState;

  NoticeTable = require('./notice_table');

  NewSendModal = require('./new_sys_roll_notice');

  Checkbox = antd.Checkbox;

  CheckboxGroup = Checkbox.Group;

  App = React.createClass({displayName: "App",
    getInitialState: function() {
      return {
        select_server: this.props.curr_server,
        player_to_send: "",
        vershard: [],
        selected_sid: [],
        loads: new ReactLoadingState(this)
      };
    },
    LoadNotices: function() {
      if (this.refs.notice_table != null) {
        this.state.loads.EnterLoading("load");
        return this.refs.notice_table.Refersh(1, (function(_this) {
          return function() {
            return _this.state.loads.LeaveLoading("load");
          };
        })(this));
      }
    },
    isServerRight: function() {
      return this.state.selected_sid.length > 0;
    },
    getBntState: function() {
      if (!this.isServerRight()) {
        return "ant-btn ant-btn-primary disabled";
      }
      return "ant-btn ant-btn-primary " + this.state.loads.GetBtnStat("load");
    },
    getTable: function() {
      if (this.state.loadstat === 0) {
        return React.createElement("div", null, "点击按钮加载ID");
      }
      if (!this.isServerRight()) {
        return React.createElement("div", null, " ");
      }
      return React.createElement(NoticeTable, Object.assign({}, this.props, {
        "ref": "notice_table",
        "server_id": this.state.selected_sid,
        "account_id": this.state.player_to_send
      }));
    },
    getSendModal: function() {
      if (!this.isServerRight()) {
        return React.createElement("div", null, "请选择服务器");
      }
      return React.createElement(NewSendModal.NewModalWithButton, Object.assign({}, this.props, {
        "ref": "notice_send",
        "modal_name": "添加跑马灯公告",
        "server_id": this.state.selected_sid,
        "account_id": "ac",
        "on_loaded": this.LoadNotices
      }));
    },
    handleUserChange: function(data) {
      if (data === "") {
        data = "请输入玩家Id";
      }
      return this.setState({
        player_to_send: data
      });
    },
    selectSid: function(e) {
      var i, len, ss, value;
      ss = [];
      for (i = 0, len = e.length; i < len; i++) {
        value = e[i];
        ss.push(value);
      }
      console.log("selectSid");
      console.log(ss);
      return this.setState({
        selected_sid: ss
      });
    },
    genGroupCheckBox: function(ar) {
      var a, i, len, plainOptions;
      plainOptions = [];
      for (i = 0, len = ar.length; i < len; i++) {
        a = ar[i];
        plainOptions.push(a);
      }
      return React.createElement("tr", null, React.createElement("td", null, React.createElement(CheckboxGroup, {
        "options": plainOptions,
        "value": this.state.selected_sid,
        "onChange": this.selectSid
      })));
    },
    gencheckbox: function() {
      var column_count, count, res, selected;
      column_count = 5;
      res = [];
      selected = [];
      count = 0;
      res.push(this.genGroupCheckBox(this.state.vershard));
      return res;
    },
    selectallcheckbox: function(e) {
      if (e.target.checked > 0) {
        return this.selectSid(this.state.vershard);
      } else {
        return this.setState({
          selected_sid: []
        });
      }
    },
    handleSubmit_ver: function() {
      var api;
      api = new Api();
      return api.Typ("getAllServer").ServerID("").AccountID("").Key(this.props.curr_key).ParamArray([]).Do((function(_this) {
        return function(result) {
          var res;
          res = JSON.parse(result);
          return _this.setState({
            vershard: res["Infos"]
          });
        };
      })(this));
    },
    render: function() {
      return React.createElement("div", null, React.createElement("div", {
        "className": "row-flex row-flex-middle row-flex-start"
      }, React.createElement(antd.Button, {
        "onClick": this.handleSubmit_ver
      }, "获取所有服务器")), React.createElement("label", null, React.createElement(Checkbox, {
        "onChange": this.selectallcheckbox
      }), "全选"), React.createElement("div", {
        "style": {
          marginBottom: 24
        }
      }, React.createElement("br", null)), React.createElement("div", null, this.gencheckbox()), React.createElement("div", {
        "style": {
          marginBottom: 24
        }
      }, React.createElement("br", null)), React.createElement("div", {
        "className": "row-flex row-flex-middle row-flex-start"
      }, React.createElement(antd.Button, {
        "className": this.getBntState(),
        "onClick": this.LoadNotices
      }, "加载跑马灯公告"), this.getSendModal()), React.createElement("hr", null), this.getTable());
    }
  });

  module.exports = App;

}).call(this);