package data

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func ReadCsv() map[string][]float64 {
	file, err := os.Open("temperature_data.csv")
	if err != nil {
		log.Fatalf("No se pudo acceder al archivo csv: %v\n", err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	record, err := reader.Read()
	headers := make([]string, 0)
	var tempData map[string][]float64
	tempData = make(map[string][]float64, 0)
	//Obtiene los headers
	for _, header := range record {
		headers = append(headers, header)
	}
	//Obtiene los valores de cada columna
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalf("No se pudo leer el archivo csv: %v\n", err)
		}
		for i, header := range headers {
			float, err := strconv.ParseFloat(record[i], 64)
			if err != nil {
				fmt.Printf("No se pudo parsear a float64: %v\n", err)
			}
			tempData[header] = append(tempData[header], float)
		}
	}
	return tempData
}
