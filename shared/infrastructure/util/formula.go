package util

func FormulaUpdateProductStock(newStock int, stockCurrent int, reserveStock int) int {
	var totalNewStock int
	var returnFinalStock int

	totalNewStock = newStock - (stockCurrent - reserveStock)
	returnFinalStock = totalNewStock + stockCurrent

	return returnFinalStock
}
