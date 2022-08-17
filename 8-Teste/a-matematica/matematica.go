package amatematica

// Media é responsável por calcular o que você ja sabe ..... :)
func media(numeros ...float64) float64 {
	total := 0.0
	for _, num := range numeros {
		total += num
	}

	return total / float64(len(numeros))
}
