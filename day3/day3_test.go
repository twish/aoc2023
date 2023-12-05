package day3

import "testing"

func TestOne(t *testing.T) {
	type args struct {
		inputFile string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"CalibrationNumberOfParts", args{inputFile: "./calib_data"}, 4361},
		{"NumberOfParts", args{inputFile: "./input_data"}, 527369},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := One(tt.args.inputFile); got != tt.want {
				t.Errorf("One() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTwo(t *testing.T) {
	type args struct {
		inputFile string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"GearRatioTotal", args{inputFile: "./input_data"}, 73074886},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Two(tt.args.inputFile); got != tt.want {
				t.Errorf("Two() = %v, want %v", got, tt.want)
			}
		})
	}
}
