package anal

func Accuracy() {
	panic("not implemented")
}

/*
[0][0] TruePositive
[0][1] TrueNegative
[1][0] FalsePositive
[1][1] FalseNegative
*/
type ConfusionMatrix [2][2]int

func AccuracyConfusionMatrix(matrix ConfusionMatrix) float64 {
	return (float64(matrix[0][0]) + float64(matrix[0][1])) / (float64(matrix[0][0]) + float64(matrix[0][1]) + float64(matrix[1][0]) + float64(matrix[1][1]))
}
