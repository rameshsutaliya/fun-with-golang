package array_hashing

import "testing"

func TestIsValidPalindrome(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Valid: Input string is a palindrome.",
			args: args{str: "Was it a car or a cat I saw?"},
			want: true,
		},
		{
			name: "Valid: Input with mixed values",
			args: args{"a12ba22"},
			want: false,
		},
		{
			name: "Invalid: Palindrome",
			args: args{str: "tab a cat"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidPalindrome(tt.args.str); got != tt.want {
				t.Errorf("IsValidPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
