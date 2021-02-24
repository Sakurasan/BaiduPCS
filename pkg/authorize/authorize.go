package authorize

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

var (
	// apikey          = "7Rkx8ubzsGPdZR7AE37u87dEfVM9TIxU"
	// secretkey       = "lrGxmFG63lfgQZrWzol5c0LadQIrkXaT"

	apikey    = "0coygXKzfFnZHafXfNj17zhj88H2ibN5"
	secretkey = "B665xWSbKbIuUGVPhPPwMDT2Fqar1WY4"

	authorizeUrl    = "https://openapi.baidu.com/oauth/2.0/authorize"
	access_tokenUrl = "https://openapi.baidu.com/oauth/2.0/token?"

	redirect_uri = "http://34c6772eb9e6.ngrok.io/pcs/"
)

func getAuthCode() (string, error) {
	// client_id := "7Rkx8ubzsGPdZR7AE37u87dEfVM9TIxU"
	// Secretkey := "lrGxmFG63lfgQZrWzol5c0LadQIrkXaT"
	response_type := "token"

	v := url.Values{}
	v.Add("grant_type", "authorization_code")
	v.Add("client_id", apikey)
	v.Add("response_type", response_type)
	v.Add("redirect_uri", redirect_uri)
	v.Add("scope", "basic,netdisk")
	v.Add("state", "STATE")
	http.DefaultClient.Get(authorizeUrl)

	req, _ := http.NewRequest(http.MethodGet, authorizeUrl, strings.NewReader(v.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	rspByte, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return "", err
	}
	fmt.Println(string(rspByte))

	return "", nil
}

func GetAccessTokenWithCode(code string) (string, error) {
	var (
		grant_type = "authorization_code"
		// code          = code
		client_id     = apikey
		client_secret = secretkey
		redirect_uri  = redirect_uri
	)
	v := url.Values{}
	v.Add("grant_type", grant_type)
	v.Add("code", code)
	v.Add("client_id", client_id)
	v.Add("client_secret", client_secret)
	v.Add("redirect_uri", redirect_uri)

	req, err := http.NewRequest("GET", access_tokenUrl+v.Encode(), nil)
	if err != nil {
		return "", err
	}
	// req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	bytereq, _ := httputil.DumpRequest(req, true)
	fmt.Println(string(bytereq))
	fmt.Println(req.URL.String(), req.URL.Query())
	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	rspbyte, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return "", err
	}
	// fmt.Println(string(rspbyte))
	return string(rspbyte), nil

}

func refreshAccessToken() string {
	return ""
}

// http://openapi.baidu.com/oauth/2.0/authorize?response_type=code&client_id=0coygXKzfFnZHafXfNj17zhj88H2ibN5&redirect_uri=http://ed18d21d1087.ngrok.io/pcs/&scope=basic,netdisk&display=tv&qrcode=1&force_login=1

// https://openapi.baidu.com/oauth/2.0/token?grant_type=authorization_code&code=CODE&client_id=0coygXKzfFnZHafXfNj17zhj88H2ibN5&client_secret=B665xWSbKbIuUGVPhPPwMDT2Fqar1WY4&redirect_uri=oob

// https://openapi.baidu.com/oauth/2.0/token?grant_type=authorization_code&code=a71be3659432dbecba00c094b5693e36&client_id=0coygXKzfFnZHafXfNj17zhj88H2ibN5&client_secret=B665xWSbKbIuUGVPhPPwMDT2Fqar1WY4&redirect_uri=http://ed18d21d1087.ngrok.io/pcs/
