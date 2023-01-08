package encoding

import (
	"reflect"
	"testing"
)

func TestGetBytes(t *testing.T) {
	type args struct {
		data    []byte
		charset Charset
	}
	tests := []struct {
		name      string
		args      args
		wantCdata []byte
		wantErr   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCdata, err := GetBytes(tt.args.data, tt.args.charset)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotCdata, tt.wantCdata) {
				t.Errorf("GetBytes() = %v, want %v", gotCdata, tt.wantCdata)
			}
		})
	}
}

func TestGetBytes0(t *testing.T) {
	type args struct {
		data    []byte
		charset Charset
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetBytes0(tt.args.data, tt.args.charset)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBytes0() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBytes0() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertByte2String(t *testing.T) {
	type args struct {
		byte    []byte
		charset Charset
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertByte2String(tt.args.byte, tt.args.charset); got != tt.want {
				t.Errorf("ConvertByte2String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertToString(t *testing.T) {
	type args struct {
		src     string
		srcCode string
		tagCode string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertToString(tt.args.src, tt.args.srcCode, tt.args.tagCode); got != tt.want {
				t.Errorf("ConvertToString() = %v, want %v", got, tt.want)
			}
		})
	}
}
