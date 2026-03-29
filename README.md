# Prediccion de temperatura con Regresion Lineal 🌡️
## Implementacion en Golang
Este proyecto consiste en el uso de regresion lineal para predecir las temperaturas en base a humedad, precipitacion y presion.
Este sistema intenta ser 0 dependencias, no utiliza librerías de ML externas, implementando cada cálculo de forma nativa.
## Decisiones en el desarrollo
Se optó por implementar un Min-Max Scaler en lugar de Z-Score para garantizar que todas las características se encontraran en el rango $[0, 1]$ al observar que los datos no siguen una distribucion normal,
de esta manera se pudo normalizar valores atipicos como la presion o humedad
## Estructura del proyecto
### 1. data/readcsv.go
Este paquete contiene el lector para obtener los datos del CSV y organizarlos en un HashMap
### 2. calculos/estadisticas.go
Este paquete contiene los calculos de estadistica para la descripción y preparación de los datos.

Media ($\bar{x}$): $\bar{x} = \frac{1}{n} \sum_{i=1}^{n} x_i$

Mediana ($\tilde{x}$): Valor central de los datos ordenados.

Desviación Estándar ($\sigma$): $\sigma = \sqrt{\frac{\sum (x_i - \bar{x})^2}{n}}$

Asimetría (Pearson): $A_p = \frac{3(\bar{x} - \tilde{x})}{\sigma}$

Min-Max Scaler: $x_{norm} = \frac{x - x_{min}}{x_{max} - x_{min}}$

Gradientes:

- Peso ($w$): $\frac{\partial J}{\partial w} = (\hat{y} - y) \cdot x$

- Sesgo ($b$): $\frac{\partial J}{\partial b} = (\hat{y} - y)$
### 3. calculos/regresion.go
Implementación de la prediccion y optimizacion

Regresión lineal multiple: $\hat{y} = (\sum_{i=1}^{k} w_i \cdot x_i) + b$

Entrenamiento: El modelo ajusta los pesos mediante iteraciones, actualizando el peso y sesgo en cada iteracion (1000).

- $w = w - \alpha \cdot \frac{1}{n} \sum Gradiente_w$

- $b = b - \alpha \cdot \frac{1}{n} \sum Gradiente_b$

Desnormalización: Para retornar el valor a la escala Celsius

$y_{celsius} = (\hat{y}_{norm} \cdot (max - min)) + min$

## 4. temperature_data.csv
Dataset con 1386 registros del clima
