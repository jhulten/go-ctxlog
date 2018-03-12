package ctxlog

import (
	"reflect"
	"testing"
)

func Test_stringFormatter_Format(t *testing.T) {
	type args struct {
		i interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"string",
			args{"log string"},
			"log string",
		},
		{
			"slice of strings",
			args{[]string{"log", "string"}},
			"log, string",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &stringFormatter{}
			if got := f.Format(tt.args.i); got != tt.want {
				t.Errorf("stringFormatter.Format() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_formatValue(t *testing.T) {
	type args struct {
		v reflect.Value
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
			if got := formatValue(tt.args.v); got != tt.want {
				t.Errorf("formatValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_head(t *testing.T) {
	type args struct {
		v reflect.Value
	}
	tests := []struct {
		name  string
		args  args
		want  reflect.Value
		want1 reflect.Value
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := head(tt.args.v)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("head() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("head() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
