package basic

import (
	"BaiduPCS/tests"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"testing"
	"time"
)

func TestFilemetas(t *testing.T) {
	type args struct {
		t *FilemetasReqType
	}
	var reqType = &FilemetasReqType{
		Access_token: tests.PCS.Access_token,
		Fsids:        []uint64{1113188906626601, 214582380365016},
	}

	ts := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "Filemetas",
			args: args{t: reqType},
		},
	}
	for _, tt := range ts {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Filemetas(tt.args.t)
			if err != nil {
				t.Errorf("Filemetas() error = %v", err)
				return
			}
			// b, _ := json.Marshal(got)
			// t.Logf("Filemetas() = %v", string(b))
			// t.Log(got.List)
			for _, v := range got.List {
				if v.IsDir == 1 {
					return
				}
				log.Printf("%s \n %s \n", v.Path, v.DLink+"&access_token="+tests.PCS.Access_token)
				log.Println(v.FsID)``
				request(v.DLink + "&access_token=" + tests.PCS.Access_token)
			}
		})
	}
}

// https://d.pcs.baidu.com/file/1261d72d03471f7b7b805fd60e024b8d?
// fid=2082810368-250528-414244021542671&
// rt=pr&
// sign=FDtAERV-DCb740ccc5511e5e8fedcff06b081203-Rnos3iOhNnMF1pS44AUWwor%2BJw8%3D&
// expires=8h
// &chkbd=0
// &chkv=2
// &dp-logid=4111511902857508725&
// dp-callid=0&
// dstime=1596179809&
// r=802381259
func request(dlink string) {
	// curl -L -X GET 'https://d.pcs.baidu.com/file/e521f566d01d7b85203b33bc55013bbd?fid=2987416974-250528-214582380365016&rt=pr&sign=FDtAERV-DCb740ccc5511e5e8fedcff06b081203-chhfDZ30jP5nT8%2FxXmYGunyatm0%3D&expires=8h&chkbd=0&chkv=2&dp-logid=958503253314344520&dp-callid=0&dstime=1614101243&r=349171534&origin_appid=23674228&file_type=0&access_token=121.9e1b535796673c3e6cdcdf60b5df92e4.Ynl-AOYVtTV2ULUa9DSr4SU2wNILH7HoU2tFDsO.DO8SkA'
	// -H 'User-Agent: pan.baidu.com'
	req, _ := http.NewRequest("GET", dlink, nil)
	req.Header.Set("User-Agent", "pan.baidu.com")

	c := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
		Timeout: 30 * time.Second,
	}
	rsp, err := c.Do(req)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(rsp.Header.Get("Location"))

}

func ariac()  {
	// /Volumes/Mac/DownLoad/TD/
	// aria2c -c -s10 -k1M -x16 --enable-rpc=false -d "/Volumes/Mac/DownLoad/TD/" -o '要保存的文件名' --header "User-Agent: Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36" --header "Cookie: 你的cookie内容" "文件的下载地址"
}