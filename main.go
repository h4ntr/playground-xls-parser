package main

import (
    "encoding/csv"
    "fmt"
    "os"
    "log"
    "io"
    //"bufio"
    "github.com/ctessum/macreader"
    "strings"
)

const sSourceFileName = "./test.csv"

func parseHistory(s string) {
    ss := strings.Split(s,",")
    for _, sChange := range ss {
        sChange = strings.TrimSpace(sChange)
        fmt.Println("-------->",sChange)
    }
}

func main(){

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
/*
    scanner := bufio.NewScanner(f)
    for scanner.Scan(){
        s := scanner.Text()
        fmt.Println(s[0:30])
    }
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
*/

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
                    parseHistory(line[col])
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
