package command

func attach(p *Parser) {
	p.RegisterCommands(
		CreateTask,
		SelectProject,
		GetTasks,
		CloseTask,
	)
}
