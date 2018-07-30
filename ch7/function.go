package main

func Average(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}

func F() (int, int) {
	return 5, 6
}

func F1() int {
	return F2()
}

func F2() (r int) {
	r = 1
	return
}
