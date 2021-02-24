package basic

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

var (
	b  = 1
	kb = 1024 * b
	mb = 1024 * kb
	gb = 1024 * mb
	tb = 1024 * gb
)

func GetQuotaInfo(access_token string) (string, error) {
	var (
		_url = "https://pan.baidu.com/api/quota"
	)
	v := url.Values{}
	v.Add("access_token", access_token)
	req, _ := http.NewRequest("GET", _url+"?"+v.Encode(), nil)
	req.Header.Set("User-Agent", "pan.baidu.com'")

	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	bytersp, err := ioutil.ReadAll(rsp.Body)
	return string(bytersp), nil

}
