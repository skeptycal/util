package build

import (
	"reflect"
	"testing"
)

func Test_goose_Active(t *testing.T) {
	type fields struct {
		goos   string
		goarch []string
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		// TODO: Add test cases.
		{"blank", fields{"", []string{""}}, []string{""}},
		{"darwin", fields{"darwin", []string{"amd64"}}, []string{""}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &goose{
				goos:   tt.fields.goos,
				goarch: tt.fields.goarch,
			}
			if got := g.Active(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("goose.Active() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_goose_String(t *testing.T) {
	type fields struct {
		goos   string
		goarch []string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
		{"test", fields{"", []string{"", "", "", ""}}, "GOOS=${} GOARCH=${[   ]}"},
		{"darwin", fields{"darwin", []string{"amd64"}}, "GOOS=${darwin} GOARCH=${[amd64]}"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &goose{
				goos:   tt.fields.goos,
				goarch: tt.fields.goarch,
			}
			if got := g.String(); got != tt.want {
				t.Errorf("goose.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
