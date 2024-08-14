package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func ChoiceCheckout() {
	hashSum := os.Args[2]
	data, err := os.ReadFile("./vcs/log.txt")
	check(err)
	dirPath := filepath.Join("./vcs/commits", hashSum)
	s := entriesInIndex()

	if !strings.Contains(string(data), hashSum) {
		fmt.Println("Commit does not exist.")
	} else {
		for i := 0; i < len(s)-1; i++ {
			fileName := filepath.Join(dirPath, s[i])

			dataFromArch, err := os.ReadFile(fileName)
			check(err)

			err = os.WriteFile(s[i], dataFromArch, 0666)
			check(err)
			/*err = os.Remove(s[i])
			check(err)
			file, err := os.Create(s[i])
			check(err)
			defer file.Close()
			_, err = fmt.Fprint(file, string(dataFromArch))
			check(err)*/
		}
		fmt.Printf("Switched to commit %s.", hashSum)
	}
}

func ChoiceLog() {
	dataFromLog := ReadLog()
	if len(dataFromLog) == 0 {
		fmt.Println("No commits yet.")
	} else {
		fmt.Print(string(dataFromLog))
	}
}

func ChoiceCommit() {
	hashSum := HashFiles()
	data, err := os.ReadFile("./vcs/log.txt")
	check(err)

	//checking changes
	if strings.Contains(string(data), hashSum) {
		fmt.Println("Nothing to commit.")
	} else {
		//creating dir and files
		dirPath := filepath.Join("./vcs/commits", hashSum)
		os.MkdirAll(dirPath, os.ModePerm)

		s := entriesInIndex()

		for i := 0; i < len(s)-1; i++ {
			fileName := filepath.Join(dirPath, s[i])
			file, err := os.Create(fileName)
			check(err)

			dataFromSrc, err := os.ReadFile(s[i])
			check(err)

			_, err = fmt.Fprint(file, string(dataFromSrc))
			check(err)
		}

		//writing into log
		dataConfig, err := os.ReadFile("./vcs/config.txt")
		check(err)

		logEntry := "commit " + hashSum + "\nAuthor: " + strings.TrimSpace(string(dataConfig)) + "\n" + os.Args[2]

		dataFromLog := ReadLog()

		logEntry = logEntry + "\n\n" + string(dataFromLog)
		err = os.WriteFile("./vcs/log.txt", []byte(logEntry), 0666)
		check(err)

		fmt.Println("Changes are committed.")
	}
}

func ChoiceConfig() {
	data, err := os.ReadFile("./vcs/config.txt")
	check(err)

	if len(os.Args) == 3 {
		err := os.WriteFile("./vcs/config.txt", []byte(os.Args[2]), 0644)
		check(err)
		fmt.Printf("The username is %s.\n", os.Args[2])
	} else if len(data) == 0 {
		fmt.Println("Please, tell me who you are.")
	} else {
		fmt.Printf("The username is %s.\n", data)
	}
}
func ChoiceAdd() {
	if len(os.Args) == 2 {
		data, err := os.ReadFile("./vcs/index.txt")
		check(err)
		if len(data) == 0 {
			fmt.Println("Add a file to the index.")
		} else {
			fmt.Printf("Tracked files:\n%s", data)
		}
	} else {
		file, err := os.Open(os.Args[2])
		if err != nil {
			fmt.Printf("Can't find '%s'.", os.Args[2])
		} else {
			file.Close()
			file, err = os.OpenFile("./vcs/index.txt", os.O_RDWR|os.O_APPEND, 0644)
			check(err)
			_, err = fmt.Fprint(file, os.Args[2]+"\n")
			check(err)
			fmt.Printf("The file '%s' is tracked.\n", os.Args[2])
		}
	}
}
