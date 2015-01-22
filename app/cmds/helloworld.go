package cmds

import "github.com/SimonRichardson/revel-cmd-example/app/notes"

type Hello struct{}

func (c Hello) Execute(note notes.Note) Result {
	return OK(notes.NewString(note, "Hello"))
}

type World struct{}

func (c World) Execute(note notes.Note) Result {
	str := note.(notes.String)
	return OK(notes.NewString(note, str.String()+" World!"))
}
