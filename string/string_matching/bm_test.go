package string_matching

import (
	"testing"
)

func TestBMSearch(t *testing.T) {
	type args struct {
		text    string
		pattern string
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
				text:    "substring search algorithm",
				pattern: "search",
			},
			want:  10,
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := BMSearch(tt.args.text, tt.args.pattern)
			if got != tt.want {
				t.Errorf("BMSearch() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("BMSearch() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
