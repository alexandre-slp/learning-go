package main

import (
	"github.com/alexandre-slp/learning-go/mocking/service"
	"testing"
)

var speakMock func(msg string) string

type spMock struct{}

func (spMock) Speak(msg string) string {
	return "hello ale 123"
}

func Test_doSomething(t *testing.T) {
	t.Parallel()

	type args struct {
		msg     string
		spkServ service.Speaker
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{msg: "ale", spkServ: service.Sp{}},
			want: "hello ale",
		},
		{
			name: "2",
			args: args{msg: "ale", spkServ: spMock{}},
			want: "hello ale 123",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := doSomething(tt.args.msg, tt.args.spkServ); got != tt.want {
				t.Errorf("doSomething() = %v, want %v", got, tt.want)
			}
		})
	}
}
