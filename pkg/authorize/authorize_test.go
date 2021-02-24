package authorize

import (
	"testing"
)

func Test_getAuthCode(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "test",
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getAuthCode()
			if (err != nil) != tt.wantErr {
				t.Errorf("getAuthCode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getAuthCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_GetAccessToken(t *testing.T) {
	type args struct {
		code string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "token test",
			args: args{code: "cbd8de8cbeeb5357f00be6e928ac6e89"},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAccessTokenWithCode(tt.args.code)
			if err != nil {
				t.Errorf("getAccessToken() error = %v", err)
				return
			}
			if got != tt.want {
				t.Errorf("getAccessToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
