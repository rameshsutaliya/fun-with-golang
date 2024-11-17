package array_hashing

import (
	"reflect"
	"testing"
)

func TestTwoSumIndexes(t *testing.T) {
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
			name: "Valid: Valid sorted array data",
			args: args{
				arr:    []int{2, 7, 11, 15},
				target: 9,
			},
			want: []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TwoSumIndexes(tt.args.arr, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TwoSumIndexes() = %v, want %v", got, tt.want)
			}
		})
	}
}
