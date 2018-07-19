package fundamentals

// Constants ...
func Constants() (float64, float64, float64) {
	var boilingF float64 = 212.0
	var freezingF float64 = 32.0
	var absoluteZeroF float64 = -459.67

	var celsius []float64 = Convert(boilingF, freezingF, absoluteZeroF)

	return celsius[0], celsius[1], celsius[2]
}
