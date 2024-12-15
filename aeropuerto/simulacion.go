package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Definición de estructuras
type Avion struct {
	ID            int
	Categoria     string
	NumPasajeros  int
}

type Aeropuerto struct {
	pistas        chan struct{}
	puertasEspera chan Avion
	mu            sync.Mutex
}

// Función para crear un nuevo aeropuerto
func nuevoAeropuerto(numPistas int) *Aeropuerto {
	return &Aeropuerto{
		pistas:        make(chan struct{}, numPistas),
		puertasEspera: make(chan Avion, 100),
	}
}

// Generar avión con categoría aleatoria
func generarAvion(id int) Avion {
	numPasajeros := rand.Intn(200)
	var categoria string
	
	switch {
	case numPasajeros > 100:
		categoria = "A"
	case numPasajeros >= 50 && numPasajeros <= 100:
		categoria = "B"
	default:
		categoria = "C"
	}
	
	return Avion{
		ID:           id,
		Categoria:    categoria,
		NumPasajeros: numPasajeros,
	}
}

// Simular aterrizaje
func (a *Aeropuerto) aterrizar(avion Avion) {
	a.mu.Lock()
	fmt.Printf("Avión %d (Categoría %s, %d pasajeros) solicitando aterrizaje\n", 
		avion.ID, avion.Categoria, avion.NumPasajeros)
	a.mu.Unlock()
	
	// Simular tiempo de aterrizaje con variabilidad
	tiempo := time.Duration(rand.Intn(3000)) * time.Millisecond
	time.Sleep(tiempo)
	
	a.pistas <- struct{}{}
	defer func() { <-a.pistas }()
	
	a.mu.Lock()
	fmt.Printf("Avión %d aterrizó en pista después de %v\n", avion.ID, tiempo)
	a.mu.Unlock()
}

// Simular desembarque
func (a *Aeropuerto) desembarcar(avion Avion) {
	a.puertasEspera <- avion
	
	tiempo := time.Duration(rand.Intn(2000)) * time.Millisecond
	time.Sleep(tiempo)
	
	a.mu.Lock()
	fmt.Printf("Desembarco del avión %d con %d pasajeros completado\n", 
		avion.ID, avion.NumPasajeros)
	a.mu.Unlock()
	
	<-a.puertasEspera
}

// Simular aeropuerto con múltiples aviones
func SimularAeropuerto(numAviones int, numPistas int) float64 {
	inicio := time.Now()
	aeropuerto := nuevoAeropuerto(numPistas)
	
	var wg sync.WaitGroup
	
	for i := 1; i <= numAviones; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			
			avion := generarAvion(id)
			aeropuerto.aterrizar(avion)
			aeropuerto.desembarcar(avion)
		}(i)
	}
	
	wg.Wait()
	
	duracion := time.Since(inicio)
	fmt.Printf("\nTiempo total de simulación: %.2f segundos\n", duracion.Seconds())
	
	return duracion.Seconds()
}