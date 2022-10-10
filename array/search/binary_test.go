package search

import "testing"

func TestBinarySearch(t *testing.T) {
	type args struct {
		sortedArray []int
		lookingFor  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				sortedArray: []int{1, 3, 4, 6, 7, 9, 10, 11, 13},
				lookingFor:  6,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinarySearch(tt.args.sortedArray, tt.args.lookingFor); got != tt.want {
				t.Errorf("BinarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
