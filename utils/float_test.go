package utils

import (
	"fmt"
	"testing"
)

func Test_isClose(t *testing.T) {
	type args struct {
		a float64
		b float64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{a: 0.3, b: 0.1 + 0.2},
			want: true,
		},
		{
			name: "1",
			args: args{a: 0.7071067811865475, b: 0.7071},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isClose(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("isClose() = %v, want %v", got, tt.want)
			} else {
				fmt.Println(got, tt.want)
			}
		})
	}
}
