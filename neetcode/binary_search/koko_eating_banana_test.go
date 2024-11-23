package binary_search

import "testing"

func TestMinimumEatingRate(t *testing.T) {
	type args struct {
		piles []int
		hours int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Valid: Test case 1",
			args: args{piles: []int{3, 6, 7, 11}, hours: 8},
			want: 4,
		},
		{
			name: "Valid: Test case 2",
			args: args{piles: []int{30, 11, 23, 4, 20}, hours: 5},
			want: 30,
		},
		{
			name: "Valid: Test case 3",
			args: args{piles: []int{30, 11, 23, 4, 20}, hours: 6},
			want: 23,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinimumEatingRate(tt.args.piles, tt.args.hours); got != tt.want {
				t.Errorf("MinimumEatingRate() = %v, want %v", got, tt.want)
			}
		})
	}
}
