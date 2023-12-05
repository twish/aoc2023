package day1

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOne_Get_55701_AsValueForSolution(t *testing.T) {
	total := Two("./input_data")
	fmt.Println(total)
	assert.Equal(t, 55701, total)
}

func TestOne_Get_281_AsValueForSolution(t *testing.T) {
	total := Two("./calib_data")
	fmt.Println(total)
	assert.Equal(t, 281, total)
}

func Test_getFirstNumWithWord(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"firstIntShouldReturn", args{line: "oneight1four4"}, "1"},
		{"firstWordShouldReturn", args{line: "three1four4"}, "3"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, getFirstNumWithWord(tt.args.line), "getFirstNumWithWord(%v)", tt.args.line)
		})
	}
}

func Test_getLastNumWithWord(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"firstIntShouldReturn", args{line: "1four4oneight"}, "8"},
		{"firstWordShouldReturn", args{line: "three1four4five"}, "5"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, getLastNumWithWord(tt.args.line), "getLastNumWithWord(%v)", tt.args.line)
		})
	}
}
