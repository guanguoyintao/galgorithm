package string_matching

import "testing"

func TestSundaySearch(t *testing.T) {
	type args struct {
		base    string
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
				base:    "substring search algorithm",
				pattern: "search",
			},
			want:  10,
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := SundaySearch(tt.args.base, tt.args.pattern)
			if got != tt.want {
				t.Errorf("SundaySearch() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("SundaySearch() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
