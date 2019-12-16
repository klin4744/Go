package main

func main() {}

func twoSum(xi []int, target int) bool {
	m := make(map[int]bool)
	for _, value := range xi {
		matchingPair := target - value
		_, ok := m[matchingPair]
		if ok {
			return true
		}
		m[value] = true

	}
	return false
}
