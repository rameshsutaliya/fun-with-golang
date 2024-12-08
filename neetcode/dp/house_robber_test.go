package dp

import "testing"

func Test_rob(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test: 1 Rob alternative hoses.",
			args: args{nums: []int{1, 2, 3, 1}},
			want: 4,
		},
		{
			name: "Test: 2- Rob alternative hoses.",
			args: args{nums: []int{2, 7, 9, 3, 1}},
			want: 12,
		},
		{
			name: "Test: 3- Rob considering adjacent houses and their adjacent",
			args: args{nums: []int{4, 2, 3, 4}},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rob(tt.args.nums); got != tt.want {
				t.Errorf("rob() = %v, want %v", got, tt.want)
			}
		})
	}
}
