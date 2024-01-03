package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func startRepl(cfg *config){
	scanner:=bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("pokedexcli >")
		scanner.Scan()
		text:=scanner.Text()
		cleaned:=cleanInput(text)
		if len(cleaned)==0{
			continue
		}
		commandName:=cleaned[0]
		args:=[]string{}
		if len(cleaned)>1{
			args=cleaned[1:]
		}
		availableGetCommand:=getCliCommand()
		command, ok:=availableGetCommand[commandName]
		if !ok {
			fmt.Println("Invalid command")	
			continue
		}
		err:=command.callback(cfg, args...)
		if err!=nil {
			fmt.Println(err)
		}
		
	}
	
}

func cleanInput(s string) []string {
     lowercase:=strings.ToLower(s)
	 words:=strings.Fields(lowercase) //equivalent to split()
	 return words
}

type cliCommand struct {
	name string
	description string
	callback func(*config, ...string) error
}

func getCliCommand() map[string]cliCommand{
	return map[string]cliCommand{
		"help":{
			name: "help",
			description:"give all command available",
			callback:callbackHelp,
		},
		"map":{
			name: "map",
			description:"give successif 20 location area",
			callback:callbackMap,
		},
		"mapb":{
			name: "mapb",
			description:"give successif 20 location area",
			callback:callbackMapB,
		},
        "explore":{
			name: "explore {location_area}",
			description:"list the pokemon in location area",
			callback:callbackExplorePokedex,
		},
		"catch":{
			name: "catch {pokemon_name}",
			description:"list catch area",
			callback:callbackCatch,
		},
		"exit":{
			name:"exit",
			description:"turn off pokedex",
			callback:callbackExit,
		},
	}
}

