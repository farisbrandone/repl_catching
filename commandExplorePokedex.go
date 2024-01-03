package main

import (
	"log"
	"fmt"
	"errors"
)

func callbackExplorePokedex(cfg *config, args ...string) error {
    
   if len(args)!=1 {
	return errors.New("No location area provided")
   }
   locationAreaName:=args[0]

	if cfg.totalCount!=nil {
		a:= *cfg.totalCount
		if cfg.count >=a{
		   return errors.New("you are in the last page")
	   } 
	}
 


	locationArea, err :=cfg.pokeapiClients.GetLocationArea(locationAreaName)
	if err!=nil{
		log.Fatal(err)
	}

  fmt.Printf("Pokemon in %s:\n", locationArea.Name)
  for _, pokemon:=range locationArea.PokemonEncounters{
      fmt.Printf("- %s\n", pokemon.Pokemon.Name)
  }
  

  
  return nil
}