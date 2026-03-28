package main

import (
	"fmt"
	"regresionlineal/calculos"
	"regresionlineal/data"
)

func main() {
	//Programa para predecir la temperatura en base a relative_humidity,precipitation,pressure
	csvData := data.ReadCsv()
	for key, value := range csvData {
		fmt.Printf("Cantidad de datos de %v: %v\n", key, len(value))
	}
	fmt.Printf("%v\n", "-----------")
	mediaTemperature := calculos.CalcularMedia(csvData["temperature"])
	fmt.Printf("Media temperatura: %v\n", mediaTemperature)
	mediaRelativeHumidity := calculos.CalcularMedia(csvData["relative_humidity"])
	fmt.Printf("Media humedad relativa: %v\n", mediaRelativeHumidity)
	mediaPrecipitation := calculos.CalcularMedia(csvData["precipitation"])
	fmt.Printf("Media precipitacion: %v\n", mediaPrecipitation)
	mediaPressure := calculos.CalcularMedia(csvData["pressure"])
	fmt.Printf("Media pressure: %v\n", mediaPressure)
	desviacionTemperature := calculos.CalcularDesviacionEstandar(csvData["temperature"], mediaTemperature)
	fmt.Printf("%v\n", "-----------")
	fmt.Printf("Desviacion temperatura: %v\n", desviacionTemperature)
	desviacionRelativeHumidity := calculos.CalcularDesviacionEstandar(csvData["relative_humidity"], mediaRelativeHumidity)
	fmt.Printf("Desviacion humedad relativa: %v\n", desviacionRelativeHumidity)
	desviacionPrecipitation := calculos.CalcularDesviacionEstandar(csvData["precipitation"], mediaPrecipitation)
	fmt.Printf("Desviacion precipitacion: %v\n", desviacionPrecipitation)
	desviacionPressure := calculos.CalcularDesviacionEstandar(csvData["pressure"], mediaPressure)
	fmt.Printf("Desviacion presion: %v\n", desviacionPressure)
	fmt.Printf("%v\n", "-----------")
	//Sesgada a la derecha
	asimetriaTemperature := calculos.CalcularAsimetria(csvData["temperature"], mediaTemperature, desviacionTemperature)
	fmt.Printf("Asimetria de la temperatura: %v\n", asimetriaTemperature)
	//Sesgada a la izquierda
	asimetriaRelativeHumidity := calculos.CalcularAsimetria(csvData["relative_humidity"], mediaRelativeHumidity, desviacionRelativeHumidity)
	fmt.Printf("Asimetria humedad relativa: %v\n", asimetriaRelativeHumidity)
	//Sesgada a la derecha
	asimetriaPrecipitation := calculos.CalcularAsimetria(csvData["precipitation"], mediaPrecipitation, desviacionPrecipitation)
	fmt.Printf("Asimetria precipitacion: %v\n", asimetriaPrecipitation)
	//Casi distribucion normal
	asimetriaPressure := calculos.CalcularAsimetria(csvData["pressure"], mediaPressure, desviacionPressure)
	fmt.Printf("Asimetria presion: %v\n", asimetriaPressure)
	fmt.Printf("%v\n", "-----------")
	//Debido a que no todos los datos siguen una distribucion normal se normaliza por min max scaler
	normalizeData := make(map[string][]float64)
	x_min_max := make(map[string]map[string]float64)
	for key, slice := range csvData {
		x_min, x_max := slice[0], slice[0]
		x_min_max[key] = make(map[string]float64)
		for _, value := range slice {
			if value < x_min {
				x_min = value
				x_min_max[key]["x_min"] = x_min
			}
			if value > x_max {
				x_max = value
				x_min_max[key]["x_max"] = x_max
			}
		}
		//Itera segun el numero de filas en cada columna
		for i := 0; i < len(slice); i++ {
			minMax := calculos.MinMaxScaler(slice[i], x_min, x_max)
			normalizeData[key] = append(normalizeData[key], minMax)
		}
	}
	for key, value := range x_min_max {
		fmt.Printf("%v: X_min:%v X_max: %v\n", key, value["x_min"], value["x_max"])
	}
	fmt.Printf("%v\n", "-----------")
	//Configuraciones iniciales
	targetY := normalizeData["temperature"]
	muestras := len(targetY)
	featuresX := make([][]float64, muestras)
	for i := 0; i < muestras; i++ {
		fila := []float64{
			normalizeData["relative_humidity"][i],
			normalizeData["precipitation"][i],
			normalizeData["pressure"][i],
		}
		featuresX[i] = fila
	}
	fmt.Printf("%v\n", featuresX)
	fmt.Printf("%v\n", "-----------")
	numFeatures := len(featuresX[0])
	iteraciones := 1000
	aprendizaje := calculos.Entrenamiento(iteraciones, numFeatures, muestras, targetY, featuresX)
	fmt.Printf("%v\n", "-----------")
	fmt.Printf("%v\n", aprendizaje)
	fmt.Printf("%v\n", "-----------")
	var relative_humidity, precipitation, pressure float64
	fmt.Print("Ingresa la humedad relativa: \n")
	fmt.Scan(&relative_humidity)
	fmt.Print("Ingresa la precipitacion: \n")
	fmt.Scan(&precipitation)
	fmt.Print("Ingresa la presion a nivel del mar: \n")
	fmt.Scan(&pressure)
	temperatura := calculos.PredecirTemperatura(relative_humidity, precipitation, pressure, aprendizaje, x_min_max)
	fmt.Printf("%v\n", "-----------")
	fmt.Printf("Temperatura predicha por la regresion lineal: %v\n", temperatura)
	fmt.Printf("%v\n", "-----------")
}
