package main

import (
	"encoding/csv"
	"io"
	"os"
)

func getQuizes(csvPath string) []quiz {
	file, err := os.Open(csvPath)
	check(err)
	defer file.Close()

	result := []quiz{}
	r := csv.NewReader(file)
	r.Read() // ignore header
	for record, err := r.Read(); err != io.EOF; record, err = r.Read() {
		if err == nil && len(record) > 4 {
			q := quiz{
				question:  record[0],
				correct:   record[1],
				incorrect: [3]string{record[2], record[3], record[4]},
			}
			result = append(result, q)
		}
	}
	return result
}
