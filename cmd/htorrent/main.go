package main

import "github.com/Voornaamenachternaam/htorrent/cmd/htorrent/cmd"

func main() {
	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
