package dog

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rand.Seed(time.Now().UnixNano())
		Years(rand.Intn(10))
	}
}

func ExampleYears() {
	fmt.Println(Years(1))
	// Output:
	// 7
}

func TestYears(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "0 years",
			args: args{n: 0},
			want: 0,
		},
		{
			name: "1 year",
			args: args{n: 1},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Years(tt.args.n); got != tt.want {
				t.Errorf("Years() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rand.Seed(time.Now().UnixNano())
		YearsTwo(rand.Intn(10))
	}
}

func ExampleYearsTwo() {
	fmt.Println(YearsTwo(1))
	// Output:
	// 7
}

func TestYearsTwo(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "0 years",
			args: args{n: 0},
			want: 0,
		},
		{
			name: "1 year",
			args: args{n: 1},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := YearsTwo(tt.args.n); got != tt.want {
				t.Errorf("YearsTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}
