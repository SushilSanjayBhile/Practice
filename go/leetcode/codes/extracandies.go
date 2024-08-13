package codes

func KidsWithCandies(candies []int, extraCandies int) []bool {
	result := make([]bool, len(candies))

	max := candies[0]
	for _, v := range candies {
		if max < v {
			max = v
		}
	}

	for k, v := range candies {
		if v+extraCandies >= max {
			result[k] = true
		} else {
			result[k] = false
		}
	}
	return result
}
