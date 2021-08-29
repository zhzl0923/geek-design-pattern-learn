package memento

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMemento(t *testing.T) {
	in := NewInputText()
	snapshot := NewSnapshot()

	tests := []struct {
		input string
		want  string
	}{
		{
			input: ":list",
			want:  "",
		},
		{
			input: "hello",
			want:  "",
		},
		{
			input: ":list",
			want:  "hello",
		},
		{
			input: "world",
			want:  "",
		},
		{
			input: ":list",
			want:  "helloworld",
		},
		{
			input: ":undo",
			want:  "",
		},
		{
			input: ":list",
			want:  "hello",
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			switch tt.input {
			case ":list":
				assert.Equal(t, tt.want, in.String())
			case ":undo":
				in.SetContent(snapshot.Pop())
			default:
				snapshot.Push(in.content)
				in.Append(tt.input)
			}
		})

	}
}
