package basic

import (
	"log"
	"testing"
)

func TestGetUserInfo(t *testing.T) {
	type args struct {
		access_token string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "get user info",
			args: args{access_token: "121.9e1b535796673c3e6cdcdf60b5df92e4.Ynl-AOYVtTV2ULUa9DSr4SU2wNILH7HoU2tFDsO.DO8SkA"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetUserInfo(tt.args.access_token)
			if err != nil {
				t.Errorf("GetUserInfo() error = %v", err)
				return
			}
			log.Println(got)
		})
	}
}
