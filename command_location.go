package main

import "fmt"

func commandLocation(cfg *config, params ...string) error {
	if len(params) == 0 {
		location, err := cfg.apiClient.LocationInfo()
		if err != nil {
			return err
		}

		fmt.Println("You are at", location.Name)
		fmt.Println()
		// TODO: add some more information about the location

		fmt.Println("You can move using 'location next' or 'location previous' commands")
		fmt.Println("Previous:", location.Previous, "| Next:", location.Next)

		return nil
	}

	switch params[0] {
	case "next", "previous":
		location, err := cfg.apiClient.Move(params[0])
		if err != nil {
			return err
		}

		fmt.Println("You are now at", location.Name)
		fmt.Println()
		fmt.Println("Previous:", location.Previous, "| Next:", location.Next)
		return nil

	case "search":
		pokemon, err := cfg.apiClient.Search()
		if err != nil {
			return err
		}

		fmt.Println("You found a wild <lvl", pokemon.Level, ">", pokemon.Name)
		fmt.Println()
		fmt.Println("HP:", pokemon.HP)
		fmt.Println("Attack:", pokemon.Attack)
		fmt.Println("Defense:", pokemon.Defense)
		fmt.Println("Special Attack:", pokemon.SpecialAttack)
		fmt.Println("Special Defense:", pokemon.SpecialDefense)
		fmt.Println("Speed:", pokemon.Speed)
		return nil

	default:
		return fmt.Errorf("invalid parameter. Use 'next' or 'previous' to move, or 'search' to find wild pokemon")
	}
}
