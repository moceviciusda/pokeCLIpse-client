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
		fmt.Println("Next:", location.Next, "Previous:", location.Previous)

		return nil
	}

	direction := params[0]
	if direction != "next" && direction != "previous" {
		return fmt.Errorf("invalid direction. Use 'next' or 'previous'")
	}

	location, err := cfg.apiClient.Move(direction)
	if err != nil {
		return err
	}

	fmt.Println("You are now at", location.Name)

	return nil
}
