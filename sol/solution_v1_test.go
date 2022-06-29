package sol

import "testing"

func BenchmarkTestV1(b *testing.B) {
	name := "cbbd"
	for idx := 0; idx < b.N; idx++ {
		longestPalindromeDP(name)
	}
}
func Test_longestPalindromeDP(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "aaaaa",
			args: args{s: "aaaaa"},
			want: "aaaaa",
		},
		{
			name: "babad",
			args: args{s: "babad"},
			want: "aba",
		},
		{
			name: "cbbd",
			args: args{s: "cbbd"},
			want: "bb",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindromeDP(tt.args.s); got != tt.want {
				t.Errorf("longestPalindromeDP() = %v, want %v", got, tt.want)
			}
		})
	}
}
