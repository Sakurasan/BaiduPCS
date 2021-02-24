package basic

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type FileListReqType struct {
	Access_token string //是	---	query	接口鉴权参数
	Dir          string //否	---	---	需要list的目录，以/开头的绝对路径, 默认为/
	// time表示先按文件类型排序，后按修改时间排序
	// name表示先按文件类型排序，后按文件名称排序
	// size表示先按文件类型排序， 后按文件大小排序
	Order     string //否	---	---	排序字段：默认为name
	Desc      string //否	---	---	1 该KEY存在为降序，否则为升序 （注：排序的对象是当前目录下所有文件，不是当前分页下的文件）
	Start     string //否	---	---	起始位置，从0开始
	Limit     string //否	---	---	每页条目数，默认为1000，最大值为10000
	Web       string //否			值为web时， 返回dir_empty属性 和 缩略图数据
	Folder    string //否			是否只返回文件夹，0 返回所有，1 只返回目录条目，且属性只返回path字段。
	Showempty string
}
type FileListRspType struct {
	Errno    int    `json:"errno,omitempty"`
	GUIDInfo string `json:"guid_info,omitempty"`
	List     []struct {
		ServerFilename string `json:"server_filename,omitempty"`
		Privacy        int    `json:"privacy,omitempty"`
		Category       int    `json:"category,omitempty"`
		Unlist         int    `json:"unlist,omitempty"`
		Pl             int    `json:"pl,omitempty"`
		FsID           int64  `json:"fs_id,omitempty"`
		DirEmpty       int    `json:"dir_empty,omitempty"`
		ServerAtime    int    `json:"server_atime,omitempty"`
		ServerCtime    int    `json:"server_ctime,omitempty"`
		LocalMtime     int    `json:"local_mtime,omitempty"`
		Wpfile         int    `json:"wpfile,omitempty"`
		Size           int    `json:"size,omitempty"`
		Isdir          int    `json:"isdir,omitempty"`
		Share          int    `json:"share,omitempty"`
		Path           string `json:"path,omitempty"`
		LocalCtime     int    `json:"local_ctime,omitempty"`
		ServerMtime    int    `json:"server_mtime,omitempty"`
		Empty          int    `json:"empty,omitempty"`
		OperID         int    `json:"oper_id,omitempty"`
		Thumbs         struct {
			URL3 string `json:"url3,omitempty"`
			URL2 string `json:"url2,omitempty"`
			URL1 string `json:"url1,omitempty"`
		} `json:"thumbs,omitempty"`
		Md5 string `json:"md5,omitempty"`
	} `json:"list,omitempty"`
	RequestID int64 `json:"request_id,omitempty"`
	GUID      int   `json:"guid,omitempty"`
}

//多参数,建议传struct进来
func GetFileList(t *FileListReqType) (*FileListRspType, error) {
	var (
		_url = "https://pan.baidu.com/rest/2.0/xpan/file"
	)
	v := url.Values{}
	v.Add("method", "list")

	v.Add("access_token", t.Access_token)
	v.Add("dir", defaultValue(t.Dir, "/"))
	v.Add("order", defaultValue(t.Order, "name")) //D^name,time,size
	v.Add("desc", defaultValue(t.Desc, "1"))
	v.Add("start", defaultValue(t.Start, "0"))
	v.Add("limit", defaultValue(t.Limit, "1000"))
	v.Add("web", defaultValue(t.Web, "web"))
	v.Add("folder", defaultValue(t.Folder, "0"))
	// v.Add("showempty", t.Showempty)

	req, _ := http.NewRequest("GET", _url+"?"+v.Encode(), nil)
	b, _ := httputil.DumpRequest(req, true)
	fmt.Println(string(b))
	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	var RspT FileListRspType
	// br, _ := ioutil.ReadAll(rsp.Body)
	// fmt.Println(string(br))
	if err := json.NewDecoder(rsp.Body).Decode(&RspT); err != nil {
		return nil, err
	}
	return &RspT, nil
}

func GetFileListByRecursion() {

}

func defaultValue(obj, default_v string) string {
	if obj != "" {
		return obj
	}
	return default_v
}
