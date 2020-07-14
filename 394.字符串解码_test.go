package dailyleetcode

import "testing"

func Test_decodeString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test case 1", args{"3[a]2[bc]"}, "aaabcbc"},
		{"test case 2", args{"3[a2[c]]"}, "accaccacc"},
		{"test case 3", args{"2[abc]3[cd]ef"}, "abcabccdcdcdef"},
		{"test case 4", args{"abc3[cd]xyz"}, "abccdcdcdxyz"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decodeString(tt.args.s); got != tt.want {
				t.Errorf("decodeString() = %v, want %v", got, tt.want)
			}
		})
	}
}
