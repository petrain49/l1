package main

/*
Дана последовательность температурных колебаний:
-25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
Объединить данные значения в группы с шагом в 10 градусов.
Последовательность в подмножноствах не важна.

Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.
*/
func combineTemps(arr []float32) map[float32][]float32 {
	steps := [...]float32{0, -10, -20, -30, -40, -50, -60, 10, 20, 30, 40, 50, 60}

	sorted := make(map[float32][]float32)

	for _, t := range arr {
		for _, step := range steps {
			if abs(t-step) < 10 {
				sorted[step] = append(sorted[step], t)
				break
			}
		}
	}

	return sorted
}

func abs(n float32) float32 {
	if n < 0 {
		return -n
	}
	return n
}
