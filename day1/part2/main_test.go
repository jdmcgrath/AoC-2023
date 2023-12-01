package main

import (
	"reflect"
	"testing"
)

func Test_getDigitsInString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Rubbish at front",
			args: args{
				s: "zoseight234",
			},
			want: []string{
				"8", "2", "3", "4",
			},
		},
		{
			name: "Example",
			args: args{
				s: "two1nine",
			},
			want: []string{
				"2", "1", "9",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getDigitsInString(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getDigitsInString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getFirstAndLastFromString(t *testing.T) {
	type args struct {
		inputString string
	}
	tests := []struct {
		name      string
		args      args
		wantFirst string
		wantLast  string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFirst, gotLast := getFirstAndLastFromString(tt.args.inputString)
			if gotFirst != tt.wantFirst {
				t.Errorf("getFirstAndLastFromString() gotFirst = %v, want %v", gotFirst, tt.wantFirst)
			}
			if gotLast != tt.wantLast {
				t.Errorf("getFirstAndLastFromString() gotLast = %v, want %v", gotLast, tt.wantLast)
			}
		})
	}
}

func Test_processFile(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			processFile(tt.args.path)
		})
	}
}

func Test_processLine(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := processLine(tt.args.line); got != tt.want {
				t.Errorf("processLine() = %v, want %v", got, tt.want)
			}
		})
	}
}
