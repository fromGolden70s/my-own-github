package main

import (
	"fmt"
	"os"
)

func main() {
	DirectoryCreator()

	if len(os.Args) == 1 {
		os.Args = append(os.Args, "--help")
	}

	ArgsHandler(os.Args[1])
}
func ArgsHandler(arg string) {

	switch arg {
	case "checkout":
		if len(os.Args) == 2 {
			fmt.Println("Commit id was not passed.")
		} else {
			ChoiceCheckout()
		}
	case "config":
		ChoiceConfig()
	case "add":
		ChoiceAdd()
	case "commit":
		if len(os.Args) == 2 {
			fmt.Println("Message was not passed.")
		} else {
			ChoiceCommit()
		}
	case "log":
		ChoiceLog()
	case "--help":
		fmt.Println("These are SVCS commands:\n" +
			"config     Get and set a username.\n" +
			"add        Add a file to the index.\n" +
			"log        Show commit logs.\n" +
			"commit     Save changes.\n" +
			"checkout   Restore a file.")
	default:
		fmt.Printf("'%s' is not a SVCS command.", arg)
	}
}
