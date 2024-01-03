package main

import (
	"time"
	"github.com/farisbrandone/repl_catching/internal/pokeapi"
)

type config struct {
	pokeapiClients pokeapi.Client
	nextLocationUrlArea *string //peut ne pas etre initialiser car possiblement nil
	prevLocationUrlArea *string
	count int
	totalCount *int
	caugthPokemon map[string]pokeapi.Pokemon
}


func main(){
   cfg :=config{
	pokeapiClients: pokeapi.NewClient(time.Hour),
	count:20,
	caugthPokemon: make(map[string]pokeapi.Pokemon),

   }
	startRepl(&cfg)
}