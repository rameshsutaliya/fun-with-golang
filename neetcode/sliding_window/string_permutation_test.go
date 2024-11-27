package sliding_window

import "testing"

func Test_checkInclusion(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test:1, When s2 contains a permutation of s1",
			args: args{s1: "ab", s2: "eidbaooo"},
			want: true,
		},
		{
			name: "Test:2, When s2 does not contains a permutation of s1",
			args: args{s1: "ab", s2: "eidboaoo"},
			want: false,
		},
		{
			name: "Test:1, When s1 and s2 are same",
			args: args{s1: "ab", s2: "ab"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkInclusion(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("checkInclusion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isEqual(t *testing.T) {
	type args struct {
		arr1 []int
		arr2 []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test: 1, When both arrays are same",
			args: args{arr1: []int{1, 2, 3, 4, 5}, arr2: []int{1, 2, 3, 4, 5}},
			want: true,
		},
		{
			name: "Test: 1, When both arrays are not same",
			args: args{arr1: []int{1, 2, 3, 4, 5}, arr2: []int{1, 2, 3, 5, 5}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isEqual(tt.args.arr1, tt.args.arr2); got != tt.want {
				t.Errorf("isEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}
