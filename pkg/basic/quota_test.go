package basic

import (
	"BaiduPCS/tests"
	"fmt"
	"log"
	"testing"

	"github.com/tidwall/gjson"
)

func TestGetQuotaInfo(t *testing.T) {
	type args struct {
		access_token string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "Quota test",
			args: args{access_token: tests.PCS.Access_token},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetQuotaInfo(tt.args.access_token)
			if err != nil {
				t.Errorf("GetQuotaInfo() error = %v", err)
				return
			}
			log.Println(got)
			total_quota := gjson.Get(got, "total").Uint()
			used_quota := gjson.Get(got, "used").Uint()
			if Quotient := total_quota / uint64(tb); Quotient != 0 {
				fmt.Printf("%dG/%dT\n", used_quota/uint64(gb), Quotient)
			} else {
				fmt.Printf("%dG/%dG\n", used_quota/uint64(gb), total_quota/uint64(gb))
			}
		})
	}
}
