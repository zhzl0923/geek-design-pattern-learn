package memento

import (
	"container/list"
	"strings"
)

type InputText struct {
	content string
}

func NewInputText() *InputText {
	return &InputText{content: ""}
}

func (in *InputText) Append(input string) {
	in.content = strings.Join([]string{in.content, input}, "")
}

func (in *InputText) SetContent(content string) {
	in.content = content
}

func (in *InputText) String() string {
	return in.content
}

type Snapshot struct {
	snapshots *list.List
}

func NewSnapshot() *Snapshot {
	return &Snapshot{snapshots: list.New()}
}

func (s *Snapshot) Push(in string) {
	s.snapshots.PushBack(in)
}

func (s *Snapshot) Pop() string {
	e := s.snapshots.Back()
	if e != nil {
		s.snapshots.Remove(e)
		return e.Value.(string)
	}
	return ""
}
