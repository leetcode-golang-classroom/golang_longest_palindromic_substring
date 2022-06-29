package sol

import "testing"

func BenchmarkTest(b *testing.B) {
	name := "cbbd"
	for idx := 0; idx < b.N; idx++ {
		longestPalindrome(name)
	}
}
func Test_longestPalindrome(t *testing.T) {
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
			want: "bab",
		},
		{
			name: "cbbd",
			args: args{s: "cbbd"},
			want: "bb",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
