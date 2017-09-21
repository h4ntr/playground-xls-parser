package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"log"
	"io"
	"github.com/ctessum/macreader"
	"strings"
	"regexp"
)

const sSourceFileName = "./test.csv"

func ParseToken(s string){
	const sBucksExp = `^\$([\d](.\d){0,1})k$`
	var BucksExp = regexp.MustCompile(sBucksExp)
	var doesMatch = BucksExp.MatchString(s)
	if doesMatch {
		fmt.Printf("Token \"%v\" matches `%v` = %v\n", s, sBucksExp, doesMatch)
		fmt.Printf("Submatch: \"%q\"\n\n", BucksExp.FindAllStringSubmatch(s,-1))
	}
}

func ParseChangeRecord(s string){
	ss := strings.Split(s, " ")
	for _, sToken := range ss{
		ParseToken(strings.Trim(sToken," \t"))
	}
}

func ParseHistory(s string) {
	ss := strings.Split(s, ";")
	for _, sChgRecord := range ss {
		ParseChangeRecord(sChgRecord)
	}
}

func main() {

	fmt.Println(">>I'll try to do a little bit of parsing for you now...")

	// I can use Cobra (https://github.com/spf13/cobra) to create a nice CLI later.
	// A simpler implementation may be achieved with package "flag"

	// Open file and pass it to the csv parser.
	// Split lines into fields, one of them History
	f, err := os.Open(sSourceFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	csvReader := csv.NewReader(macreader.New(f))

	var headers []string
	for lineNum := 1; ; lineNum++ {

		line, err := csvReader.Read()
		if err == io.EOF {
			fmt.Println(">>> EOF reached.")
			break
		} else if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("\nLine %v:\n", lineNum)

		if lineNum == 1 {
			headers = line
		}

		// No error checks starting from here. Need to be added

		for col := 0; col < len(line); col++ {
			if lineNum == 1 {
				fmt.Printf("\t%v\n", headers[col])
			} else {
				fmt.Printf("\t%v: \"%v\"\n", headers[col], line[col])
				if headers[col] == "History" {
					ParseHistory(line[col])
				}
			}
		}
	}

	// Parse History fields using regular experssions recgonizing
	// change dates, sums, currencies, and titles

	// Save it. Where?
	// Join with historical exchanges rates. How?
	// Build charts. How?
	// Add filters, dimensions. Which and how?
}
