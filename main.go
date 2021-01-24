package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type bodyData struct {
	Weight float64
	Fat    float64
	Bmi    float64
	Date   string
	Time   string
	LogID  float64
	Source string
}

const (
	dataLimit  = 295
	poundsToKg = 0.4535924
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Target dir required.")
	}
	targetFolder := os.Args[1]

	infos, err := ioutil.ReadDir(targetFolder)
	if err != nil {
		log.Fatal("File reading error:", err)
	}

	targetFiles := []os.FileInfo{}
	for _, f := range infos {
		if strings.HasSuffix(f.Name(), ".json") && strings.Contains(f.Name(), "weight") {
			targetFiles = append(targetFiles, f)
		}
	}

	fmt.Println("weight files to transform:", len(targetFiles))
	var filteredData []bodyData
	for _, f := range targetFiles {
		if !f.IsDir() {
			absPath := filepath.Join(targetFolder, f.Name())
			file, _ := os.Open(absPath)
			defer file.Close()
			decoder := json.NewDecoder(file)

			_, err := decoder.Token()
			if err != nil {
				log.Fatal(err)
			}
			for decoder.More() {
				var x bodyData
				err := decoder.Decode(&x)
				if err != nil {
					log.Fatal(err)
				}
				filteredData = append(filteredData, x)
			}
			_, err = decoder.Token()
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	var chunks [][]bodyData
	for i := 0; i < len(filteredData); i += dataLimit {
		end := i + dataLimit
		if end > len(filteredData) {
			end = len(filteredData)
		}
		s := filteredData[i:end]
		chunks = append(chunks, s)
	}
	for i, c := range chunks {
		csvFile, _ := os.Create(fmt.Sprintf("heracles%d.csv", i))
		writer := csv.NewWriter(csvFile)
		defer csvFile.Close()

		var lines [][]string
		for _, d := range c {
			kg := math.Round(d.Weight*poundsToKg*100) / 100
			fatKg := kg / 100 * d.Fat
			layout := "01/02/06 15:04:05"
			y := fmt.Sprintf("%s %s", d.Date, d.Time)
			t, _ := time.Parse(layout, y)

			line := []string{t.Format("2006-01-02 15:04:05"), fmt.Sprintf("%.2f", kg)}
			if fatKg != 0 {
				fat := fmt.Sprintf("%.2f", fatKg)
				line = append(line, fat)
			}
			lines = append(lines, line)
		}
		writer.WriteAll(lines)
	}
}
