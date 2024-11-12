package array_hashing

import "testing"

func TestIsAnagram(t *testing.T) {
	type args struct {
		first  string
		second string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Valid Anagram",
			args: args{
				first:  "racecar",
				second: "carrace",
			},
			want: true,
		},
		{
			name: "Invalid Anagram",
			args: args{first: "jar", second: "jam"},
			want: false,
		},
		{
			name: "Valid: Empty string",
			args: args{
				first:  "",
				second: "",
			},
			want: true,
		},
		{
			name: "Uneven length strings",
			args: args{
				first:  "abcdef",
				second: "abc",
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsAnagram(tt.args.first, tt.args.second); got != tt.want {
				t.Errorf("IsAnagram() = %v, want %v", got, tt.want)
			}
		})
	}
}
