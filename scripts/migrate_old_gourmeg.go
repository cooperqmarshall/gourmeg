// Migrates the date from the original gourmeg to the new
// gourmeg. This requires a csv with two columns: list_name
// and url. This script will iterate through each row and add
// the list if that name has not been added and add the recipe
// to that list. If another row has the same list name, that
// recipe will be added to the same list.
package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
)

func main() {
	file, err := os.Open("data.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	csv_reader := csv.NewReader(file)
	records, err := csv_reader.ReadAll()

    lists := make(map[string]string)

    for i, record := range records {
		if i == 0 {
			continue
		}
		if val, ok := lists[record[0]]; val == "" || !ok {
            fmt.Printf("adding new list: %s\n", record[0])
            req_body := bytes.NewBuffer([]byte(fmt.Appendf([]byte("name="), record[0])))
            res, err := http.Post("https://gourmeg.org/list?parent_id=0", "application/x-www-form-urlencoded", req_body)
            if err != nil {
		        panic(err)
            }

            body, err := io.ReadAll(res.Body)
            if err != nil {
		        panic(err)
            }

            re := regexp.MustCompile(`href="/list/(.*)"`)
            list_id := re.FindSubmatch(body)[1]
            lists[record[0]] = string(list_id)
            fmt.Printf("added new list: %s\n", list_id)
        }

        // add recipe
        fmt.Printf("Adding recipe (%s) to list %s\n", record[1], lists[record[0]])
        fmt.Printf("https://gourmeg.org/recipe?list_id=%s&ignore_duplicates=true\n", lists[record[0]])
        req_body := bytes.NewBuffer([]byte(fmt.Appendf([]byte("url="), record[1])))
        res, err := http.Post(fmt.Sprintf("https://gourmeg.org/recipe?list_id=%s&ignore_duplicates=true", lists[record[0]]),
                                    "application/x-www-form-urlencoded",
                                    req_body)
        if err != nil {
            panic(err)
        }
        fmt.Printf("%s\n",res.Status)
    }
}
