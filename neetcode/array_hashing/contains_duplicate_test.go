package array_hashing

import "testing"

func TestContainsDuplicate(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Valid-Contains duplicate",
			args: args{arr: []int{1, 2, 3, 3}},
			want: true,
		},
		{
			name: "Valid-No Duplicate",
			args: args{arr: []int{1, 2, 3, 4}},
			want: false,
		},
		{
			name: "Valid-Empty Array",
			args: args{arr: []int{}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsDuplicate(tt.args.arr); got != tt.want {
				t.Errorf("ContainsDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
