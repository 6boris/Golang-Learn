package Sort

import (
	"fmt"
	"math/rand"
	"testing"
)

const MIN_NUM = 1
const MAX_NUM = 100
const ARRAY_NUM = 10000000

type Data struct {
	Input  []int
	Output []int
}

var sourceData Data

// 初始化入口
func init() {
	sourceData = initData()
}

// 初始化数据
func initData() Data {
	sourceData := new(Data)
	for i := 0; i < ARRAY_NUM; i++ {
		sourceData.Input = append(sourceData.Input, MIN_NUM+rand.Intn(MAX_NUM-MIN_NUM))
		sourceData.Output = append(sourceData.Output, sourceData.Input[i])
	}
	sourceData.Output = QuickSort(sourceData.Input)

	return *sourceData
}

// 获取一个测试数据，获取的每个数据互相不影响
func getData() Data {
	var res = new(Data)
	for i := 0; i < len(sourceData.Input); i++ {
		res.Input = append(res.Input, sourceData.Input[i])
		res.Output = append(res.Output, sourceData.Output[i])
	}
	return *res
}

func isEqualData(data1 []int, data2 []int) bool {
	if len(data1) != len(data2) {
		return false
	}

	for i := 0; i < len(data1); i++ {
		if data1[i] != data2[i] {
			return false
		}
	}
	return true
}

func TestSort(t *testing.T) {

	fmt.Println(MIN_NUM + rand.Intn(MAX_NUM-MIN_NUM))

	//冒泡排序测试
	t.Run("TestBubbleSort...", func(t *testing.T) {
		var tmp = getData()
		got := BubbleSort(tmp.Input)
		want := tmp.Output

		if !isEqualData(got, want) {
			t.Error("Got:", got, "Want:", want)
		}
	})

	// 插入排序测试
	t.Run("TestInsertSort...", func(t *testing.T) {
		var tmp = getData()
		got := InsertSort(tmp.Input)
		want := tmp.Output

		if !isEqualData(got, want) {
			t.Error("Got:", got, "Want:", want)
		}
	})

	// 选择排序测试
	t.Run("TestSelectSort...", func(t *testing.T) {
		var tmp = getData()
		got := SelectSort(tmp.Input)
		want := tmp.Output

		if !isEqualData(got, want) {
			t.Error("Got:", got, "Want:", want)
		}
	})

	// 快速排序测试
	t.Run("TestQuickSort...", func(t *testing.T) {
		var tmp = getData()
		got := QuickSort(tmp.Input)
		want := tmp.Output

		if !isEqualData(got, want) {
			t.Error("Got:", got, "Want:", want)
		}

	})

	// 归并排序测试
	t.Run("TestMergeSort...", func(t *testing.T) {
		var tmp = getData()
		got := MergeSort(tmp.Input)
		want := tmp.Output

		if !isEqualData(got, want) {
			t.Error("Got:", got, "Want:", want)
		}
	})

	// 堆排序测试
	t.Run("TestHeapSort...", func(t *testing.T) {
		var tmp = getData()
		got := HeapSort(tmp.Input)
		want := tmp.Output

		if !isEqualData(got, want) {
			t.Error("Got:", got, "Want:", want)
		}
	})

	// 希尔排序测试
	t.Run("TestShellSort...", func(t *testing.T) {
		var tmp = getData()
		got := ShellSort(tmp.Input)
		want := tmp.Output

		if !isEqualData(got, want) {
			t.Error("Got:", got, "Want:", want)
		}
	})

	// 鸡尾酒排序测试
	t.Run("TestCocktailShakerSort...", func(t *testing.T) {
		var tmp = getData()
		got := CocktailShakerSort(tmp.Input)
		want := tmp.Output

		if !isEqualData(got, want) {
			t.Error("Got:", got, "Want:", want)
		}
	})

	// 梳排序测试
	//t.Run("TestCombSort...", func(t *testing.T) {
	//	var tmp = getData()
	//	got := CombSort(tmp.Input)
	//	want := tmp.Output
	//
	//	if !isEqualData(got, want) {
	//		t.Error("Got:", got, "Want:", want)
	//	}
	//})

	// 计数排序测试
	t.Run("TestCountingSort...", func(t *testing.T) {
		var tmp = getData()
		got := CountingSort(tmp.Input)
		want := tmp.Output

		if !isEqualData(got, want) {
			t.Error("Got:", got, "Want:", want)
		}
	})

	// 地精排序测试
	t.Run("TestGnomeSort...", func(t *testing.T) {
		var tmp = getData()
		got := GnomeSort(tmp.Input)
		want := tmp.Output

		if !isEqualData(got, want) {
			t.Error("Got:", got, "Want:", want)
		}
	})

	// 奇偶排序测试
	t.Run("TestOddEvenSort...", func(t *testing.T) {
		var tmp = getData()
		got := OddEvenSort(tmp.Input)
		want := tmp.Output

		if !isEqualData(got, want) {
			t.Error("Got:", got, "Want:", want)
		}
	})

}
