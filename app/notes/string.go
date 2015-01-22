package notes

type String struct {
	Render
	str string
}

func NewString(n Note, str string) String {
	return String{
		Render: NewRender(n.Type(), n.Controller()),
		str:    str,
	}
}

func (s String) String() string {
	return s.str
}
