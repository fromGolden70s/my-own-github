package main

import (
	"log"
	"os"
	"strings"
)

func entriesInIndex() []string {
	data, err := os.ReadFile("./vcs/index.txt")
	check(err)
	return strings.Split(string(data), "\n")
}

func ReadLog() []byte {
	dataFromLog, err := os.ReadFile("./vcs/log.txt")
	check(err)
	return dataFromLog
}

func DirectoryCreator() {
	err := os.MkdirAll("./vcs", 0666)
	check(err)
	file, err := os.OpenFile("./vcs/config.txt", os.O_CREATE, 0666)
	check(err)
	file.Close()
	file, err = os.OpenFile("./vcs/index.txt", os.O_CREATE, 0666)
	check(err)
	file.Close()
	file, err = os.OpenFile("./vcs/log.txt", os.O_CREATE, 0666)
	check(err)
	file.Close()
	file, err = os.OpenFile("./tracked_file.txt", os.O_CREATE, 0666)
	check(err)
	file.Close()
	file, err = os.OpenFile("./untracked_file.txt", os.O_CREATE, 0666)
	check(err)
	file.Close()
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
