package main

import (
	"math"
	"testing"
)

func TestEstimateCircleArea(t *testing.T) {
	ray := 1.0
	samples := 1000000
	expectedArea := math.Pi * ray * ray

	area := EstimateCircleArea(ray, samples)

	if math.Abs(area-expectedArea) > 0.01 {
		t.Fatalf("Expected area: %f, got: %f", expectedArea, area)
	}
}

func TestEstimateCircleAreaParallel(t *testing.T) {
	ray := 1.0
	samples := 1000000
	expectedArea := math.Pi * ray * ray

	area := EstimateCircleAreaParallel(ray, samples)

	if math.Abs(area-expectedArea) > 0.01 {
		t.Fatalf("Expected area: %f, got: %f", expectedArea, area)
	}
}

func TestEstimateCircleAreaConcurrent(t *testing.T) {
	ray := 1.0
	samples := 1000000
	expectedArea := math.Pi * ray * ray

	area := EstimateCircleAreaConcurrent(ray, samples)

	if math.Abs(area-expectedArea) > 0.01 {
		t.Fatalf("Expected area: %f, got: %f", expectedArea, area)
	}
}

func BenchmarkEstimateCircleArea(b *testing.B) {
	ray := 1.0
	samples := 1000000

	for i := 0; i < b.N; i++ {
		EstimateCircleArea(ray, samples)
	}
}

func BenchmarkEstimateCircleAreaParallel(b *testing.B) {
	ray := 1.0
	samples := 1000000

	for i := 0; i < b.N; i++ {
		EstimateCircleAreaParallel(ray, samples)
	}
}

func BenchmarkEstimateCircleAreaConcurrent(b *testing.B) {
	ray := 1.0
	samples := 1000000

	for i := 0; i < b.N; i++ {
		EstimateCircleAreaConcurrent(ray, samples)
	}
}
