package dp

import "testing"

func Test_tribonacci(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test:1",
			args: args{n: 0},
			want: 0,
		},
		{
			name: "Test:2",
			args: args{n: 2},
			want: 1,
		},
		{
			name: "Test:3",
			args: args{n: 25},
			want: 1389537,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tribonacci(tt.args.n); got != tt.want {
				t.Errorf("tribonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}
