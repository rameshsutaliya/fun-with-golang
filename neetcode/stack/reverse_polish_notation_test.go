package stack

import "testing"

func TestEvalRPN(t *testing.T) {
	type args struct {
		tokens []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Valid: Single Operation",
			args: args{tokens: []string{"2", "1", "+"}},
			want: 3,
		},
		{
			name: "Valid: Multiple Operations",
			args: args{
				tokens: []string{"2", "1", "+", "3", "*"},
			},
			want: 9,
		},
		{
			name: "Valid: Numbers, Operations grouped",
			args: args{
				tokens: []string{"4", "13", "5", "/", "+"},
			},
			want: 6,
		},
		{
			name: "Valid: Higher number of operations",
			args: args{tokens: []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}},
			want: 22,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EvalRPN(tt.args.tokens); got != tt.want {
				t.Errorf("EvalRPN() = %v, want %v", got, tt.want)
			}
		})
	}
}
