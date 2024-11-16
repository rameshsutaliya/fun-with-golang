package array_hashing

import (
	"reflect"
	"testing"
)

func TestProductWithoutIncludingSelf(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{
			name:    "Valid-Test: Input list with length 2",
			args:    args{arr: []int{1, 2}},
			want:    []int{2, 1},
			wantErr: false,
		},
		{
			name:    "Valid-Test: Input list with positive numbers",
			args:    args{arr: []int{1, 2, 3, 4}},
			want:    []int{24, 12, 8, 6},
			wantErr: false,
		},
		{
			name:    "Valid-Test: Input list with integer numbers",
			args:    args{arr: []int{-1, 1, 0, -3, 3}},
			want:    []int{0, 0, 9, 0, 0},
			wantErr: false,
		},
		{
			name:    "In-Valid-Test: Input list with length <2",
			args:    args{arr: []int{1}},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ProductWithoutIncludingSelf(tt.args.arr)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProductWithoutIncludingSelf() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductWithoutIncludingSelf() got = %v, want %v", got, tt.want)
			}
		})
	}
}
