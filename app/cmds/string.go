package cmds

import (
	"net/http"

	"github.com/SimonRichardson/revel-cmd-example/app/notes"
)

type RenderString struct{}

func (c RenderString) Execute(note notes.Note) Result {
	return Result{
		Continue: false,
		Result:   note.(notes.String).String(),
		Code:     http.StatusOK,
	}
}
