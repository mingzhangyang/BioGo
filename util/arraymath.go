package util

// MeanUint8 calculate the mean value of an array of uint8
func MeanUint8(arr []uint8) float64 {
	var sum int
	for _, v := range arr {
		sum += int(v)
	}
	// return float64(sum / len(arr))
	// This is probelmatic. A integer divided by another integer alwasys get a integer.
	// Type conversion to float64 should be done before division.
	return float64(sum) / float64(len(arr))
}

// SumInt32 calculate the sum of an array of int32
func SumInt32(arr []int32) int64 {
	var sum int64
	for _, v := range arr {
		sum += int64(v)
	}
	return sum
}

// SumUint8 calculate the sum of an array of int32
func SumUint8(arr []uint8) int64 {
	var sum int64
	for _, v := range arr {
		sum += int64(v)
	}
	return sum
}

// SumInt64 calculate the sum of an array of int32
func SumInt64(arr []int64) int64 {
	var sum int64
	for _, v := range arr {
		sum += int64(v)
	}
	return sum
}