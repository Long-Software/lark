package template

type Template interface {
	init() string
	repository() string
	model() string
}

func NewTemplate(extension string, templateType string) string {
	var template Template
	switch extension {
	case ".php":
		template = &PHPTemplate{}
	default:
		template = &ExampleTemplate{}
	}

	switch templateType {
	case "init":
		return template.init()
	case "model":
		return template.model()
	case "repo":
		return template.repository()
	default:
		return template.init()

	}
}
