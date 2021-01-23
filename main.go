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
	dataLimit  = 300
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
		fmt.Println(s)
		chunks = append(chunks, s)
	}
	for i, c := range chunks {
		csvFile, _ := os.Create(fmt.Sprintf("heracles%d.csv", i))
		writer := csv.NewWriter(csvFile)
		defer csvFile.Close()

		var l [][]string
		for _, d := range c {
			kg := math.Round(d.Weight*poundsToKg*100) / 100
			fatKg := kg / 100 * d.Fat
			l = append(l, []string{fmt.Sprintf("%s %s", d.Date, d.Time), fmt.Sprintf("%.2f", kg), fmt.Sprintf("%.2f", fatKg)})
		}
		writer.WriteAll(l)
	}
}
