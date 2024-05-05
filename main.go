package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/biter777/countries"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintln(os.Stdout, "Usage: ibangen [country code]")
		printSupportedCountries()
	}

	flag.Parse()

	if countryCode := flag.Arg(0); countryCode != "" {
		countryCode = strings.ToUpper(countryCode)
		generator, ok := Generators[countryCode]
		if !ok {
			fmt.Printf("Could not find generator for country code '%s'.\n", strings.ToLower(countryCode))
			printSupportedCountries()
			return
		}

		fmt.Println(generator.Generate())
		return
	}

	var seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()),
	)

	i := seededRand.Intn(len(Generators))
	for _, generator := range Generators {
		if i == 0 {
			fmt.Println(generator.Generate())
			break
		}

		i -= 1
	}
}

func printSupportedCountries() {
	fmt.Println("Supported country codes:")
	for countryCode := range Generators {
		name := countries.ByName(countryCode).String()
		fmt.Printf("- '%s' (%s)\n", strings.ToLower(countryCode), name)
	}
}
