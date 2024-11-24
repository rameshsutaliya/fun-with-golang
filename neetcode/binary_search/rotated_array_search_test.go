package binary_search

import "testing"

func TestRotatedArraySearch(t *testing.T) {
	type args struct {
		arr    []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Search In Rotated Sorted array Test:1",
			args: args{arr: []int{4, 5, 6, 7, 0, 1, 2}, target: 0},
			want: 4,
		},
		{
			name: "Search In Rotated Sorted array Test:2",
			args: args{arr: []int{4, 5, 6, 7, 0, 1, 2}, target: 8},
			want: -1,
		},
		{
			name: "Search In Rotated Sorted array Test:3",
			args: args{arr: []int{0, 1, 2, 4, 5, 6, 7}, target: 5},
			want: 4,
		},
		{
			name: "Search In Rotated Sorted array Test:4",
			args: args{arr: []int{0}, target: 0},
			want: 0,
		},
		{
			name: "Search In Rotated Sorted array Test:5",
			args: args{arr: []int{3, 0}, target: 0},
			want: 1,
		},
		{
			name: "Search In Rotated Sorted array Test:6",
			args: args{arr: []int{3, 0}, target: 3},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RotatedArraySearch(tt.args.arr, tt.args.target); got != tt.want {
				t.Errorf("RotatedArraySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_binarySearch(t *testing.T) {
	type args struct {
		arr    []int
		left   int
		right  int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Binary Search Test:1",
			args: args{arr: []int{1}, target: 1},
			want: 0,
		},
		{
			name: "Binary Search Test:2",
			args: args{arr: []int{1, 2, 3, 4, 5, 6}, target: 1},
			want: 0,
		},
		{
			name: "Binary Search Test:3",
			args: args{arr: []int{1, 2, 3, 4, 5, 6}, target: 10},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binarySearch(tt.args.arr, tt.args.left, tt.args.right, tt.args.target); got != tt.want {
				t.Errorf("binarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rotationPoint(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Rotation point test:1",
			args: args{arr: []int{4, 5, 6, 7, 0, 1, 2}},
			want: 4,
		},
		{
			name: "Rotation point test:2",
			args: args{arr: []int{4, 5, 6, 7, 0}},
			want: 4,
		},
		{
			name: "Rotation point test:3",
			args: args{arr: []int{0, 1, 2, 4, 5, 6, 7}},
			want: 0,
		},
		{
			name: "Rotation point test:4",
			args: args{arr: []int{2, 1}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotationPoint(tt.args.arr); got != tt.want {
				t.Errorf("rotationPoint() = %v, want %v", got, tt.want)
			}
		})
	}
}
