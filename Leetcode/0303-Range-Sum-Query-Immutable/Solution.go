package Solution

type NumArray struct {
	Sum []int
}

func Constructor(nums []int) NumArray {
	var res NumArray
	res.Sum = make([]int, len(nums)+1)
	for k, v := range nums {
		res.Sum[k+1] = res.Sum[k] + v
	}
	return res
}

func (this *NumArray) SumRange(i int, j int) int {
	return this.Sum[j+1] - this.Sum[i]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
