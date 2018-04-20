package util

// MeanUint8 calculate the mean value of an array of int
func MeanUint8(arr []uint8) float64 {
	var sum int
	for _, v := range arr {
		sum += int(v)
	}
	return float64(sum / len(arr))
}