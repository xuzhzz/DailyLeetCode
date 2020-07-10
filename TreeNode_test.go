package dailyleetcode

import (
	"reflect"
	"testing"
)

func Test_buildTree(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{"test str", args{in: "[123, 1234, 52151, 15d, null]"}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildTree(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
