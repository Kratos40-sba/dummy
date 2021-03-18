package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func random(n int) []float64 {
	l := make([]float64, n)
	for i := 0; i < n; i++ {
		l[i] = float64(i * rand.Intn(10))

	}
	return l
}
func avg(l []float64) float64 {
	var s float64
	for _, v := range l {
		s = v + s
	}
	return s / float64(len(l))
}
func variance(l []float64, avg float64) float64 {
	var s float64
	for i := 0; i < len(l); i++ {
		s = math.Sqrt(math.Pow(s+l[i]-avg, 2))
	}
	return s / float64(len(l))
}
func cartType(variance float64) float64 {
	return math.Pow(variance, 2)
}
func main() {
	rand.Seed(time.Now().UnixNano())
	l := random(20)
	avg := avg(l)
	v := variance(l, avg)
	e := cartType(v)
	fmt.Println("Numbers :", l)
	fmt.Println("Avg :", avg)
	fmt.Println("Variance / Ecart-type :", v, e)

}
