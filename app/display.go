package app

import (
	"fmt"

	"github.com/divjotarora/todo-cli/command"
	"github.com/elliotchance/orderedmap"
)

// Display prints a list of Nameable instances to the console.
func Display(ctx *command.Context) {
	displayHelper(ctx.CurrentListing().Items(), "")
}

func displayHelper(items *orderedmap.OrderedMap, prefix string) {
	for idx, item := 0, items.Front(); item != nil; idx, item = idx+1, item.Next() {
		item := item.Value.(command.ListItem)
		fmt.Printf("%v%v. %v\n", prefix, idx, item.Name())

		if children := item.Children(); children.Len() > 0 {
			displayHelper(children, prefix+"  ")
		}
	}
}
