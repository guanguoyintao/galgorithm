package string_matching

import "testing"

func TestBFSearch(t *testing.T) {
	type args struct {
		s    string
		temp string
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 bool
	}{
		{
			name: "1",
			args: args{
				s:    "substring search algorithm",
				temp: "search",
			},
			want:  10,
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := BFSearch(tt.args.s, tt.args.temp)
			if got != tt.want {
				t.Errorf("BFSearch() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("BFSearch() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
