package googleClient

import (
	"testing"
	"fmt"
	"io/ioutil"
	"net/url"
	"go-client/client/authorizationProvider"
)

func TestGetGoogle(t *testing.T) {
	googleAddress, _ := url.Parse("https://www.google.com")
	token := "eyJhbGciOiJSUzI1NiJ9.eyJleHBpcnlfdGltZSI6MTQ3NzQ3NzEyMDMwOSwidXNlcl90eXBlIjoidXNlciIsInJlYWxtIjoiMWViNjVmZGYtOTY0My00MTdmLTk5NzQtYWQ3MmNhZTBlMTBmIiwiY2lzX3V1aWQiOiI5YjQyYjU1Yy1mOWNmLTQxZmQtOTY2ZS01NWFjMmI0NTg1MjkiLCJyZWZlcmVuY2VfaWQiOiJhN2Q3YTUzYS1lMWJjLTQ5MzYtOTE2Mi04OTU0Nzg1ODA1OGUiLCJpc3MiOiJodHRwczpcL1wvaWRicm9rZXIud2ViZXguY29tXC9pZGIiLCJ0b2tlbl90eXBlIjoiQmVhcmVyIiwiY2xpZW50X2lkIjoiQzk2ZDM4OWQ2MzJjOTZkMDM4ZDhmNDA0YzM1OTA0YjUxMDg5ODhiZDZkNjAxZDRiNDdmNGVlYzg4YTU2OWQ1ZGIiLCJ0b2tlbl9pZCI6Ik1HVTBNek01WldFdFpEZ3laaTAwT0RabExXRXdORFF0TURsaFl6SXpNekJrWTJaa01tRXhaVE5qT1RRdFpqWXoiLCJwcml2YXRlIjoiZXlKaGJHY2lPaUprYVhJaUxDSmpkSGtpT2lKS1YxUWlMQ0psYm1NaU9pSkJNVEk0UTBKRExVaFRNalUySW4wLi5DUVFTM2Y5Z3MzS0VnaGRJYTg3Mnl3LkV5Rmd4TEZiMFRFUGlCTlM2d0dpbnJKZEJFTzhpeEV0aTZrZG82V1pvdjg4SmZHdnc0ODZqZjU1TE02WUlUVmpFbndKZGVLMVhkWWFxVzAzY2NXaHJqZGJ4MmFoQVZuVTU2ZEVHTEFQa2JWbndZWnItODRMY0ZtSVlhM01aY09yQ1NIS1BpTmozWF8xYmJ1Q1NaM1JMZkNOZ3pBSnp4TjF5R1laOW1VekxnZVdfLU8xT0NrOXZNTWpRemJfNzRiV180aXVab3BVTDQ1TkdabHJ0MUlyZEV4WERqU09hX0d5YmJzcFFlVW5OWk0yYkgzYlU1YTRhQU5UVllkVjdyV3lfVklDY09uMFVTWG96VEhQYmdoTWxXeVRidG40a1ZnUHc0SVFSb2p2U3dxcC1fQjNMc3QyQVVxSDBHTVVyWEZWaUJfUll6WTZIajk1Y2xmUzM4SmJNaFNXRHl2dzR1aTB5V2tPQUk3Vl9DaUp0ZHVkMzRwdU5UdmJuT2xHVHZpTVJ4bVJKblhmLUl3NEZfc18yV0JqbGhPZmp5U0Y5WFlEbmNGNnlHZzBBa2djQ1RXY0FUV3FmRHFsTXRrdlhYOUt2bGUwS2NSNUFlYUJSRVotODJVNThMX29BMnZYalE2bFhJWFFXUnZ5Z0tBZDVGcTg3V0ZtSDlKSlB0WUR1UFNPMVh1cE1vU19CV2cwMF9XOWJoYWFrdy5kTkhPd1g4QVlVTlhnODNJS1p3WDhRIiwidXNlcl9tb2RpZnlfdGltZXN0YW1wIjoiMjAxNjA5MjcxNzA4MzUuNjY0WiJ9.iBGqcjEnm6rhhZ0NaD9-IL0CHRDUmbUEYMX1UsUcYThUg4qEGQDZ5tQClBziYgDGRiOKPBfCqBP31iQu04y090IUy_l8FQNOQmsYMxAWLEHHbilEheKyzqTdhZS7YGc-3bAYNUtZOdLUgJW7OBuRDCB40yBCxgR-zvdmuSDdNJod-bpccwtt1hQ_vvv7zh3JnLveYZWTlILfKNTR-yKNVC7mFJho71imExN_uCP-8PNpzULE_sxVcoY9det-lehxfDYiEVrCVzqKFJVbm6PsY5_eoxnefWleeUFS0VH5IkkEIkq7JsMLeS7yMs9-M7R1nZdN0LubFCTOfWbTzjk9gQ"

	builder := ((*googleClientFactoryBuilder)(nil)).NewBuilder()
	builder.BaseUrl = googleAddress
	simpleProvider := authorizationProvider.SimpleBearerAuthorizationProvider{Token:token}
	builder.AuthorizationProvider = simpleProvider
	response, err := builder.Builder().newGoogleClient().GetGoogle()
	defer response.Body.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}