package main

import (
	"log"
	"fmt"
	"errors"
	"math/rand"
)

func callbackCatch(cfg *config, args ...string) error {
    
   if len(args)!=1 {
	return errors.New("No pokemon name provided")
   }
   pokemonName:=args[0]

	pokemon, err :=cfg.pokeapiClients.GetPokemon(pokemonName)
	if err!=nil{
		log.Fatal(err)
	}

   const treshHold =50
   
	randNum:= rand.Intn(pokemon.BaseExperience)
	fmt.Println(pokemon.BaseExperience, randNum, treshHold)
	if randNum>treshHold {
		return fmt.Errorf("fail to catch %s!\n", pokemonName)
	}

  fmt.Printf("was caugth %s:\n",  pokemonName)
  
  return nil
}