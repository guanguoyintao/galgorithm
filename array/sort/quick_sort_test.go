package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuickSortWithSwap(t *testing.T) {
	type args struct {
		list []int
		low  int
		high int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{list: []int{2, 3, 5, 7, 18, 20, 1, 0, -1, 4, 3}, low: 0, high: 10},
			want: []int{-1, 0, 1, 2, 3, 3, 4, 5, 7, 18, 20},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := QuickSortBySwap(tt.args.list, tt.args.low, tt.args.high)
			assert.Equal(t, result, tt.args.list, "sort Error!")
		})
	}
}

func TestQuickSortWithDig(t *testing.T) {
	type args struct {
		list []int
		low  int
		high int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{list: []int{2, 3, 5, 7, 18, 20, 1, 0, -1, 4, 3}, low: 0, high: 10},
			want: []int{-1, 0, 1, 2, 3, 3, 4, 5, 7, 18, 20},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, QuickSortByDig(tt.args.list, tt.args.low, tt.args.high), "QuickSortWithDig(%v, %v, %v)", tt.args.list, tt.args.low, tt.args.high)
		})
	}
}
