package str28

import (
	"reflect"
	"testing"
)

func Test_strStr(t *testing.T) {
	type args struct {
		haystack string
		needle   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strStr(tt.args.haystack, tt.args.needle); got != tt.want {
				t.Errorf("strStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getNext(t *testing.T) {
	type args struct {
		pa string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"abcy", args{"abcy"}, []int{0, 0, 0, 0}},
		{"ababcdaba", args{"ababcdaba"}, []int{0, 0, 1, 2, 0, 0, 1, 2, 3}},
		{"aabaaac", args{"aabaaac"}, []int{0, 1, 0, 1, 2, 2, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getNext(tt.args.pa); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getNext() = %v, want %v", got, tt.want)
			}
		})
	}
	// fmt.Printf("%v", getNext(tests[1].args.pa))
}
