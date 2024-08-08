package main

import (
	"math/rand/v2"
	"runtime"
	"sync"
)

func EstimateCircleArea(ray float64, samples int) float64 {
	pointsInside := 0

	for i := 0; i < samples; i++ {
		x := rand.Float64()
		y := rand.Float64()

		if x*x+y*y <= ray {
			pointsInside++
		}
	}

	return (float64(pointsInside) / float64(samples)) * 4
}

func EstimateCircleAreaParallel(ray float64, samples int) float64 {
	pointsInside := 0
	ch := make(chan int, samples)

	for i := 0; i < samples; i++ {
		go func() {
			x := rand.Float64()
			y := rand.Float64()

			if x*x+y*y <= ray {
				ch <- 1
			} else {
				ch <- 0
			}
		}()
	}

	for i := 0; i < samples; i++ {
		pointsInside += <-ch
	}

	return (float64(pointsInside) / float64(samples)) * 4
}

func estimateCircleAreaWorker(ray float64, samples int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < samples; i++ {
		x := rand.Float64()
		y := rand.Float64()

		if x*x+y*y <= ray {
			ch <- 1
		} else {
			ch <- 0
		}
	}
}

func EstimateCircleAreaConcurrent(ray float64, samples int) float64 {
	var wg sync.WaitGroup
	pointsInside := 0
	ch := make(chan int, samples)
	samplesPerRoutine := samples / runtime.NumCPU()

	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)
		go estimateCircleAreaWorker(ray, samplesPerRoutine, ch, &wg)
	}

	wg.Wait()
	close(ch)

	for points := range ch {
		pointsInside += points
	}

	return (float64(pointsInside) / float64(samples)) * 4
}
