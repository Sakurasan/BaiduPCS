package basic

import (
	"BaiduPCS/tests"
	"fmt"
	"log"
	"testing"
)

func TestGetFileList(t *testing.T) {
	type args struct {
		t *FileListReqType
	}
	var reqType = &FileListReqType{
		Access_token: tests.PCS.Access_token,
		Order:        "time",
		Desc:         "1",
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "TestGetFileList",
			args: args{t: reqType},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetFileList(tt.args.t)
			if err != nil {
				t.Errorf("GetFileList() error = %v", err)
				return
			}
			// b, _ := json.Marshal(got)
			// t.Logf("GetFileList() = %v", string(b))
			fdir := func(isdir int) string {
				if isdir == 1 {
					return "Dir"
				}
				return "file"
			}
			fsize := func(size int) string {
				if size == 0 {
					return "-"
				} else if q := size / gb; q > 0 {
					return fmt.Sprintf("%d G", q)
				} else {
					return fmt.Sprintf("%d M", size/mb)
				}
			}
			for _, v := range got.List {
				log.Printf("%3s %50s %5s", fdir(v.Isdir), v.Path, fsize(v.Size))
				log.Println(v.FsID)
			}

		})
	}
}
