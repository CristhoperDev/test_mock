package test

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/CristhoperDev/test_wrapper/client"
	"github.com/CristhoperDev/test_wrapper/mock"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func init() {
	client.Client = &mock.MockClient{}
	viper.Set("VISANET_DEV_AUTHORIZATION_API", "https://apitestenv.vnforapps.com/api.security/v1/security")
	viper.Set("VISANET_DEV_SESSION_API", "https://apitestenv.vnforapps.com/api.ecommerce/v2/ecommerce/token/session/")
	viper.Set("VISANET_DEV_USER", "integraciones.visanet@necomplus.com")
	viper.Set("VISANET_DEV_MERCHANT_ID", "522591303")
	viper.Set("VISANET_DEV_PASSWD", "d5e7nk$M")
}


func TestCreateRepoSuccess(t *testing.T) {
	// build response JSON
	json := `eyJraWQiOiJmWk1tV3pZR0RBckxHektvalNCK2w3SjFhMnNPXC9zQnNwOTlNNmNuM3F5MD0iLCJhbGciOiJSUzI1NiJ9.eyJzdWIiOiJjNTAxNjYyOS04Zjc2LTQ1M2QtYjhlNC01MGJjZDI5YjI2NTAiLCJldmVudF9pZCI6IjYyYTM2MjBmLWVlNzItNDdkMy1hODdjLWRmMmQ3ZjI0YWU0MSIsInRva2VuX3VzZSI6ImFjY2VzcyIsInNjb3BlIjoiYXdzLmNvZ25pdG8uc2lnbmluLnVzZXIuYWRtaW4iLCJhdXRoX3RpbWUiOjE1OTQ4MzQ3MjksImlzcyI6Imh0dHBzOlwvXC9jb2duaXRvLWlkcC51cy1lYXN0LTEuYW1hem9uYXdzLmNvbVwvdXMtZWFzdC0xXzJjSjFTZTFmSSIsImV4cCI6MTU5NDgzODMyOSwiaWF0IjoxNTk0ODM0NzI5LCJqdGkiOiIyZDA2ZDU5Ni0xMGRmLTRmYzQtODA1Yi0zZjEyNzk4NTRlNGIiLCJjbGllbnRfaWQiOiIxMGx2MDYxN281ZGljNTFlYnNucWVpaWpiNyIsInVzZXJuYW1lIjoiaW50ZWdyYWNpb25lcy52aXNhbmV0QG5lY29tcGx1cy5jb20ifQ.HKj00LRJVtUv3kUHu83JWCTdHjMcAtzzJEMEX7aXNkr-0cQ5d3ML_RLn5bqhK44S8VKCRBUZzY-eCiBllXVPicTxdmhHIg4GbkQwpKGhHlhGpkQpRKsNmTO1xQ3IkaSHKEkl1GdngPdtet0rYHTefy16xJXrluREizpNepI-BYY4-KVVcdZpDIbNg0r5xiXlzaQ4dPagNfqT6XpeJz3dHcRhQ74NKOtl0HqkMZAWx_Qj6zZjddqQXi9-HcLy9Q3FG3PghlQCv-qNHKki0FpT1nVH6FMeQtkGNSa5AJ_SSqbohrWs3-ZtK-AijqyVQqRPRvkCfRWcKEEuj-qOgyag3w`
	//json := `Unauthorized access`
	// create a new reader with that JSON
	r := ioutil.NopCloser(bytes.NewReader([]byte(json)))
	mock.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 401,
			Body:       r,
		}, nil
	}

	user := viper.Get("VISANET_DEV_USER")
	pwd := viper.Get("VISANET_DEV_PASSWD")
	encodedCredential := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", user, pwd)))
	fmt.Println(encodedCredential)

	header := http.Header{
		"Authorization": {
			"Basic " + encodedCredential,
		},
	}

	response, err := client.Post(viper.GetString("VISANET_DEV_AUTHORIZATION_API"), nil, header)

	body, err := ioutil.ReadAll(response.Body)


	assert.Nil(t, err)
	assert.EqualValues(t, json,
			string(body))
}