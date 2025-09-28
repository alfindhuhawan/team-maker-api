package util

import (
	"encoding/csv"
	"log"
	"os"
)

func PrintCSV(data [][]string, folderFilename string) error {

	if _, err := os.Stat("./report"); os.IsNotExist(err) {
		err := os.Mkdir("./report", 0755)
		if err != nil {
			log.Fatal(err)
		}
	}

	csvFile, err := os.Create(folderFilename)

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	csvwriter := csv.NewWriter(csvFile)

	for _, dataRow := range data {
		_ = csvwriter.Write(dataRow)
	}
	csvwriter.Flush()
	csvFile.Close()

	return nil
}
