package main

import (
	"testing"
)

func TestSimulacionAeropuerto(t *testing.T) {
	testCases := []struct {
		nombre       string
		numAviones   int
		numPistas    int
	}{
		{"Simulación 30 aviones", 30, 3},
		{"Simulación 50 aviones", 50, 5},
		{"Simulación 20 aviones", 20, 2},
	}

	for _, tc := range testCases {
		t.Run(tc.nombre, func(t *testing.T) {
			SimularAeropuerto(tc.numAviones, tc.numPistas)
		})
	}
}

func BenchmarkSimulacionAeropuerto(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SimularAeropuerto(50, 3)
	}
}