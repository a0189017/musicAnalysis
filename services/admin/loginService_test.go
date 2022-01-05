package adminServices

import (
	"api/services"
	"reflect"
	"testing"
	"time"
)

func TestGetAdminLogin(t *testing.T) {
	type args struct {
		data MemberLogin
	}
	today := time.Now().Format("2006-01-02 15:04:05")
	tests := []struct {
		name           string
		args           args
		wantResultCode int
		wantResult     interface{}
	}{
		{"loginSuccess", args{MemberLogin{"12345@gmail.com", "5643"}}, 0, services.PasswordMd5("12345@gmail.com" + today + "project@#$S_ecrEt")},
		{"loginFAIL", args{MemberLogin{"123", "456"}}, 5, "帳號/密碼錯誤！"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResultCode, gotResult := GetAdminLogin(tt.args.data)
			if gotResultCode != tt.wantResultCode {
				t.Errorf("GetAdminLogin() gotResultCode = %v, want %v", gotResultCode, tt.wantResultCode)
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("GetAdminLogin() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
