package template 

type ExampleTemplate struct{}

func (e *ExampleTemplate) init() string {
	return "Example Init Template"
}

func (e *ExampleTemplate) repository() string {
	return "Example Repository Template"
}

func (e *ExampleTemplate) model() string {
	return "Example Model Template"
}
