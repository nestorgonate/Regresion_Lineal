package calculos

import (
	"fmt"
	"regresionlineal/models"
)

func RegresionLineal(datos, pesos []float64, sesgo float64) float64 {
	var prediccion float64
	for i := 0; i < len(datos); i++ {
		prediccion = pesos[i] * datos[i]
	}
	return prediccion + sesgo
}

func Entrenamiento(iteraciones, numFeatures, muestras int, targetY []float64,
					featuresX [][]float64) models.Aprendizaje{
	aprendizaje := models.Aprendizaje{
		Pesos:           []float64{1, 0.1, 0.1},
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
			prediccion := RegresionLineal(featuresX[j], aprendizaje.Pesos, aprendizaje.Sesgo)
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
		errorCuadraticoMedio := errorCuadratico / n
		if i%10 == 0 {
			fmt.Printf("Iteración %d - Sesgo: %v, Pesos: %v\n, Error: %v\n", i, aprendizaje.Sesgo, aprendizaje.Pesos, errorCuadraticoMedio)
		}
	}
	return aprendizaje
}

//Desnormalizar un valor usando x=(x_normalizado*(x_max-x_min))+x_min
func Desnormalizar(valorNormalizado, x_min, x_max float64) float64 {
    return (valorNormalizado * (x_max - x_min)) + x_min
}

func PredecirTemperatura(relative_humidity, precipitation, pressure float64,
						aprendizaje models.Aprendizaje, x_min_max map[string]map[string]float64) float64{
	normalizarHumedad := MinMaxScaler(relative_humidity, x_min_max["relative_humidity"]["x_min"], x_min_max["relative_humidity"]["x_max"])
	normalizarPrecipitacion := MinMaxScaler(precipitation, x_min_max["precipitation"]["x_min"], x_min_max["precipitation"]["x_max"])
	normalizarPresion := MinMaxScaler(pressure, x_min_max["pressure"]["x_min"], x_min_max["pressure"]["x_max"])
	//Prediccion y = w1*x1+w2*x2+w3*x3+b
	temperatura := (aprendizaje.Pesos[0]*normalizarHumedad)+
					(aprendizaje.Pesos[1]*normalizarPrecipitacion)+
					(aprendizaje.Pesos[2]*normalizarPresion)+
					aprendizaje.Sesgo
	temperaturaCelsius := Desnormalizar(temperatura, x_min_max["temperature"]["x_min"], x_min_max["temperature"]["x_max"])
	return temperaturaCelsius
}