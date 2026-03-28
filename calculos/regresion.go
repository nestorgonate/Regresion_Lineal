package calculos

func RegresionLineal(datos, pesos []float64, sesgo float64) float64{
	var prediccion float64
	for i:=0;i<len(datos);i++{
		prediccion = pesos[i]*datos[i]
	}
	return prediccion+sesgo
}