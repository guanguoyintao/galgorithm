package max_depth

import (
	"testing"

	"github.com/guanguoyintao/galgorithm/tree"
)

func TestMaxDepth(t *testing.T) {
	type args struct {
		root *tree.BinaryTreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				root: tree.NewPIBinaryTreeNode([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}),
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DFSMaxDepth(tt.args.root); got != tt.want {
				t.Errorf("MaxDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
