package Solution

import "fmt"

//	主要是用了 matrix[index / width][index % width] 来转换矩阵
//	上面这个公式会根据列的长度来输出一个矩阵
func matrixReshape(nums [][]int, r int, c int) [][]int {
	n := len(nums)    //	行数
	m := len(nums[0]) //	列数
	if m*n != r*c {
		return nums
	}
	//	初始化一下Slice,因为不知道怎么初始化多维Silce所以只能用这种比较笨的方法了
	reshapedNums := [][]int{}
	for i := 0; i < r; i++ {
		tmp := make([]int, c, c)
		reshapedNums = append(reshapedNums, tmp)
	}
	fmt.Println(r, c, m, n)
	for i := 0; i < r*c; i++ {
		reshapedNums[i/c][i%c] = nums[i/m][i%m]
		fmt.Print(reshapedNums)
		fmt.Println("  [", i/c, "]", "[", i%c, "]", "[", i/m, "]", "[", i%m, "]")
	}
	return reshapedNums
}

func matrixReshape1(nums [][]int, r int, c int) [][]int {
	n := len(nums)    //	行数
	m := len(nums[0]) //	列数
	//
	if m*n != r*c {
		return nums
	}

	reshapedNums := [][]int{}
	index := 0
	for i := 0; i < r; i++ {
		tmp := []int{}
		for j := 0; j < c; j++ {
			tmp = append(tmp, nums[index/n][index%n])
			index++
		}
		reshapedNums = append(reshapedNums, tmp)
	}
	return reshapedNums
}
