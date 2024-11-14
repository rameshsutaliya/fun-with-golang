package array_hashing

import (
	"reflect"
	"testing"
)

func TestGroupAnagram(t *testing.T) {
	type args struct {
		list []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "Valid: Multiple Anagram Grouping",
			args: args{list: []string{"act", "pots", "tops", "cat", "stop", "hat"}},
			want: [][]string{
				{"act", "cat"},
				{"pots", "tops", "stop"},
				{"hat"},
			},
		},
		{
			name: "Valid: Empty String",
			args: args{list: []string{""}},
			want: [][]string{
				{""},
			},
		},
		{
			name: "Valid: Single Character string",
			args: args{list: []string{"a"}},
			want: [][]string{
				{"a"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GroupAnagram(tt.args.list); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GroupAnagram() = %v, want %v", got, tt.want)
			}
		})
	}
}
