package basic

import (
	"encoding/json"
	"net/http"
	"net/url"
)

type FilemetasReqType struct {
	Access_token string   //是	---	query	接口鉴权参数
	Fsids        []uint64 //是	---	-----	文件id数组，数组中元素是uint64类型，数组大小上限是：100
	Path         string   //否	---	-----	查询共享目录时需要，格式： /uk-fsid 其中uk为共享目录创建者id， fsid对应共享目录的fsid
	Thumb        string   //否	---	-----	是否需要缩略图地址，0为否，1为是，默认为0
	Dlink        string   //否	---	-----	是否需要下载地址，0为否，1为是，默认为0
	Extra        string   //否	---	-----	图片是否需要拍摄时间、原图分辨率等其他信息，0 否、1 是，默认0
}

type FilemetasRspType struct {
	ErrorCode    int    `json:"errno"`
	ErrorMsg     string `json:"errmsg"`
	RequestID    int
	RequestIDStr string `json:"request_id"`
	List         []struct {
		FsID        uint64            `json:"fs_id"`
		Path        string            `json:"path"`
		Category    int               `json:"category"`
		FileName    string            `json:"filename"`
		IsDir       int               `json:"isdir"`
		Size        int               `json:"size"`
		Md5         string            `json:"md5"`
		DLink       string            `json:"dlink"`
		Thumbs      map[string]string `json:"thumbs"`
		ServerCtime int               `json:"server_ctime"`
		ServerMtime int               `json:"server_mtime"`
		DateTaken   int               `json:"date_taken"`
		Width       int               `json:"width"`
		Height      int               `json:"height"`
	}
}

func Filemetas(t *FilemetasReqType) (*FilemetasRspType, error) {
	var (
		_url = "https://pan.baidu.com/rest/2.0/xpan/multimedia"
	)
	fsids, err := json.Marshal(t.Fsids)
	if err != nil {
		return nil, err
	}
	v := url.Values{}
	v.Add("method", "filemetas")

	v.Add("access_token", t.Access_token)
	v.Add("fsids", string(fsids))
	v.Add("dlink", defaultValue(t.Dlink, "1"))
	v.Add("thumb", defaultValue(t.Thumb, "1"))
	v.Add("extra", defaultValue(t.Extra, "1"))

	req, err := http.NewRequest("GET", _url+"?"+v.Encode(), nil)
	if err != nil {
		return nil, err
	}

	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	// br, _ := ioutil.ReadAll(rsp.Body)
	// fmt.Println(string(br))

	rspType := new(FilemetasRspType)
	if err := json.NewDecoder(rsp.Body).Decode(rspType); err != nil {
		return nil, err
	}
	return rspType, nil

}
