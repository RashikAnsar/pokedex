package main

import (
	"errors"
	"fmt"
)

func callbackMapb(cfg *config, args ...string) error {
	if cfg.prevLocationAreasApi == nil {
		return errors.New("you're on the first page")
	}
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.prevLocationAreasApi)
	if err != nil {
		return err
	}

	cfg.nextLocationAreasApi = resp.Next
	cfg.prevLocationAreasApi = resp.Previous

	fmt.Println("Location areas: ")
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}

	return nil
}
