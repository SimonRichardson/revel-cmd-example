package controllers

import (
	"github.com/SimonRichardson/revel-cmd-example/app/cmds"
	"github.com/SimonRichardson/revel-cmd-example/app/notes"
	"github.com/revel/revel"
)

var (
	HelloRunner cmds.Runner = cmds.Runner{
		map[string][]cmds.Command{
			"Index": []cmds.Command{
				cmds.Hello{},
				cmds.World{},
				cmds.RenderString{},
			},
		},
	}
)

type Hello struct {
	App
}

func (c Hello) Index() revel.Result {
	note := notes.NewRender("Index", c)
	code, result := HelloRunner.Run(note)
	c.Response.Status = code
	return result
}
