package adminServices

import (
	"testing"
)

func TestGetAnalysisUsetime(t *testing.T) {
	type args struct {
		data map[string]string
	}

	tests := []struct {
		name           string
		args           args
		wantResultCode int
		wantResult     interface{}
	}{
		{"AnalysisUsetimeSuccess", args{map[string]string{"startdate": "2021-10-01", "enddate": "2021-10-31"}}, 0, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResultCode, _ := GetAnalysisUsetime(tt.args.data)
			if gotResultCode != tt.wantResultCode {
				t.Errorf("GetAnalysisUsetime() gotResultCode = %v, want %v", gotResultCode, tt.wantResultCode)
			}
		})
	}
}

func TestGetAnalysisState(t *testing.T) {
	tests := []struct {
		name           string
		wantResultCode int
		wantResult     interface{}
	}{
		{"AnalysisState", 0, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResultCode, _ := GetAnalysisState()
			if gotResultCode != tt.wantResultCode {
				t.Errorf("GetAnalysisState() gotResultCode = %v, want %v", gotResultCode, tt.wantResultCode)
			}
		})
	}
}

func TestGetAnalysisFeedback(t *testing.T) {
	tests := []struct {
		name           string
		wantResultCode int
		wantResult     interface{}
	}{
		{"AnalysisFeedback", 0, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResultCode, _ := GetAnalysisFeedback()
			if gotResultCode != tt.wantResultCode {
				t.Errorf("GetAnalysisFeedback() gotResultCode = %v, want %v", gotResultCode, tt.wantResultCode)
			}
		})
	}
}

func TestGetAnalysisAge(t *testing.T) {
	tests := []struct {
		name           string
		wantResultCode int
		wantResult     interface{}
	}{
		{"AnalysisAge", 0, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResultCode, _ := GetAnalysisAge()
			if gotResultCode != tt.wantResultCode {
				t.Errorf("GetAnalysisAge() gotResultCode = %v, want %v", gotResultCode, tt.wantResultCode)
			}
		})
	}
}

func Test_getOnlieCount(t *testing.T) {
	var number int
	tests := []struct {
		name       string
		wantResult int
	}{
		{"OnlieCount", number},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := getOnlieCount(); gotResult != tt.wantResult {
				t.Errorf("getOnlieCount() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
