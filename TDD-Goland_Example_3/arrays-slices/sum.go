package slices

func Sum(nums []int) int {
	var total int
	for _, v := range nums {
		total += v
	}
	return total
}

/*
My initial approach, clearly not the most optimal way
Issues with this approach, I limit the function to take in
two slices, should be able to handle more!

Next Im running two for loops to append to the same total slice
Not dynamic if Im summing more than two slices!
*/
func SumAll(s1, s2 []int) []int {
	var total []int
	var sum1 int
	var sum2 int
	for _, v := range s1 {
		sum1 += v
	}
	total = append(total, sum1)

	for _, v := range s2 {
		sum2 += v
	}
	total = append(total, sum2)

	return total
}

//Version 2 of the function, performs better O(n)
func SumAllV2(numbersToSum ...[]int) []int {
	var sums []int
	for _, nums := range numbersToSum {
		sums = append(sums, Sum(nums))
	}

	return sums
}

func SumAllTails(numsToSum ...[]int) []int {
	var total []int
	for _, nums := range numsToSum {

		if len(nums) == 0 {
			total = append(total, 0)
		} else {
			//Tail says first element in slice to the end
			tail := nums[1:]
			total = append(total, Sum(tail))
		}
	}
	return total

}
