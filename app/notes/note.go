package notes

import "github.com/revel/revel"

type Renderer interface {
	RenderError(err error) revel.Result
	Render(extraRenderArgs ...interface{}) revel.Result
	RenderText(text string, objs ...interface{}) revel.Result
}

type Note interface {
	Controller() Renderer
	Type() string
}

type Render struct {
	typ  string
	cont Renderer
}

func NewRender(typ string, cont Renderer) Render {
	return Render{
		typ:  typ,
		cont: cont,
	}
}

func (r Render) Type() string {
	return r.typ
}

func (r Render) Controller() Renderer {
	return r.cont
}
