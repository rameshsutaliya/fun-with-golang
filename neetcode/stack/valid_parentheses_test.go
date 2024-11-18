package stack

import "testing"

func TestIsValidParentheses(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Valid: Single Parentheses",
			args: args{str: "()"},
			want: true,
		},
		{
			name: "Valid: Multiple Continues Parentheses",
			args: args{str: "()[]{}"},
			want: true,
		},
		{
			name: "Valid: Multiple nested Parentheses",
			args: args{str: "(()){([])}()()"},
			want: true,
		},
		{
			name: "Invalid: All Open Parentheses",
			args: args{str: "(((({{{[[[[]]"},
			want: false,
		},
		{
			name: "Invalid: Invalid mapping of Parentheses",
			args: args{str: "(({{]})]"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidParentheses(tt.args.str); got != tt.want {
				t.Errorf("IsValidParentheses() = %v, want %v", got, tt.want)
			}
		})
	}
}
