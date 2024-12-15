---
title: "<font color='#01fb6a'>Práctica 3</font>"
---
### Sistemas Distribuidos
#### Autor: Carlos Nebril (carlosnj)

<br>
<hr>

## Índice

### [1. Descripción del Problema](#descripción-del-problema-1)

### [2. Arquitectura del Sistema](#arquitectura-del-sistema-1)
##### • [2.1 Estructura del Código](#estructura-del-código-1)
##### • [2.2 Componentes Estructurales](#componentes-estructurales-1)
##### • [2.3 Mecanismos de Concurrencia](#mecanismos-de-concurrencia-1)

### [3. Análisis de la Ejecución](#análisis-de-la-ejecución-1)
##### • [3.1 Diagrama UML de Secuencia](#diagrama-uml-de-secuencia-1)
##### • [3.2 Participantes del Diagrama](#participantes-del-diagrama-1)
##### • [3.3 Flujo de Interacciones](#flujo-de-interacciones-1)

### [4. Análisis de Rendimiento](#análisis-de-rendimiento-1)
##### • [4.1 Configuraciones de Prueba](#configuraciones-de-prueba-1)
##### • [4.2 Resultados Estadísticos](#resultados-estadísticos-1)
##### • [4.3 Gráfico de Rendimiento](#gráfico-de-rendimiento-1)

### [5. Principales Hallazgos](#principales-hallazgos-1)

### [6. Conclusiones](#conclusiones-1)


<br>
<hr>

## 1. Descripción del <font color='#01fb6a'>Problema</font>

Esta práctica aborda la simulación de un aeropuerto gestionando la llegada de aviones con concurrencia en el lenguaje Go. El objetivo principal es evaluar el rendimiento de la priorización por categorías y el uso eficiente de recursos limitados como pistas y puertas de desembarque.

En la simulación, una cantidad `N` de aviones llega a un aeropuerto con recursos limitados:  
- **<font color='#01fb6a'>Pistas de aterrizaje</font>**.  
- **<font color='#01fb6a'>Puertas de desembarque</font>**.  

Los aviones se clasifican en tres categorías:  
- **<font color='#01fb6a'>Categoría A</font>:** Más de 100 pasajeros (Alta prioridad).  
- **<font color='#01fb6a'>Categoría B</font>:** Entre 50 y 100 pasajeros (Prioridad normal).  
- **<font color='#01fb6a'>Categoría C</font>:** Menos de 50 pasajeros (Baja prioridad).  

El orden de atención depende de la prioridad y la disponibilidad de recursos. El objetivo es medir el tiempo total de simulación bajo diferentes configuraciones.  

## 2. <font color='#01fb6a'>Arquitectura</font> del Sistema

### 2.1 Estructura del <font color='#01fb6a'>Código</font>

- <font color='#01fb6a'>'main.go'</font>: Orquesta las simulaciones y mide los tiempos de ejecución.
- <font color='#01fb6a'>'simulacion.go'</font>: Contiene la lógica del aeropuerto y las operaciones de aterrizaje y desembarque.
- <font color='#01fb6a'>'simulacion_test.go'</font>: Define pruebas para validar el comportamiento del sistema.


### 2.2 Componentes <font color='#01fb6a'>Estructurales</font>

La simulación se compone de varios componentes fundamentales que modelan un sistema de control de tráfico aéreo:

- **<font color='#01fb6a'>Avión</font>**: 
  - Estructura que representa cada aeronave
  - Propiedades:
    - ID único
    - Categoría (A, B, C)
    - Número de pasajeros
  - Genera aleatoriamente sus características en tiempo de ejecución

- **<font color='#01fb6a'>Aeropuerto</font>**: 
  - Gestor central de recursos
  - Componentes clave:
    - Canal de pistas limitado
    - Cola de puertas de espera
    - Mecanismo de sincronización mutex

### 2.3 Mecanismos de <font color='#01fb6a'>Concurrencia</font>

- **<font color='#01fb6a'>Canales Go</font>**: 
  - `pistas`: Control de ocupación de pistas
  - `puertasEspera`: Gestión de desembarque
- **<font color='#01fb6a'>WaitGroup</font>**: Sincronización de goroutines
- **<font color='#01fb6a'>Mutex</font>**: Protección de secciones críticas
- **<font color='#01fb6a'>Concurrencia</font>**: Múltiples aviones procesados simultáneamente

## 3.1 <font color='#01fb6a'>Análisis</font> de la <font color='#01fb6a'>Ejecución</font> 

### 3.1 Diagrama UML de Secuencia

<div style="text-align: center;">
![Gentrificacion en la República Bananera](../images/diagrama_secuencia.svg)
</div>

### • **Participantes del Diagrama**

#### 1. <font color='#01fb6a'>`Main`</font>
- Punto de entrada de la aplicación
- Responsable de iniciar la simulación
- Invoca el método de simulación principal

#### 2. <font color='#01fb6a'>`SimularAeropuerto()`</font>
- Función central que orquesta toda la simulación
- Gestiona el flujo general de la ejecución
- Crea y coordina los recursos del aeropuerto

#### 3. <font color='#01fb6a'>`Torre de Control`</font>
- Gestor principal de los recursos
- Coordina la asignación de pistas
- Controla el flujo de aviones

#### 4. <font color='#01fb6a'>`Aviones`</font>
- Entidades dinámicas que solicitan servicios
- Generan solicitudes de aterrizaje y desembarque
- Tienen características únicas (ID, categoría, pasajeros)

#### 5. <font color='#01fb6a'>`Pistas`</font>
- Recurso limitado para aterrizaje
- Gestionan el tiempo y la disponibilidad de aterrizaje

#### 6. <font color='#01fb6a'>`Puertas de Desembarque`</font>
- Recursos para procesamiento posterior al aterrizaje
- Manejan el desembarque de pasajeros

### • **Flujo de Interacciones**

#### 1. Inicio de Simulación
1. <font color='#01fb6a'>`Main`</font> inicia <font color='#01fb6a'>`SimularAeropuerto()`</font>
2. <font color='#01fb6a'>`SimularAeropuerto()`</font> crea la <font color='#01fb6a'>`Torre de Control`</font>

#### 2. Ciclo de Procesamiento de Aviones

##### Generación de Avión
1. <font color='#01fb6a'>`SimularAeropuerto()`</font> genera un nuevo <font color='#01fb6a'>`Avión`</font>
2. <font color='#01fb6a'>`Avión`</font> recibe detalles aleatorios (pasajeros, categoría)

##### Solicitud de Aterrizaje
1. <font color='#01fb6a'>`Avión`</font> solicita aterrizaje a la <font color='#01fb6a'>`Torre de Control`</font>
2. <font color='#01fb6a'>`Torre de Control`</font> activa el recurso de <font color='#01fb6a'>`Pista`</font>
3. Tiempo de aterrizaje es aleatorio (simulando condiciones reales)
4. <font color='#01fb6a'>`Pista`</font> confirma aterrizaje al <font color='#01fb6a'>`Avión`</font>

##### Proceso de Desembarque
1. <font color='#01fb6a'>`Avión`</font> solicita desembarque a <font color='#01fb6a'>`Puertas de Desembarque`</font>
2. <font color='#01fb6a'>`Puerta`</font> procesa el desembarque
3. Tiempo de desembarque es variable
4. <font color='#01fb6a'>`Puerta`</font> confirma desembarque completado

#### 3. Ciclo de Concurrencia
1. El diagrama muestra un bucle que se repite para cada <font color='#01fb6a'>`avión`</font>
2. Cada <font color='#01fb6a'>`avión`</font> se procesa de manera concurrente
3. Uso de <font color='#01fb6a'>`goroutines`</font> y <font color='#01fb6a'>`canales`</font> para gestionar concurrencia

## 4. Análisis de Rendimiento

### 4.1 Configuraciones de Prueba

<table border="1" cellpadding="10" cellspacing="0" style="border-collapse: collapse; text-align: center; width: 100%;">
  <thead>
    <tr style="background-color: #01fb6a;">
      <th>Configuración</th>
      <th>Aviones A</th>
      <th>Aviones B</th>
      <th>Aviones C</th>
      <th>Num. Pistas</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td>Caso Base</td>
      <td>10</td>
      <td>10</td>
      <td>10</td>
      <td>3</td>
    </tr>
    <tr>
      <td>Carga Alta</td>
      <td>20</td>
      <td>5</td>
      <td>5</td>
      <td>3</td>
    </tr>
    <tr>
      <td>Carga Baja</td>
      <td>5</td>
      <td>5</td>
      <td>20</td>
      <td>3</td>
    </tr>
  </tbody>
</table>

### 4.2 Resultados Estadísticos

#### Tiempos de Simulación (10 Ejecuciones)

<table border="1" cellpadding="10" cellspacing="0" style="border-collapse: collapse; text-align: center; width: 100%;">
  <thead>
    <tr style="background-color: #01fb6a;">
      <th>Configuración</th>
      <th>Mínimo</th>
      <th>Máximo</th>
      <th>Promedio</th>
      <th>Desviación</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td>Caso Base</td>
      <td>3.89</td>
      <td>4.86</td>
      <td>4.44</td>
      <td>0.29</td>
    </tr>
    <tr>
      <td>Carga Alta (A prioritarios)</td>
      <td>4.02</td>
      <td>4.74</td>
      <td>4.43</td>
      <td>0.22</td>
    </tr>
    <tr>
      <td>Carga Baja (C dominante)</td>
      <td>3.89</td>
      <td>4.78</td>
      <td>4.47</td>
      <td>0.30</td>
    </tr>
  </tbody>
</table>

### 4.3 Gráfico de <font color='#01fb6a'>Rendimiento</font>

<div style="text-align: center;">
![Gráfico de Rendimiento](../images/grafico_rendimiento.png)
</div>

## 5. Principales <font color='#01fb6a'>Hallazgos</font>

1. **<font color='#01fb6a'>Consistencia del Sistema</font>:**
   - Los resultados muestran que el sistema mantiene **<font color='#01fb6a'>tiempos promedio</font>** muy similares (~4.45 segundos) bajo diferentes configuraciones.
   - La **<font color='#01fb6a'>variabilidad</font>** entre ejecuciones es baja, con una desviación estándar promedio de 0.27 segundos, lo que indica estabilidad en el rendimiento.

2. **<font color='#01fb6a'>Eficiencia de Priorización</font>:**
   - La priorización de **<font color='#01fb6a'>Categorías A</font>** garantiza que los aviones con mayor prioridad sean atendidos rápidamente, reduciendo el impacto de cargas mixtas.
   - Las configuraciones con mayor proporción de aviones de Categoría A muestran menor varianza en los tiempos totales.

3. **<font color='#01fb6a'>Tolerancia a Cargas Variables</font>:**
   - El sistema demuestra adaptabilidad frente a distribuciones de tráfico diferentes, manteniendo tiempos promedio similares en los escenarios de **<font color='#01fb6a'>Carga Alta</font>** y **<font color='#01fb6a'>Carga Baja</font>**.
   - La implementación de mecanismos concurrentes, como goroutines y canales, asegura que los recursos sean utilizados de manera eficiente incluso bajo cargas desbalanceadas.

## 6. Conclusiones

1. **<font color='#01fb6a'>Escalabilidad y Concurrencia</font>:**
   - El diseño concurrente con goroutines, canales, y mutex permite gestionar múltiples aviones simultáneamente, logrando un **<font color='#01fb6a'>uso óptimo</font>** de recursos limitados.
   - La arquitectura es fácilmente escalable para aumentar el número de aviones o recursos, manteniendo la estabilidad del sistema.

2. **<font color='#01fb6a'>Rendimiento Consistente</font>:**
   - Los resultados demuestran que el sistema opera de manera consistente, con tiempos de simulación promedio estables y baja variabilidad, lo que lo hace adecuado para situaciones con **<font color='#01fb6a'>requerimientos estrictos</font>** de tiempo.

3. **<font color='#01fb6a'>Prioridades Efectivas</font>:**
   - La implementación de un esquema de prioridades basado en categorías mejora la gestión del tráfico aéreo, especialmente bajo **<font color='#01fb6a'>cargas mixtas</font>**. Esto permite optimizar el tiempo total de simulación y asegura que los aviones de mayor importancia sean atendidos rápidamente.

4. **<font color='#01fb6a'>Flexibilidad</font>:**
   - El sistema se adapta bien a diferentes escenarios de carga, mostrando su capacidad para manejar situaciones diversas sin comprometer el rendimiento. Esto lo hace ideal para aplicaciones que requieren **<font color='#01fb6a'>ajustes dinámicos</font>** en tiempo de ejecución.

En general, la simulación cumple su propósito de evaluar el **<font color='#01fb6a'>rendimiento</font>** en escenarios complejos, validando la eficacia de las técnicas de concurrencia y priorización implementadas en Go.

El código fuente completo de este proyecto está disponible en el siguiente repositorio: [https://github.com/cnebril2020/practica3_sistemasdistribuidos](https://github.com/cnebril2020/practica3_sistemasdistribuidos)



