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

const defaultSeed = 6

func main() {
	rand.Seed(time.Now().UnixNano())
	var defaultSeparators = [] string {
		"#",
		"-",
		"_",
		"|",
		"!",
	}

	separator, ok := os.LookupEnv("SEPARATOR")
	if !ok{
		separator = defaultSeparators[rand.Intn(len(defaultSeparators))]
		log.Printf(">>> Using default separator %s", separator)
	}

	seed, err := strconv.Atoi(os.Getenv("SEED_RANGE"))
	if err != nil {
		seed = defaultSeed
		log.Printf(">>> Using default seed %d", defaultSeed)
	}

	passwordResult, err := diceware.Generate(seed)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(strings.Join(passwordResult, separator))
}
