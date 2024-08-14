package main

import (
	"crypto/sha256"
	"fmt"
	"os"
)

func HashFiles() string {
	s := entriesInIndex()

	file, err := os.Create("sys-temp.txt")
	check(err)

	for i := 0; i < len(s)-1; i++ {
		dataFromSrc, err := os.ReadFile(s[i])
		check(err)
		_, err = fmt.Fprint(file, dataFromSrc)
		check(err)

	}

	dataFromTemp, err := os.ReadFile("sys-temp.txt")
	check(err)

	sha256Hash := sha256.New()
	sha256Hash.Write(dataFromTemp)
	return fmt.Sprintf("%x", sha256Hash.Sum(nil))
}
