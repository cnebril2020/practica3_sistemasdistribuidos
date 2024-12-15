---
title: "<font color='#01fb6a'>Práctica 2</font>"
---
### Sistemas Distribuidos
#### Autor: Carlos Nebril (carlosnj)

<br>
<hr>

## Índice

### [1. Arquitectura del Sistema](#arquitectura-del-sistema-1)
##### • [1.1 Componentes Estructurales](#componentes-estructurales-1)
##### • [1.2 Mecanismos de Sincronización](#mecanismos-de-sincronización-1)

### [2. Análisis de la Ejecución](#análisis-de-la-ejecución-1)

##### • [2.1 Diagrama UML de Flujo](#diagrama-uml-de-flujo-1)
##### • [2.2 Ejemplo de Análisis del Flujo](#ejemplo-de-análisis-del-flujo-1)
##### • [2.3 Métricas Iniciales](#métricas-iniciales-1)

### [3. Análisis de Tests de Rendimiento](#análisis-de-tests-de-rendimiento-1)

##### • [3.1 Configuraciones Evaluadas](#configuraciones-evaluadas-1)
##### • [3.2 Análisis de Resultados](#análisis-de-resultados-1)
##### • [3.3 Hallazgos Clave](#hallazgos-clave-1)

### [4. Conclusiones](#conclusiones-1)


<br>
<hr>

## 1. <font color='#01fb6a'>Arquitectura</font> del Sistema

El sistema implementa un aeropuerto utilizando Go y sus características de concurrencia. La arquitectura se basa en los siguientes componentes principales:

### 1.1 Componentes Estructurales

- **<font color='#01fb6a'>Torre de Control</font>**: Gestiona el tráfico aéreo mediante un canal buffered y mutex para control de concurrencia
- **<font color='#01fb6a'>Pistas</font>**: Recursos compartidos con control de ocupación mediante mutex
- **<font color='#01fb6a'>Puertas de Embarque</font>**: Recursos compartidos con control de ocupación mediante mutex
- **<font color='#01fb6a'>Aviones</font>**: Implementados como goroutines independientes

### 1.2 Mecanismos de Sincronización

- **<font color='#01fb6a'>Canales</font>**: Utilizados para la gestión de la cola de la torre de control
- **<font color='#01fb6a'>Mutex</font>**: Empleados para el control de acceso a recursos compartidos
- **<font color='#01fb6a'>WaitGroup</font>**: Coordina la finalización de todas las goroutines de aviones

## 2. Análisis de la Ejecución

### 2.1 Diagrama <font color='#01fb6a'>UML</font> de <font color='#01fb6a'>Flujo</font>

<div style="text-align: center;">
![Diagrama de Flujo UML](images/diagrama_UML.svg)
</div>

### 2.2 Ejemplo de Análisis del Flujo

Basado en la salida del programa con la <font color='#01fb6a'>configuración base</font> (básica):
```go
NumAviones: 10
NumPistas: 2
NumPuertas: 3
MaxColaAviones: 5
...
```

1. **Fase Inicial (Solicitud de Permiso)**:
   - Los 10 aviones intentan solicitar permiso <font color='#01fb6a'>simultáneamente</font>
   - Solo 5 aviones logran <font color='#01fb6a'>entrar</font> en la cola
   - 5 aviones son <font color='#01fb6a'>rechazados</font> por cola llena

2. **Fase de Aterrizaje**:
   - Las dos pistas se utilizan eficientemente <font color='#01fb6a'>en paralelo</font>
   - Secuencia de aterrizajes (ejemplo):
     - Pista 0: 7 → 8 → 5
     - Pista 1: 6 → 9

3. **Fase de Desembarque**:
   - Las tres puertas se utilizan de manera <font color='#01fb6a'>rotativa</font>
   - Secuencia de desembarques (ejemplo):
     - Puerta 0: 7 → 9
     - Puerta 1: 6 → 5
     - Puerta 2: 8

### 2.3 Métricas <font color='#01fb6a'>Iniciales</font>

1. <font color='#01fb6a'>**Tasa de Aceptación**</font>:
   - 50% de los aviones procesados exitosamente (5 de 10)
   - 50% rechazados por limitación de cola

2. <font color='#01fb6a'>**Utilización de Recursos**</font>:
   - **Pistas**: Distribución equilibrada entre las dos pistas
   - **Puertas**: Todas las puertas fueron utilizadas al menos una vez

3. <font color='#01fb6a'>**Patrones de Tiempo**</font>:
   - Los aterrizajes se solapan eficientemente entre las dos pistas
   - El desembarque muestra una distribución uniforme entre las puertas disponibles

## 3. Análisis de <font color='#01fb6a'>Tests de Rendimiento</font>

Esta sección presenta los resultados de las pruebas de rendimiento realizadas sobre diferentes configuraciones del sistema de simulación de aeropuerto. Se realizaron <font color='#01fb6a'>5 ejecuciones completas</font> para obtener datos estadísticamente significativos.

### 3.1 Configuraciones Evaluadas

1. <font color='#01fb6a'>**Configuración Base**</font> (detallada)
   - 10 aviones
   - 2 pistas
   - 3 puertas
   - Cola máxima de 5 aviones
   - Tiempo base de control: 2s
   - Tiempo base de pista: 3s
   - Tiempo de puerta: 4s
   - Variación de tiempo: 25%

2. <font color='#01fb6a'>**Cola Duplicada**</font>
   - Igual que la configuración base pero con cola máxima de 10 aviones

3. <font color='#01fb6a'>**Variación Tiempo 25%**</font>
   - Igual que la configuración base pero con variación de tiempo ajustada al 25%

4. <font color='#01fb6a'>**Cola Doble + Variación 25%**</font>
   - Combinación de cola duplicada y variación de tiempo 25%

5. <font color='#01fb6a'>**Pistas x5**</font>
   - Igual que la configuración base pero con 10 pistas

6. <font color='#01fb6a'>**Pistas x5 + Tiempo x5**</font>
   - 10 pistas
   - Tiempo base de pista multiplicado por 5 (15s)

## 3.2 Resultados Detallados

### <font color='#01fb6a'>Tiempos</font> de ejecución (en segundos)

| Configuración               | Test 1  | Test 2  | Test 3  | Test 4  | Test 5  | Promedio | Desv. Est. |
|----------------------------|---------|---------|---------|---------|---------|----------|------------|
| Configuración Base         | 14.36   | 13.42   | 14.44   | 16.13   | 14.76   | 14.62    | ±0.98      |
| Cola Duplicada             | 23.38   | 22.13   | 20.61   | 21.56   | 23.18   | 22.17    | ±1.12      |
| Variación Tiempo 25%       | 12.88   | 12.61   | 14.37   | 14.30   | 15.22   | 13.88    | ±1.05      |
| Cola Doble + Variación 25% | 20.80   | 20.83   | 21.39   | 22.80   | 22.79   | 21.72    | ±0.91      |
| Pistas x5                  | 12.89   | 13.48   | 12.81   | 13.05   | 13.30   | 13.11    | ±0.28      |
| Pistas x5 + Tiempo x5      | 24.60   | 27.57   | 24.62   | 24.53   | 24.14   | 25.09    | ±1.39      |

<div style="text-align: center;">
![Gráfica Tiempos de Ejecución](images/metricas.png)
</div>

### Tiempo Total de Ejecución

- <font color='#01fb6a'>Promedio</font>: 110.59 segundos
- <font color='#01fb6a'>Desviación Estándar</font>: ±2.13 segundos

## 3.2 Análisis de <font color='#01fb6a'>Resultados</font>

1. **<font color='#01fb6a'>Configuración Base</font> (14.62s ±0.98s)**
   - Establece el punto de referencia para comparar otras configuraciones
   - Muestra un rendimiento consistente con variabilidad moderada

2. **<font color='#01fb6a'>Cola Duplicada</font> (22.17s ±1.12s)**
   - Incremento del 51.6% en tiempo de ejecución
   - Peor rendimiento que la configuración base
   - Demuestra que aumentar la cola no mejora el rendimiento

3. **<font color='#01fb6a'>Variación Tiempo 25%</font> (13.88s ±1.05s)**
   - Mejora del 5.1% respecto a la base
   - Sugiere que la variabilidad en los tiempos puede reducir cuellos de botella

4. **<font color='#01fb6a'>Cola Doble + Variación 25%</font> (21.72s ±0.91s)**
   - Rendimiento similar a la cola duplicada
   - Confirma que el tamaño de cola aumentado no beneficia al sistema

5. **<font color='#01fb6a'>Pistas x5</font> (13.11s ±0.28s)**
   - Mejor rendimiento de todas las configuraciones
   - Mejora del 10.3% respecto a la base
   - Menor variabilidad (más consistente)

6. **<font color='#01fb6a'>Pistas x5 + Tiempo x5</font> (25.09s ±1.39s)**
   - Peor rendimiento general
   - Mayor variabilidad entre pruebas
   - El aumento en tiempo de pista neutraliza los beneficios de tener más pistas

## 3.3 Hallazgos <font color='#01fb6a'>Clave</font>

1. <font color='#01fb6a'>**Estabilidad del Sistema**</font>
   - La configuración más estable: Pistas x5 (±0.28s)
   - La configuración menos estable: Pistas x5 + Tiempo x5 (±1.39s)

2. <font color='#01fb6a'>**Cuellos de Botella Identificados**</font>
   - El tamaño de la cola no es un factor limitante
   - El número de pistas afecta significativamente al rendimiento
   - Los tiempos de operación son críticos para el rendimiento general

## 4. Conclusiones

El trabajo analiza un sistema de simulación de un aeropuerto en Go, enfocándose en la gestión de la <font color='#01fb6a'>concurrencia</font> y la <font color='#01fb6a'>sincronización</font>. Se destacan los siguientes hallazgos principales:

1. **Optimización de Recursos**: Aumentar el número de pistas fue la mejora más efectiva, <font color='#01fb6a'>reduciendo</font> el tiempo de ejecución y <font color='#01fb6a'>estabilizando</font> el rendimiento. Por otro lado, <font color='#01fb6a'>incrementar</font> el tamaño de la cola no mejoró la eficiencia, subrayando la importancia de una asignación óptima de recursos.

2. **Efectos de la Variabilidad**: Introducir variaciones de tiempo mejoró ligeramente la ejecución, sugiriendo que una cierta <font color='#01fb6a'>flexibilidad temporal</font> puede aliviar cuellos de botella, aunque la <font color='#01fb6a'>configuración base</font> sigue siendo un buen punto de comparación.

3. **Complejidad Temporal**: Aumentar los tiempos de pista mostró que, aunque es posible <font color='#01fb6a'>ampliar recursos físicos</font>, los costes temporales pueden <font color='#01fb6a'>anular</font> dichos beneficios, evidenciando la importancia de un balance entre disponibilidad y eficiencia.

En resumen, este análisis enfatiza que mejorar la <font color='#01fb6a'>eficiencia</font> de los recursos es más valioso que simplemente incrementar su número. Los resultados proporcionan una base sólida para <font color='#01fb6a'>futuras optimizaciones</font> en sistemas distribuidos.

El código fuente completo de este proyecto está disponible en el siguiente repositorio: [https://github.com/cnebril2020/sistemasdistribuidos](https://github.com/cnebril2020/sistemasdistribuidos)
