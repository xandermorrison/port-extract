package main

import (
	"io"
	"os"
	"fmt"
	"log"
	"regexp"
	"strings"
)

func main() {
	path := os.Args[1]

	file, err := os.Open(path)
	if err != nil {
		log.Fatalln("Can't open file")
	}
	defer file.Close()

	contents, _ := io.ReadAll(file)

	re := regexp.MustCompile(`(?m)^(\d+)/`)
	matches := re.FindAllStringSubmatch(string(contents), -1)

	ports := make([]string, 0)
	for _, match := range matches {
		ports = append(ports, match[1])
	}

	fmt.Println(strings.Join(ports, ","))
}
