package template

type PHPTemplate struct{}

func (p *PHPTemplate) init() string {
	return `<?php
		echo "hello World!";
	`
}

func (p *PHPTemplate) repository() string {
	return "PHP Repository"
}

func (p *PHPTemplate) model() string {
	return "PHP Model"
}
