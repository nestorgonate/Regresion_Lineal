package main

import (
	"fmt"
	"regresionlineal/calculos"
	"regresionlineal/data"
	"regresionlineal/models"
)

func main() {
	//Programa para predecir la temperatura en base a relative_humidity,precipitation,pressure
	csvData := data.ReadCsv()
	for key, value := range csvData{
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
	for key, slice := range csvData {
		x_min, x_max := slice[0], slice[0]
		for _, value := range slice {
			if value < x_min {
				x_min = value
			}
			if value > x_max {
				x_max = value
			}
		}
		//Itera segun el numero de filas en cada columna
		for i := 0; i < len(slice); i++ {
			minMax := calculos.MinMaxScaler(slice[i], x_min, x_max)
			normalizeData[key] = append(normalizeData[key], minMax)
		}
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
	aprendizaje := models.Aprendizaje{
		Pesos:           []float64{0.1, 0.1, 0.1},
		Sesgo:           0.0,
		TasaAprendizaje: 0.01,
	}
	//Entrenamiento
	for i := 0; i < iteraciones; i++ {
		var errorAcumuladoSesgo float64 //Sumatoria de errores de prediccion
		errorAcumuladoPesos := make([]float64, numFeatures)
		var errorCuadratico float64
		//Itera j que es cada tiempo del clima
		for j := 0; j < muestras; j++ {
			//Regesion lineal simple: w*x+b
			prediccion := calculos.RegresionLineal(featuresX[j], aprendizaje.Pesos, aprendizaje.Sesgo)
			//Obtener error de prediccion
			errorPrediccion := prediccion - targetY[j]
			//Elevar el error de prediccion al cuadrado
			errorCuadratico += errorPrediccion * errorPrediccion
			//Sumatoria del error en la prediccion (Sumatoria del error)
			errorAcumuladoSesgo += errorPrediccion
			//Itera k que es cada valor de j para actualizar el peso
			for k := 0; k < numFeatures; k++ {
				//Calcular gradiente del peso (Sumatoria del error * x)
				errorAcumuladoPesos[k] += errorPrediccion * featuresX[j][k]
			}
		}
		//Actualizer parametros
		n := float64(muestras)
		//Actualizar sesgo sesgo=sesgo-tasaAprendizaje*(Σerror/n)
		aprendizaje.Sesgo -= aprendizaje.TasaAprendizaje * (errorAcumuladoSesgo / n)
		//Actualizar cada peso peso=peso-tasaAprendizaje(Σ(error*x)/n)
		for k := 0; k < numFeatures; k++ {
			aprendizaje.Pesos[k] -= aprendizaje.TasaAprendizaje * (errorAcumuladoPesos[k] / n)
		}
		//Calcular error cuadratico medio
		errorCuadraticoMedio := errorCuadratico/n
		if i%10 == 0 {
			fmt.Printf("Iteración %d - Sesgo: %.5f, Pesos: %.5f\n, Error: %.5f\n", i, aprendizaje.Sesgo, aprendizaje.Pesos, errorCuadraticoMedio)
		}
	}
	fmt.Printf("%v\n", "-----------")
}
