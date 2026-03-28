package models

type Aprendizaje struct {
	Pesos           []float64 //Almacena un peso para cada feature
	Sesgo           float64
	TasaAprendizaje float64
}
