package zsh

import (
	"reflect"
	"testing"
)

func Test_winsize_String(t *testing.T) {
	type fields struct {
		row    uint16
		col    uint16
		xpixel uint16
		ypixel uint16
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
		{"example ws", fields{36, 78, 1404, 1512}, "TTY row: 36, col: 78 resolution 1404x1512"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ws := &winsize{
				row:    tt.fields.row,
				col:    tt.fields.col,
				xpixel: tt.fields.xpixel,
				ypixel: tt.fields.ypixel,
			}
			if got := ws.String(); got != tt.want {
				t.Errorf("winsize.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_winsize_Col(t *testing.T) {
	// ws := winsize{36, 78, 1404, 1512}
	type fields struct {
		row    uint16
		col    uint16
		xpixel uint16
		ypixel uint16
	}
	tests := []struct {
		name   string
		fields fields
		want   uint
	}{
		// TODO: Add test cases.
		{"example ws", fields{36, 78, 1404, 1512}, 78},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ws := &winsize{
				row:    tt.fields.row,
				col:    tt.fields.col,
				xpixel: tt.fields.xpixel,
				ypixel: tt.fields.ypixel,
			}
			if got := ws.Col(); got != tt.want {
				t.Errorf("winsize.Col() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_winsize_Row(t *testing.T) {
	type fields struct {
		row    uint16
		col    uint16
		xpixel uint16
		ypixel uint16
	}
	tests := []struct {
		name   string
		fields fields
		want   uint
	}{
		// TODO: Add test cases.
		{"example ws", fields{36, 78, 1404, 1512}, 36},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ws := &winsize{
				row:    tt.fields.row,
				col:    tt.fields.col,
				xpixel: tt.fields.xpixel,
				ypixel: tt.fields.ypixel,
			}
			if got := ws.Row(); got != tt.want {
				t.Errorf("winsize.Row() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_winsize_X(t *testing.T) {
	type fields struct {
		row    uint16
		col    uint16
		xpixel uint16
		ypixel uint16
	}
	tests := []struct {
		name   string
		fields fields
		want   uint
	}{
		// TODO: Add test cases.
		{"example ws", fields{36, 78, 1404, 1512}, 1404},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ws := &winsize{
				row:    tt.fields.row,
				col:    tt.fields.col,
				xpixel: tt.fields.xpixel,
				ypixel: tt.fields.ypixel,
			}
			if got := ws.X(); got != tt.want {
				t.Errorf("winsize.X() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_winsize_Y(t *testing.T) {
	type fields struct {
		row    uint16
		col    uint16
		xpixel uint16
		ypixel uint16
	}
	tests := []struct {
		name   string
		fields fields
		want   uint
	}{
		// TODO: Add test cases.
		{"example ws", fields{36, 78, 1404, 1512}, 1512},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ws := &winsize{
				row:    tt.fields.row,
				col:    tt.fields.col,
				xpixel: tt.fields.xpixel,
				ypixel: tt.fields.ypixel,
			}
			if got := ws.Y(); got != tt.want {
				t.Errorf("winsize.Y() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewTTY(t *testing.T) {
	ws := winsize{}
	tests := []struct {
		name string
		want TTY
	}{
		// TODO: Add test cases.
		{"tty", &ws},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTTY(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTTY() = %v, want %v", got, tt.want)
			}
		})
	}
}
