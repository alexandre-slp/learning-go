package main

import (
	"github.com/alexandre-slp/learning-go/mockgen/mocks"
	"github.com/golang/mock/gomock"
	"testing"
)

func Test_2_doSomething(t *testing.T) {
	t.Parallel()
	mockCtrl := gomock.Controller{T: t}
	defer mockCtrl.Finish()

	type args struct {
		msg     string
	}
	tests := []struct {
		name string
		args args
		mock *mocks.MockSpeaker
		want string
	}{
		{
			name: "1",
			args: args{msg: "ale"},
			mock: mocks.NewMockSpeaker(&mockCtrl),
			want: "hello ale",
		},
		{
			name: "2",
			args: args{msg: "ale"},
			want: "hello ale 123",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock.EXPECT().Speak().Return().
			if got := doSomething(tt.args.msg); got != tt.want {
				t.Errorf("doSomething() = %v, want %v", got, tt.want)
			}
		})
	}
}
