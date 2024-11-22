package binary_search

import "testing"

func TestBinarySearch(t *testing.T) {
	type args struct {
		board  [][]int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Valid: When Solution exist",
			args: args{board: [][]int{
				{1, 2, 4, 8},
				{10, 11, 12, 13},
				{14, 20, 30, 40}}, target: 10},
			want: true,
		},
		{
			name: "Valid: When Solution does not exist",
			args: args{
				board:  [][]int{{1, 2, 4, 8}, {10, 11, 12, 13}, {14, 20, 30, 40}},
				target: 15,
			},
			want: false,
		},
		{
			name: "Valid: Matrix with single element, and target exists",
			args: args{board: [][]int{{1}}, target: 1},
			want: true,
		},
		{
			name: "Valid: Matrix with single element, and target does not exists",
			args: args{board: [][]int{{5}}, target: 4},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinarySearch(tt.args.board, tt.args.target); got != tt.want {
				t.Errorf("BinarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
