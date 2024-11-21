package binary_search

import "testing"

func TestSearch(t *testing.T) {
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
			name: "Valid: When target exists into the array.",
			args: args{arr: []int{-1, 0, 2, 4, 6, 8}, target: 4},
			want: 3,
		},
		{
			name: "Valid: When target at the beginning of the array",
			args: args{arr: []int{-1, 0, 2, 4, 6, 8}, target: -1},
			want: 0,
		},
		{
			name: "Valid: When target doesn't exists in the array.",
			args: args{arr: []int{-1, 0, 2, 4, 6, 8}, target: 3},
			want: -1,
		},
		{
			name: "Valid: Target at the end of the array",
			args: args{arr: []int{-1, 0, 2, 4, 6, 8}, target: 8},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Search(tt.args.arr, tt.args.target); got != tt.want {
				t.Errorf("Search() = %v, want %v", got, tt.want)
			}
		})
	}
}
