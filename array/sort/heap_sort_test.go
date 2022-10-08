package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeapSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{arr: []int{2, 3, 5, 7, 18, 20, 1, 0, -1, 4, 3}},
			want: []int{-1, 0, 1, 2, 3, 3, 4, 5, 7, 18, 20},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, HeapSort(tt.args.arr), "HeapSort(%v)", tt.args.arr)
		})
	}
}
