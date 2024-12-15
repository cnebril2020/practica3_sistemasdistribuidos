package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Establecer semilla para aleatorios
	rand.Seed(time.Now().UnixNano())

	// Almacenar tiempos de simulaciones
	tiempos := make([]float64, 3)

	// Ejecutar simulaciones de prueba
	fmt.Println("--- Simulación 1: 10 aviones A, 10 aviones B, 10 aviones C ---")
	tiempos[0] = SimularAeropuerto(30, 3)

	fmt.Println("\n--- Simulación 2: 20 aviones A, 5 aviones B, 5 aviones C ---")
	tiempos[1] = SimularAeropuerto(30, 3)

	fmt.Println("\n--- Simulación 3: 5 aviones A, 5 aviones B, 20 aviones C ---")
	tiempos[2] = SimularAeropuerto(30, 3)

	// Comparar tiempos
	fmt.Println("\nComparación de Tiempos:")
	for i, tiempo := range tiempos {
		fmt.Printf("Simulación %d: %.2f segundos\n", i+1, tiempo)
	}
}