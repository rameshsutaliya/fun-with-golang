package sliding_window

import "testing"

func Test_firstContainsSecond(t *testing.T) {
	type args struct {
		tMap map[byte]int
		sMap map[byte]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test: When s map has the t map",
			args: args{tMap: map[byte]int{'a': 5, 'b': 2, 'c': 4}, sMap: map[byte]int{'a': 6, 'b': 2, 'c': 5, 'd': 1, 'e': 4}},
			want: true,
		},
		{
			name: "Test: When s map does not have the t map",
			args: args{tMap: map[byte]int{'a': 5, 'b': 2, 'c': 4}, sMap: map[byte]int{'a': 3, 'b': 2, 'c': 5, 'd': 1, 'e': 4}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstContainsSecond(tt.args.tMap, tt.args.sMap); got != tt.want {
				t.Errorf("firstContainsSecond() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minWindow(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test: 1",
			args: args{s: "abcdefg", t: "ace"},
			want: "abcde",
		},
		{
			name: "Test: 2",
			args: args{s: "ADOBECODEBANC", t: "ABC"},
			want: "BANC",
		},
		{
			name: "Test: 3",
			args: args{s: "a", t: "a"},
			want: "a",
		},
		{
			name: "Test: 4",
			args: args{s: "a", t: "aa"},
			want: "",
		},
		{
			name: "Test: 5",
			args: args{s: "OUZODYXAZV", t: "XYZ"},
			want: "YXAZ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minWindow(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("minWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}
