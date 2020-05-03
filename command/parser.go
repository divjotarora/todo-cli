package command

// Parser represents a command parser that stores valid commands.
type Parser struct {
	commandMap map[string]Command
}

// NewParser creates a new parser.
func NewParser() *Parser {
	p := &Parser{
		commandMap: make(map[string]Command),
	}

	attach(p)
	return p
}

// Parse returns the Command for the given command name and an ok value to indicate whether or not the command was
// found.
func (p *Parser) Parse(name string) (Command, bool) {
	cmd, ok := p.commandMap[name]
	return cmd, ok
}

// RegisterCommands registers the given commands to the parser.
func (p *Parser) RegisterCommands(cmds ...Command) {
	for _, cmd := range cmds {
		p.commandMap[cmd.Name()] = cmd
	}
}

// Commands returns a list of commands that are recognized by the parser.
func (p *Parser) Commands() []Command {
	cmds := make([]Command, 0, len(p.commandMap))
	for _, cmd := range p.commandMap {
		cmds = append(cmds, cmd)
	}

	return cmds
}
