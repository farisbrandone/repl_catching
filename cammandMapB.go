package main

import (
	"log"
	"fmt"
	"errors"
)

func callbackMapB(cfg *config, args ...string) error {
	//pokeapiClient:=pokeapi.NewClient() Recupere un nouveau client

if cfg.prevLocationUrlArea == nil{
	return errors.New("you are in the first page")
}


	resp, err :=cfg.pokeapiClients.ListLocationArea(cfg.prevLocationUrlArea)
	if err!=nil{
		log.Fatal(err)
	}

  fmt.Println("location area is:")
  for _, area:=range resp.Results{
      fmt.Printf("- %s\n", area.Name)
  }
  cfg.nextLocationUrlArea=resp.Next
  cfg.prevLocationUrlArea=resp.Previous
  cfg.count-=20
  return nil
}