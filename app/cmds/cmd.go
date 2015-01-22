package cmds

import (
	"errors"
	"net/http"

	"github.com/SimonRichardson/revel-cmd-example/app/notes"
	"github.com/revel/revel"
)

const (
	CommandFailureStatusCode int = http.StatusInternalServerError
)

type Result struct {
	Continue bool
	Result   string
	Code     int
	Note     notes.Note
}

func OK(note notes.Note) Result {
	return Result{
		Continue: true,
		Code:     http.StatusOK,
		Note:     note,
	}
}

type Command interface {
	Execute(note notes.Note) Result
}

type Sequential struct {
	Commands []Command
	Fail     func(results []Result) Result
	Complete func(results []Result) Result
}

func (c Sequential) Execute(note notes.Note) Result {
	result := Result{
		Continue: false,
		Code:     CommandFailureStatusCode,
		Note:     note,
	}

	numOfCommands := len(c.Commands)
	results := make([]Result, numOfCommands, numOfCommands)
	for k, v := range c.Commands {
		result = v.Execute(result.Note)
		results[k] = result

		if !result.Continue {
			if c.Fail != nil {
				return c.Fail(results[:k+1])
			}
			return result
		}
	}

	if c.Complete != nil {
		return c.Complete(results)
	}
	return result
}

type Runner struct {
	Commands map[string][]Command
}

func NewEmptyBucket(name string) map[string][]Command {
	m := map[string][]Command{}
	m[name] = nil
	return m
}

func NewEmptyCommands(name string, cmds []Command) map[string][]Command {
	m := map[string][]Command{}
	m[name] = cmds
	return m
}

func NewRunner(cmds map[string][]Command) Runner {
	return Runner{
		Commands: cmds,
	}
}

func (c Runner) RegisterCommand(t string, command Command) {
	if c.Commands[t] == nil {
		c.Commands[t] = []Command{
			command,
		}
	} else {
		c.Commands[t] = append(c.Commands[t], command)
	}
}

func (o Runner) Run(note notes.Note) (int, revel.Result) {
	ctrl := note.Controller()

	runner := Sequential{
		Commands: o.Commands[note.Type()],
	}
	result := runner.Execute(note)
	// Nothing meaningful was executed
	if result.Result == "" && result.Code == CommandFailureStatusCode {
		return CommandFailureStatusCode, ctrl.RenderError(errors.New("Exhausted Commands."))
	}
	// No result to process
	if result.Result == "" {
		return CommandFailureStatusCode, ctrl.RenderError(errors.New("No valid result."))
	}
	// Valid all the way down.
	return result.Code, note.Controller().RenderText(result.Result)
}
