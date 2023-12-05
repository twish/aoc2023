package day2

import "testing"

func TestOne(t *testing.T) {
	type args struct {
		inputFile string
		bag       GameBag
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"GetTotal", args{inputFile: "./input_data_1", bag: GameBag{Reds: 12, Greens: 13, Blues: 14}}, 2285},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := One(tt.args.inputFile, tt.args.bag); got != tt.want {
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
		{"PowerSum", args{inputFile: "./input_data_1"}, 77021},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Two(tt.args.inputFile); got != tt.want {
				t.Errorf("Two() = %v, want %v", got, tt.want)
			}
		})
	}
}
