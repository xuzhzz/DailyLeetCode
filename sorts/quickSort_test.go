package sorts

import (
	"reflect"
	"testing"
)

func Test_quickSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test case1", args{[]int{7, 3, 8, 6, 2, 9, 1}}, []int{1, 2, 3, 6, 7, 8, 9}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := quickSort(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("quickSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
