package main
  
import (
	"fmt"
)

func callbackHelp(cfg *config, args ...string) error {
	fmt.Println("Welcome to the help command")
	fmt.Println("Here are the available command:")

    availableGetCommand:=getCliCommand()
	for _,cmd:=range availableGetCommand{
		fmt.Printf("- %s: %s\n", cmd.name, cmd.description)
	}
	return nil

}