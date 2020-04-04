package Sub

func Average(i []int) int {
	result := 0
	for _, value := range i {
		result += value
	}
	return int(result/len(i))
}
