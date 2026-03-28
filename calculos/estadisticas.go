package calculos

import (
	"math"
	"sort"
)

func CalcularMedia(datos []float64) float64 {
	var media float64
	var sumatoria float64
	n := float64(len(datos))
	if n == 0{
		return 0
	}
	for i := 0; i < len(datos); i++ {
		sumatoria += datos[i]
	}
	media = sumatoria/n
	return media
}

// Argumentos: array de datos y la media de los datos calculada con CalcularMedia
func CalcularDesviacionEstandar(datos []float64, mediaDatos float64) float64 {
	var sumatoriaCuadrados float64
	n := float64(len(datos))
	if n <= 1 {
        return 0
    }
	for i := 0; i < len(datos); i++ {
        diferencia := datos[i] - mediaDatos
        sumatoriaCuadrados += diferencia * diferencia
    }
	varianza := sumatoriaCuadrados/(n-1)
	return math.Sqrt(varianza)
}

func CalcularMediana(datos []float64) float64{
	n := len(datos)
	if n == 0{
		return 0
	}
	copia := make([]float64, n)
	copy(copia, datos)
	sort.Float64s(copia)
	//Par
	if n%2==0{
		indiceCentro1 := (n / 2) - 1
        indiceCentro2 := n / 2
        return (copia[indiceCentro1] + copia[indiceCentro2]) / 2.0
	}
	return copia[n/2]
}

//Calular la asimetria utilizando el coeficiente de pearson 3(Media-Moda)/DesviacionEstandar
func CalcularAsimetria(datos []float64, media float64, desviacion float64) float64{
	if desviacion == 0 {
        return 0
    }
	mediana := CalcularMediana(datos)
	return 3 * (media - mediana) / desviacion
}

func MinMaxScaler(x, x_min, x_max float64) float64{
	if x_max == x_min {
        return 0
    }
	var numerador float64
	var denominador float64
	numerador = x - x_min
	denominador = x_max-x_min
	return numerador/denominador
}

func CalcularError(prediccion, datoReal float64) float64{
	return prediccion-datoReal
}

func CalcularGradientePeso(prediccion, datoReal, x float64) float64{
	return (prediccion-datoReal)*x
}

func CalcularGradienteSesgo(prediccion, datoReal float64) float64{
	return (prediccion-datoReal)
}