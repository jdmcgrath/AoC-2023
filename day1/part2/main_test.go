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
