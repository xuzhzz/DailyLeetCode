package dailyleetcode

import (
	"reflect"
	"testing"
)

func Test_sortMerge(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"test default use case", args{buildListNode("[4,2,1,3]")}, buildListNode("[1,2,3,4]")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortMerge(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortMerge() = %v, want %v", got, tt.want)
			}
		})
	}
}
