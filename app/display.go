package app

import (
	"fmt"

	"github.com/cevaris/ordered_map"
	"github.com/divjotarora/todo-cli/command"
)

// Display prints a list of Nameable instances to the console.
func Display(ctx *command.Context) {
	displayHelper(ctx.CurrentListing().Items(), "")
}

func displayHelper(items *ordered_map.OrderedMap, prefix string) {
	iter := items.IterFunc()

	var idx int
	for {
		kv, ok := iter()
		if !ok {
			break
		}

		item := kv.Value.(command.ListItem)
		fmt.Printf("%v%v. %v\n", prefix, idx, item.Name())
		idx++

		if children := item.Children(); children.Len() > 0 {
			displayHelper(children, prefix+"  ")
		}
	}
}
