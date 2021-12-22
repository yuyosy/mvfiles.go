package main

import (
	"flag"
	"fmt"
	"log"
	"mvfiles/mvfiles"
	"path/filepath"
)

var (
	version        string = "1.0"
	build_datetime string = ""
)

func main() {
	fmt.Printf("mvfiles Ver.%s (Build:%s)\n", version, build_datetime)
	config := flag.String("c", "moverules.csv", "Config")
	quit := flag.Bool("q", false, "Exit without pressing any key.")
	flag.Parse()
	rules, err := mvfiles.MoveRules(*config)
	if err != nil {
		log.Fatal(err)
		// Exit
	}
	for _, rule := range rules {
		if err := mvfiles.MakeDirs(rule.Moveto); err != nil {
			fmt.Println(err)
			continue
		}
		maches, err := filepath.Glob(rule.Pattern)
		if err != nil {
			fmt.Println(err)
			continue
		}
		mvfiles.MoveFiles(maches, rule.Moveto)
		fmt.Println()
	}
	if !*quit {
		fmt.Print("Press any key to exit...")
		fmt.Scanln()
	}
}
