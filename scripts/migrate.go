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

            res, err := http.Post("http://localhost:1323/list?parent_id=0", "application/x-www-form-urlencoded", req_body)
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
        fmt.Printf("http://localhost:1323/recipe?list_id=%s&ignore_duplicates=true\n", lists[record[0]])
        req_body := bytes.NewBuffer([]byte(fmt.Sprintf("url=%s", record[1])))
        res, err := http.Post(fmt.Sprintf("http://localhost:1323/recipe?list_id=%s&ignore_duplicates=true", lists[record[0]]),
                                    "application/x-www-form-urlencoded",
                                    req_body)
        if err != nil {
            panic(err)
        }
        fmt.Printf("%s\n",res.Status)
    }
}
