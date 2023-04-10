package code

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	var min, max = prices[0], prices[0]
	var res int

	for i := 1; i < len(prices); i++ {
		if min > prices[i] {
			if res < max-min {
				res = max - min
			}

			min = prices[i]
			max = prices[i]
		} else {
			if max < prices[i] {
				max = prices[i]
			}
		}
	}

	if res < max-min {
		res = max - min
	}
	return res
}
