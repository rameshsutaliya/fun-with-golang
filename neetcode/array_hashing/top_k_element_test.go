package array_hashing

import (
	"reflect"
	"testing"
)

func TestKTopElement(t *testing.T) {
	type args struct {
		arr []int
		k   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Valid-Proper Test case",
			args: args{arr: []int{1, 2, 2, 3, 3, 3}, k: 2},
			want: []int{3, 2},
		},
		{
			name: "Valid-Same elements",
			args: args{arr: []int{7, 7}, k: 1},
			want: []int{7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := KTopElement(tt.args.arr, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("KTopElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
