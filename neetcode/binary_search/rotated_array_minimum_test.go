package binary_search

import "testing"

func TestGetMin(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test:1",
			args: args{arr: []int{3, 4, 5, 1, 2}},
			want: 1,
		},
		{
			name: "Test:2",
			args: args{arr: []int{4, 5, 6, 7, 0, 1, 2}},
			want: 0,
		},
		{
			name: "Test:3",
			args: args{arr: []int{11, 13, 15, 17}},
			want: 11,
		},
		{
			name: "Test:3",
			args: args{arr: []int{11, 13, 15, 17, 10}},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetMin(tt.args.arr); got != tt.want {
				t.Errorf("GetMin() = %v, want %v", got, tt.want)
			}
		})
	}
}
