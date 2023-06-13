package no

func NumInList(num int, nums []int) bool {
	for _, s := range nums {
		if s == num {
			return true
		}
	}
	return false
}
