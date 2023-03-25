package main

import (
	"fmt"
)

func callbackMap(cfg *config, args ...string) error {
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationAreasApi)
	if err != nil {
		return err
	}

	cfg.nextLocationAreasApi = resp.Next
	cfg.prevLocationAreasApi = resp.Previous
	// fmt.Println(resp)

	fmt.Println("Location areas: ")
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}

	return nil
}
