package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// main with args
func main() {
	args := os.Args

	dest := "./testdata"
	if len(args) > 1 {
		dest = args[1]
	}
	dest, err := filepath.Abs(dest)
	if err != nil {
		fmt.Println("error getting absolute path", err)
		return
	}

	now := time.Now()
	logTime := now.Add(time.Duration(-3 * time.Hour * 24 * 30))

	// <dest>/test_logs/<year>/<month>/<day>/<log_<time>.log>
	// ensure month and day are 2 digits
	filepattern := `%s/test_logs/%d/%02d/%02d/log_%d.json`

	type dataTime struct {
		Id        string    `json:"id"`
		Timestamp time.Time `json:"timestamp"`
	}
	idx := 0
	// create logs for the last 3 months
	for logTime.Before(now) {
		filename := fmt.Sprintf(filepattern, dest, logTime.Year(), logTime.Month(), logTime.Day(), logTime.Hour())
		// ensure directory exists
		dir := filepath.Dir(filename)
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			fmt.Println("error creating directories", err)
			return
		}

		// log entries
		logEntries := make([]dataTime, 50)
		entryTime := logTime
		// for 1-50
		for i := 0; i < 50; i++ {
			logEntries[i] = dataTime{
				Id:        fmt.Sprintf("id%d", idx),
				Timestamp: entryTime,
			}
			entryTime = entryTime.Add(time.Second * 10)
		}
		// marshal the log entries to json
		// write to file
		jsonBytes, err := json.MarshalIndent(logEntries, "", "  ")
		if err != nil {
			fmt.Println("error marshalling json", err)
			return
		}
		// write the file

		err = os.WriteFile(filename, jsonBytes, 0644)
		if err != nil {
			fmt.Println("error writing file", err)
			return

		}

		logTime = logTime.Add(time.Hour * 4)
	}
	fmt.Println("done")
}
