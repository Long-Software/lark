package remove

type RemoveArgs struct {
	Page int
	Start int
	End int
}
type RemoveCommand struct{}

func (r *RemoveCommand) Run() {
	// TODO: create a remove commnad to remove either a single page in the pdf or multiple page from start to end
}

func (r *RemoveCommand) Help() {
	// TODO: implement RemoveCommand help
	panic("TODO: Implement")
}
