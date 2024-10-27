package hashMap

import "testing"

func Test_canPermutePalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{"addsa"},
			want: true,
		},
		{
			name: "1",
			args: args{"aabbddood"},
			want: true,
		},
		{
			name: "1",
			args: args{"aabb"},
			want: true,
		},
		{
			name: "1",
			args: args{"code"},
			want: false,
		},
		{
			name: "1",
			args: args{"carerac"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canPermutePalindrome(tt.args.s); got != tt.want {
				t.Errorf("canPermutePalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
