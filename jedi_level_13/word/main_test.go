package word

import (
	"fmt"
	"github.com/alexandre-slp/learning-go/jedi_level_13/quote"
	"reflect"
	"testing"
)

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count(quote.SunAlso)
	}
}

func ExampleCount() {
	fmt.Println(Count("1 2 3"))
	// Output:
	// 3
}

func TestCount(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "with words",
			args: args{s: "abc def g h i"},
			want: 5,
		},
		{
			name: "only spaces",
			args: args{s: "   "},
			want: 0,
		},
		{
			name: "empty",
			args: args{s: ""},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Count(tt.args.s); got != tt.want {
				t.Errorf("Count() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount(quote.SunAlso)
	}
}

func ExampleUseCount() {
	fmt.Println(UseCount("1 2 2 3 3 3"))
	// Output:
	// map[1:1 2:2 3:3]
}

func TestUseCount(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want map[string]int
	}{
		{
			name: "with words",
			args: args{s: "1 2 2 3 3 3"},
			want: map[string]int{
				"1": 1,
				"2": 2,
				"3": 3,
			},
		},
		{
			name: "only spaces",
			args: args{s: "    "},
			want: map[string]int{},
		},
		{
			name: "empty",
			args: args{s: ""},
			want: map[string]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UseCount(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UseCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
