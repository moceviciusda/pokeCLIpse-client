package main

import "os"

func commandExit(cfg *config, params ...string) error {
	defer os.Exit(0)
	return nil
}
