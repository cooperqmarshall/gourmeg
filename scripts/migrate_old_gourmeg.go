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
	defer file.Close()
	if err != nil {
		panic(err)
	}
	csv_reader := csv.NewReader(file)
	records, err := csv_reader.ReadAll()

    lists := make(map[string]string)

    for _, record := range records {
        if lists[record[0]] == "" {
            fmt.Printf("adding new list: %s\n", record[0])
            req_body := bytes.NewBuffer([]byte(fmt.Sprintf("name=%s", record[0])))

            res, err := http.Post("https://new.gourmeg.org/list?parent_id=1", "application/x-www-form-urlencoded", req_body)
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
        fmt.Printf("https://new.gourmeg.org/recipe?list_id=%s&ignore_duplicates=true\n", lists[record[0]])
        req_body := bytes.NewBuffer([]byte(fmt.Sprintf("url=%s", record[1])))
        res, err := http.Post(fmt.Sprintf("https://new.gourmeg.org/recipe?list_id=%s&ignore_duplicates=true", lists[record[0]]),
                                    "application/x-www-form-urlencoded",
                                    req_body)
        if err != nil {
            panic(err)
        }
        fmt.Printf("%s\n",res.Status)
    }
}
