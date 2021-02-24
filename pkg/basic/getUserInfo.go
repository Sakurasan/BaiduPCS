package basic

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func GetUserInfo(access_token string) (string, error) {
	var (
		_url = "https://pan.baidu.com/rest/2.0/xpan/nas"
	)
	v := url.Values{}
	v.Add("method", "uinfo")
	v.Add("access_token", access_token)
	v.Add("get_unionid", "1")

	req, _ := http.NewRequest("GET", _url+"?"+v.Encode(), nil)
	req.Header.Set("User-Agent", "pan.baidu.com")
	bytereq, _ := httputil.DumpRequest(req, true)
	fmt.Println(string(bytereq))
	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	bytersp, _ := ioutil.ReadAll(rsp.Body)
	return string(bytersp), nil
}
