package main

import "testing"

func Test_test(t *testing.T) {
	t.Parallel()

	type args struct {
		d int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "d < 0",
			args: args{d: -2},
			want: 42,
		},
		{
			name: "d > 0",
			args: args{d: 4},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := test(tt.args.d); got != tt.want {
				t.Errorf("test() = %v, want %v", got, tt.want)
			}
		})
	}
}
