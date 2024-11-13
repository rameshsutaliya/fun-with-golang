package array_hashing

import (
	"reflect"
	"testing"
)

func TestGetTwoSum(t *testing.T) {
	type args struct {
		arr    []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Valid: Proper Solution",
			args: args{
				arr:    []int{3, 4, 5, 6},
				target: 7,
			},
			want: []int{0, 1},
		},
		{
			name: "Valid: Any Combination.",
			args: args{arr: []int{4, 5, 6}, target: 10},
			want: []int{0, 2},
		},
		{
			name: "Invalid: No Solution exists.",
			args: args{arr: []int{1, 2, 3, 4, 5, 6}, target: 13},
			want: []int{-1, -1},
		},
		{
			name: "Invalid: Empty input array",
			args: args{arr: []int{}, target: 5},
			want: []int{-1, -1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetTwoSum(tt.args.arr, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTwoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
