package oauth

import (
	"github.com/alvaro259818/golang-restclient/rest"
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"testing"
)

func TestMain(m *testing.M)  {
	rest.StartMockupServer()
	os.Exit(m.Run())
}

func TestOauthConstants(t *testing.T) {
	assert.EqualValues(t, "X-Public", headerXPublic)
	assert.EqualValues(t, "X-Client-Id", headerXClientId)
	assert.EqualValues(t, "X-Caller-Id",headerXCallerId)
	assert.EqualValues(t, "access_token", paramAccessToken)
}

func TestIsPublicNilRequest(t *testing.T) {
	assert.True(t, IsPublic(nil))
}

func TestIsPublicNoError(t *testing.T) {
	request := http.Request{
		Header: make(http.Header),
	}
	assert.False(t, IsPublic(&request))

	request.Header.Add("X-Public", "true")
	assert.True(t, IsPublic(&request))

}

func TestCallerIdNilRequest(t *testing.T)  {

}

func TestCallerInvalidCallerFormat(t *testing.T)  {

}

func TestGetCallerNoError(t *testing.T) {

}

func TestGetAccessTokenInvalidRestClientResponse(t *testing.T)  {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod: http.MethodGet,
		URL: "http://localhost:8080/oauth/access_token/ABC123",
		ReqBody: "",
		RespHTTPCode: -1,
		RespBody: "{}",
	})
	accessToken, err := getAccessToken("ABC123")
	assert.Nil(t, accessToken)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status())
	assert.EqualValues(t, "invalid restclient response when trying to get access token", err.Message())
}