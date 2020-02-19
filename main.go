package main

import (
	"fmt"
	"github.com/sethvargo/go-diceware/diceware"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

const words = 6

func main() {
	rand.Seed(time.Now().UnixNano())
	var defaultSeparators = []string{
		"#",
		"-",
		"_",
		"|",
		"!",
	}

	separator, ok := os.LookupEnv("SEPARATOR")
	if !ok {
		separator = defaultSeparators[rand.Intn(len(defaultSeparators))]
	}

	seed, err := strconv.Atoi(os.Getenv("WORDS"))
	if err != nil {
		seed = words
	}

	passwordResult, err := diceware.Generate(seed)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(strings.Join(passwordResult, separator))
}
