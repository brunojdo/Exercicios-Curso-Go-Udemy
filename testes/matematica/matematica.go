package matematica

import (
	"fmt"
	"strconv"
)

// Media é responsável por calculucar a
// média dos valores passados
func Media(numeros ...float64) float64 {
	total := 0.0
	for _, num := range numeros {
		total += num
	}
	media := total / float64(len(numeros))
	mediaArredondada, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", media), 64)
	return mediaArredondada
}

/* func main() {
	media := Media(4, 8, 12, 4)
	fmt.Println(media)
}
*/
