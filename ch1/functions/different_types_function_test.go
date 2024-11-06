package functions

import "testing"

func TestAdd(t *testing.T) {
	type args struct {
		num1 int
		num2 int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Valid-Positive",
			args: args{2, 5},
			want: 7,
		},
		{
			name: "Valid Sum - negative numbers",
			args: args{
				num1: -3,
				num2: -4,
			},
			want: -7,
		},
		{
			name: "Valid Sum - mix of positive and negative",
			args: args{
				num1: 5,
				num2: -3,
			},
			want: 2,
		},
		{
			name: "Valid Sum - zero",
			args: args{
				num1: 0,
				num2: 5,
			},
			want: 5,
		},
		{
			name: "Valid Sum - zero result",
			args: args{
				num1: -5,
				num2: 5,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGreet(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Greet String",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Greet()
		})
	}
}

func TestNamed(t *testing.T) {
	type args struct {
		val int
	}
	tests := []struct {
		name           string
		args           args
		wantFirstPart  int
		wantSecondPart int
	}{
		{
			name:           "Valid Split: Even Division",
			args:           args{8},
			wantFirstPart:  4,
			wantSecondPart: 4,
		},
		{
			name:           "Valid Split: Uneven Division",
			args:           args{9},
			wantFirstPart:  4,
			wantSecondPart: 5,
		},
		{
			name:           "Valid Split: Even Negative Number split",
			args:           args{-8},
			wantFirstPart:  -4,
			wantSecondPart: -4,
		},
		{
			name:           "Valid Split: Uneven negative split",
			args:           args{-13},
			wantFirstPart:  -6,
			wantSecondPart: -7,
		},
		{
			name:           "Valid Split: Zero split",
			args:           args{0},
			wantFirstPart:  0,
			wantSecondPart: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFirstPart, gotSecondPart := Named(tt.args.val)
			if gotFirstPart != tt.wantFirstPart {
				t.Errorf("Named() gotFirstPart = %v, want %v", gotFirstPart, tt.wantFirstPart)
			}
			if gotSecondPart != tt.wantSecondPart {
				t.Errorf("Named() gotSecondPart = %v, want %v", gotSecondPart, tt.wantSecondPart)
			}
		})
	}
}

func TestSwap(t *testing.T) {
	type args struct {
		num1 int
		num2 int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{
			name:  "Valid: Swap two numbers",
			args:  args{1, 2},
			want:  2,
			want1: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Swap(tt.args.num1, tt.args.num2)
			if got != tt.want {
				t.Errorf("Swap() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Swap() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestGetGreetMsg(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "Valid Greet msg",
			want: "˗ˏˋ ★ ˎˊ˗ welcome !!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetGreetMsg(); got != tt.want {
				t.Errorf("GetGreetMsg() = %v, want %v", got, tt.want)
			}
		})
	}
}
