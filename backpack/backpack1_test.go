package backpack

import "testing"

func Test_backPack1(t *testing.T) {
	type args struct {
		goods []int
		size  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test case 1", args{[]int{3, 4, 8, 5}, 10}, 9},
		{"test case 2", args{[]int{2, 3, 5, 7}, 12}, 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := backPack1(tt.args.goods, tt.args.size); got != tt.want {
				t.Errorf("backPack1() = %v, want %v", got, tt.want)
			}
		})
	}
}
