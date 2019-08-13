package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

func main() {
	counter := make(map[string]int)
	total := 0

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("failed opening file: %s", err)

	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		//re := regexp.MustCompile(`.*GET\s(.*)\sHTTP.*\s-\s-$`)
		re := regexp.MustCompile(`.*GET\s(\/.*(\.js|\.gif|\.png|\.swf)).*\sHTTP.*$`)
		rep := "$1"

		if re.MatchString(line) {
			uri := re.ReplaceAllString(line, rep)
			counter[uri]++
		}
	}

	for k, v := range counter {
		fmt.Printf("%s: %d\n", k, v)
		total += v
	}

	fmt.Println("Total:", total)

}
