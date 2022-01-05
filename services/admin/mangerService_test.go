package adminServices

import (
	"reflect"
	"testing"
)

func TestGetManagerList(t *testing.T) {
	tests := []struct {
		name           string
		wantResultCode int
		wantResult     interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResultCode, gotResult := GetManagerList()
			if gotResultCode != tt.wantResultCode {
				t.Errorf("GetManagerList() gotResultCode = %v, want %v", gotResultCode, tt.wantResultCode)
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("GetManagerList() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestDelManager(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name           string
		args           args
		wantResultCode int
		wantResult     interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResultCode, gotResult := DelManager(tt.args.data)
			if gotResultCode != tt.wantResultCode {
				t.Errorf("DelManager() gotResultCode = %v, want %v", gotResultCode, tt.wantResultCode)
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("DelManager() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestSetManager(t *testing.T) {
	type args struct {
		email ManagerEmail
	}
	tests := []struct {
		name           string
		args           args
		wantResultCode int
		wantResult     interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResultCode, gotResult := SetManager(tt.args.email)
			if gotResultCode != tt.wantResultCode {
				t.Errorf("SetManager() gotResultCode = %v, want %v", gotResultCode, tt.wantResultCode)
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("SetManager() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
