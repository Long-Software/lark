type Command interface {
	Execute()
}

type PrintCommand struct {
 Msg string
}
func (p *PrintCommand) Execute() {
 fmt.Println(p.Msg)
}
type Invoker struct {
 commands []Command
}
func (i *Invoker) AddCommand(c Command) {
 i.commands = append(i.commands, c)
}
func (i *Invoker) Run() {
 for _, cmd := range i.commands {
  cmd.Execute()
 }
}