package tree

import (
	"fmt"
	"testing"
)

func TestNewPIBinaryTreeNode(t *testing.T) {
	type args struct {
		preorder []int
		inorder  []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "1",
			args: args{
				preorder: []int{1, 2, 4, 5, 3, 6, 7},
				inorder:  []int{4, 2, 5, 1, 6, 3, 7},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewPIBinaryTreeNode(tt.args.preorder, tt.args.inorder)
			fmt.Println(got)
		})
	}
}
