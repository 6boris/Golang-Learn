package Solution

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	//	从右上角开始查找
	row := 0
	col := len(matrix[0]) - 1
	for row < len(matrix) && col >= 0 {
		//	找到目标
		if target == matrix[row][col] {
			return true
			//	当前数比目标大，坐标向左移动
		} else if target < matrix[row][col] {
			col--
			//	当前数比目标小，坐标向下移动
		} else {
			row++
		}
	}
	return false
}
