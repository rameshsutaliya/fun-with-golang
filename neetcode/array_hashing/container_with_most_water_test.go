package array_hashing

import "testing"

func TestContainerVolume(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Valid: With random value input list",
			args: args{arr: []int{1, 8, 6, 2, 5, 4, 8, 3, 7}},
			want: 49,
		},
		{
			name: "Valid: with single distance",
			args: args{arr: []int{1, 1}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainerVolume(tt.args.arr); got != tt.want {
				t.Errorf("ContainerVolume() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getMin(t *testing.T) {
	type args struct {
		val1 int
		val2 int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Valid: With two positive numbers",
			args: args{val1: 5, val2: 9},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMin(tt.args.val1, tt.args.val2); got != tt.want {
				t.Errorf("getMin() = %v, want %v", got, tt.want)
			}
		})
	}
}
