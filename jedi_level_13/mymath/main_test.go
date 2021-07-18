package mymath

import (
	"fmt"
	"testing"
)

func BenchmarkCenteredAvg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CenteredAvg([]int{1, 2, 3, 4, 5})
	}
}

func ExampleCenteredAvg() {
	fmt.Println(CenteredAvg([]int{2, 1, 1, 4, 4, 3}))
	// Output:
	// 2.5
}
func TestCenteredAvg(t *testing.T) {
	type args struct {
		xi []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "with numbers",
			args: args{xi: []int{1, 2, 3, 4}},
			want: 2.5,
		},
		{
			name: "repeated numbers",
			args: args{xi: []int{2, 1, 1, 4, 4, 3}},
			want: 2.5,
		},
		{
			name: "empty list",
			args: args{xi: []int{}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CenteredAvg(tt.args.xi); got != tt.want {
				t.Errorf("CenteredAvg() = %v, want %v", got, tt.want)
			}
		})
	}
}
