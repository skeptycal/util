package http

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestGetPage(t *testing.T) {
	type args struct {
		url   string
		key   string
		value string
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{"google", args{"https://httpbin.org/get", "Host", "httpbin.org"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotBuff, err := GetPage(tt.args.url)
			got := json.Unmarshal(gotBuff.Bytes(), v)
			got := gotBuff.Bytes()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPage() = %v, want %v", got, tt.want)
			}
		})
	}
}
