package dailyleetcode

import (
	"reflect"
	"testing"
)

func Test_buildListNode(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test build and print", args{"[1,2,3,4,5]"}, "1->2->3->4->5"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := printListNode(buildListNode(tt.args.in)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildListNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
