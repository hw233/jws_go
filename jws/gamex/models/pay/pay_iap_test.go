package pay

import (
	"testing"

	"github.com/dogenzaka/go-iap/appstore"
	"github.com/stretchr/testify/assert"
	"vcs.taiyouxi.net/platform/planx/util/logs"
)

func TestIAP(t *testing.T) {

	client := appstore.New()
	req := appstore.IAPRequest{
		ReceiptData: "ewoJInNpZ25hdHVyZSIgPSAiQWtRV1liczR1andCOTFIUjNndkJEMEhEMWhZZU41ZXVMUTZzd3ZsMGlhSWhpL0JSalJ2bmJhcVFnejNqdndYL0tFWnRrV01XNlJhSGl3dFY4Q1pqUzlob1g3VitJbUg2elR4Ty9RZDFpSk01U0RUTm5XYXpXR25FMis2SmxtUGxqL2taNWlOaUpHWlorWHpkTkIrR2w5WVpoNjlKSGhWZUdVSFZ5VFFEZ3UyOEFBQURWekNDQTFNd2dnSTdvQU1DQVFJQ0NCdXA0K1BBaG0vTE1BMEdDU3FHU0liM0RRRUJCUVVBTUg4eEN6QUpCZ05WQkFZVEFsVlRNUk13RVFZRFZRUUtEQXBCY0hCc1pTQkpibU11TVNZd0pBWURWUVFMREIxQmNIQnNaU0JEWlhKMGFXWnBZMkYwYVc5dUlFRjFkR2h2Y21sMGVURXpNREVHQTFVRUF3d3FRWEJ3YkdVZ2FWUjFibVZ6SUZOMGIzSmxJRU5sY25ScFptbGpZWFJwYjI0Z1FYVjBhRzl5YVhSNU1CNFhEVEUwTURZd056QXdNREl5TVZvWERURTJNRFV4T0RFNE16RXpNRm93WkRFak1DRUdBMVVFQXd3YVVIVnlZMmhoYzJWU1pXTmxhWEIwUTJWeWRHbG1hV05oZEdVeEd6QVpCZ05WQkFzTUVrRndjR3hsSUdsVWRXNWxjeUJUZEc5eVpURVRNQkVHQTFVRUNnd0tRWEJ3YkdVZ1NXNWpMakVMTUFrR0ExVUVCaE1DVlZNd2daOHdEUVlKS29aSWh2Y05BUUVCQlFBRGdZMEFNSUdKQW9HQkFNbVRFdUxnamltTHdSSnh5MW9FZjBlc1VORFZFSWU2d0Rzbm5hbDE0aE5CdDF2MTk1WDZuOTNZTzdnaTNvclBTdXg5RDU1NFNrTXArU2F5Zzg0bFRjMzYyVXRtWUxwV25iMzRucXlHeDlLQlZUeTVPR1Y0bGpFMU93QytvVG5STStRTFJDbWVOeE1iUFpoUzQ3VCtlWnRERWhWQjl1c2szK0pNMkNvZ2Z3bzdBZ01CQUFHamNqQndNQjBHQTFVZERnUVdCQlNKYUVlTnVxOURmNlpmTjY4RmUrSTJ1MjJzc0RBTUJnTlZIUk1CQWY4RUFqQUFNQjhHQTFVZEl3UVlNQmFBRkRZZDZPS2RndElCR0xVeWF3N1hRd3VSV0VNNk1BNEdBMVVkRHdFQi93UUVBd0lIZ0RBUUJnb3Foa2lHOTJOa0JnVUJCQUlGQURBTkJna3Foa2lHOXcwQkFRVUZBQU9DQVFFQWVhSlYyVTUxcnhmY3FBQWU1QzIvZkVXOEtVbDRpTzRsTXV0YTdONlh6UDFwWkl6MU5ra0N0SUl3ZXlOajVVUllISytIalJLU1U5UkxndU5sMG5rZnhxT2JpTWNrd1J1ZEtTcTY5Tkluclp5Q0Q2NlI0Szc3bmI5bE1UQUJTU1lsc0t0OG9OdGxoZ1IvMWtqU1NSUWNIa3RzRGNTaVFHS01ka1NscDRBeVhmN3ZuSFBCZTR5Q3dZVjJQcFNOMDRrYm9pSjNwQmx4c0d3Vi9abEwyNk0ydWVZSEtZQ3VYaGRxRnd4VmdtNTJoM29lSk9PdC92WTRFY1FxN2VxSG02bTAzWjliN1BSellNMktHWEhEbU9Nazd2RHBlTVZsTERQU0dZejErVTNzRHhKemViU3BiYUptVDdpbXpVS2ZnZ0VZN3h4ZjRjemZIMHlqNXdOelNHVE92UT09IjsKCSJwdXJjaGFzZS1pbmZvIiA9ICJld29KSW05eWFXZHBibUZzTFhCMWNtTm9ZWE5sTFdSaGRHVXRjSE4wSWlBOUlDSXlNREUyTFRBeExUQTBJREF3T2pBNE9qRTFJRUZ0WlhKcFkyRXZURzl6WDBGdVoyVnNaWE1pT3dvSkluVnVhWEYxWlMxcFpHVnVkR2xtYVdWeUlpQTlJQ0kwTURneE1EZzBPVFkzT1dGbFlUbGpZbVZqWkRoak5ETmtabVpsTUdSaU1HRTJOalJqT0dVeUlqc0tDU0p2Y21sbmFXNWhiQzEwY21GdWMyRmpkR2x2YmkxcFpDSWdQU0FpTVRBd01EQXdNREU0TnpJM016ZzRNeUk3Q2draVluWnljeUlnUFNBaU56SXlJanNLQ1NKMGNtRnVjMkZqZEdsdmJpMXBaQ0lnUFNBaU1UQXdNREF3TURFNE56STNNemc0TXlJN0Nna2ljWFZoYm5ScGRIa2lJRDBnSWpFaU93b0pJbTl5YVdkcGJtRnNMWEIxY21Ob1lYTmxMV1JoZEdVdGJYTWlJRDBnSWpFME5URTRPVFE0T1RVM056a2lPd29KSW5WdWFYRjFaUzEyWlc1a2IzSXRhV1JsYm5ScFptbGxjaUlnUFNBaVJESXhPRGt5TUVVdFJFUkZNQzAwUWpsQ0xVSXpSa0V0TmpFNE9ERkROa0l3UVVKRklqc0tDU0p3Y205a2RXTjBMV2xrSWlBOUlDSmpiMjB1ZEdGcGVXOTFlR2t1YVdaelp5NHhNQ0k3Q2draWFYUmxiUzFwWkNJZ1BTQWlNVEEyTVRrME5qWTFNaUk3Q2draVltbGtJaUE5SUNKamIyMHVkR0ZwZVc5MWVHa3VhV1p6WnlJN0Nna2ljSFZ5WTJoaGMyVXRaR0YwWlMxdGN5SWdQU0FpTVRRMU1UZzVORGc1TlRjM09TSTdDZ2tpY0hWeVkyaGhjMlV0WkdGMFpTSWdQU0FpTWpBeE5pMHdNUzB3TkNBd09Eb3dPRG94TlNCRmRHTXZSMDFVSWpzS0NTSndkWEpqYUdGelpTMWtZWFJsTFhCemRDSWdQU0FpTWpBeE5pMHdNUzB3TkNBd01Eb3dPRG94TlNCQmJXVnlhV05oTDB4dmMxOUJibWRsYkdWeklqc0tDU0p2Y21sbmFXNWhiQzF3ZFhKamFHRnpaUzFrWVhSbElpQTlJQ0l5TURFMkxUQXhMVEEwSURBNE9qQTRPakUxSUVWMFl5OUhUVlFpT3dwOSI7CgkiZW52aXJvbm1lbnQiID0gIlNhbmRib3giOwoJInBvZCIgPSAiMTAwIjsKCSJzaWduaW5nLXN0YXR1cyIgPSAiMCI7Cn0=",
	}
	resp, err := client.Verify(req)

	if err != nil {
		logs.Error("err %s", err.Error())
	}

	assert.NotEmpty(t, resp.Receipt.OriginalPurchaseDate)

	/*
		logs.Error("resp[] %v", resp)
		logs.Error("resp[] %v", resp.Receipt.ReceiptType)
		logs.Error("resp[] %v", resp.Receipt.AppItemID)
	*/

	r, err := ParseReceiptData(req.ReceiptData)

	assert.Equal(t, "com.taiyouxi.ifsg.10", r.PurchaseInfo.ProductID)
	assert.Equal(t, "1061946652", r.PurchaseInfo.ItemID)
	assert.Equal(t, "D218920E-DDE0-4B9B-B3FA-61881C6B0ABE", r.PurchaseInfo.UniqueVendorIdentifier)
	assert.Equal(t, "2016-01-04 08:08:15 Etc/GMT", r.PurchaseInfo.OriginalPurchaseDate)
	assert.Equal(t, "Sandbox", r.Environment)

	logs.Close()
}
