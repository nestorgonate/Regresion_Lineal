package models

//Struct que almacena los datos crudos del CSV
type CsvData struct {
	Temperature       []float64
	Relative_humidity []float64
	Precipitation     []float64
	Pressure          []float64
}

//Struct que almacena los datos normalizados como relative_humidity o precipitacion utilizando z score
type NormalizeData struct {
	Temperature        []float64
	Relative_humidity []float64
	Precipitation     []float64
	Pressure          []float64
}