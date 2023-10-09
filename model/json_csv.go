package model

import (
	"convert-file/util"
	"encoding/csv"
	"encoding/json"
	"log"
	"os"
)

func CSV(path string) error {

	inputFile, err := os.Open(path)
	util.IfErrorPanic(err)
	defer inputFile.Close()

	csvReader := csv.NewReader(inputFile)
	result, err := csvReader.ReadAll()
	util.IfErrorPanic(err)

	data := make(map[string]interface{})
	var jsonData []interface{}

	n := len(result[0])
	for _, values := range result[1:] {
		for i := 0; i < n; i++ {
			data[result[0][i]] = values[i]
		}
		jsonData = append(jsonData, data)
	}

	output, err := os.Create("output.json")
	util.IfErrorPanic(err)
	defer output.Close()

	encoder := json.NewEncoder(output)
	if err := encoder.Encode(jsonData); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

type Data struct {
	jsonData map[string]interface{}
}

func Json(path string) error {
	inputFile, err := os.Open(path)
	util.IfErrorPanic(err)
	defer inputFile.Close()

	var jsonData []map[string]string
	decoder := json.NewDecoder(inputFile)
	if err := decoder.Decode(&jsonData); err != nil {
		log.Fatal(err)
		return err
	}

	outputFile, err := os.Create("output.csv")
	util.IfErrorPanic(err)
	defer outputFile.Close()

	csvWrite := csv.NewWriter(outputFile)
	defer csvWrite.Flush()

	var header []string
	for key, _ := range jsonData[0] {
		header = append(header, key)
	}
	if err := csvWrite.Write(header); err != nil {
		log.Fatal(err)
		return err
	}

	for _, csvData := range jsonData {
		record := []string{}

		for _, valueString := range csvData {
			record = append(record, valueString)
		}
		if err = csvWrite.Write(record); err != nil {
			log.Fatal(err)
		}
	}
	return nil
}
