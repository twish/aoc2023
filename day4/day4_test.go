package day4

import (
	"reflect"
	"testing"
)

func Test_readNumbers(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name        string
		args        args
		wantNumbers []int
		wantErr     bool
	}{
		{"WorkingNumberStringParseTrailingWhiteSpace", args{s: " 83 86  6 31 17  9 48 53   "}, []int{83, 86, 6, 31, 17, 9, 48, 53}, false},
		{"WorkingNumberStringParse", args{s: " 83 86  6 31 17  9 48 53"}, []int{83, 86, 6, 31, 17, 9, 48, 53}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNumbers, err := readNumbers(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("readNumbers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotNumbers, tt.wantNumbers) {
				t.Errorf("readNumbers() gotNumbers = %v, want %v", gotNumbers, tt.wantNumbers)
			}
		})
	}
}

func Test_parseToCard(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name     string
		args     args
		wantCard Card
		wantErr  bool
	}{
		{"TestParseACard", args{text: "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53"}, Card{id: 1}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCard, err := parseToCard(tt.args.text)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseToCard() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotCard, tt.wantCard) {
				t.Errorf("parseToCard() gotCard = %v, want %v", gotCard, tt.wantCard)
			}
		})
	}
}

func TestOne(t *testing.T) {
	type args struct {
		inputFile string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"TestOne", args{inputFile: "./input_data"}, 21568},
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
		{"TestTwo", args{inputFile: "./input_data"}, 11827296},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Two(tt.args.inputFile); got != tt.want {
				t.Errorf("Two() = %v, want %v", got, tt.want)
			}
		})
	}
}
