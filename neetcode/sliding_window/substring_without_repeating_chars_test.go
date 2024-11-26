package sliding_window

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Valid: Test:1",
			args: args{s: ""},
			want: 0,
		},
		{
			name: "Valid: Test:2",
			args: args{s: "  "},
			want: 1,
		},
		{
			name: "Valid: Test:3",
			args: args{s: "abcdefghijklmnopqrstuvwxyz1234567890!@#$%^&*()-=_+AZXCVBNM,.;POIUYTREWQWSDFGHJK"},
			want: 71,
		},
		{
			name: "Valid: Test:4",
			args: args{s: "pwwkew"},
			want: 3,
		},
		{
			name: "Valid: Test:5",
			args: args{s: "aaaaa"},
			want: 1,
		},
		{
			name: "Valid: Test:6",
			args: args{s: "abcabcbb"},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("LengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
