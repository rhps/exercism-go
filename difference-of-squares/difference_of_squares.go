package diffsquares

func SquareOfSum(n int) int {
	tot := 0
	for i := 1; i < n+1; i++ {
		tot = tot + i
	}

	return tot * tot
}

func SumOfSquares(n int) int {
	tot := 0
	for i := 1; i < n+1; i++ {
		tot = tot + (i * i)
	}

	return tot
}

func Difference(n int) int {
	if SumOfSquares(n) >= SquareOfSum(n) {
		return SumOfSquares(n) - SquareOfSum(n)
	} else {
		return SquareOfSum(n) - SumOfSquares(n)
	}
}
