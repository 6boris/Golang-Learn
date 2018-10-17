package Order_Map

/*
Golang  map实现原理是hash  map（核心元素是桶，key通过哈希算法被归入不同的bucket中），
key是无序的，很多应用场景可能需要map  key有序（例如交易所订单撮合），C++  的stl  map
实现了key有序，实际上是TreeMap是基于树（红黑树）的实现方式，即添加到一个有序列表，在O(log  n)
的复杂度内通过key值找到value，优点是空间要求低，但在时间上不如HashMap。

闲来用go  map  +  slice切片，实现了一套key有序map数据结构，就是空间换时间的玩法，
实质是map  负责存k  v，  slice负责维护k的有序索引位置(查找key采用的是2分法)，实现后赠改删时间负责度是  O(log2n),  。

优化的一点思考：实际上主要就是在slice上维护k位置时的增改删费操作，这时候我们可根据具体应用在2分查找上下点文章。
例如可能所存的数据结构频繁操作的节点只有前面一部分，这时候我们可以加个逻辑，操作时slice时先2查找
slice子集（例如头部热点），这样可能很多增改删操作在第一时间就解决了，整体性能会有很大提升，
最好根据应用场景来具体分析解决。下面给出代码。

*/

func findIndexByBinarySearch(s []int, k int) (int, bool) {
	lo, hi := 0, len(s)
	var m int
	max := len(s)
	if max == 0 {
		return 0, false
	}
	res := false
	for lo <= hi {
		m = (lo + hi) >> 1
		if m == 0 && s[0] > k {
			return 0, res
		}
		if m == max-1 && s[max-1] < k {
			return m + 1, res
		}
		if s[m] < k && s[m+1] > k {
			return m + 1, res
		}
		if s[m] > k && s[m-1] < k {
			return m, res
		}
		if s[m] < k {
			lo = m + 1
		} else if s[m] > k {
			hi = m - 1
		} else {
			return m, true
		}
	}
	return -1, false
}

type Int_Map struct {
	dataMap  map[int]interface{}
	keyArray []int
}

func NewIntMap(cap int) *Int_Map {
	return &Int_Map{
		dataMap:  make(map[int]interface{}),
		keyArray: make([]int, 0, cap),
	}
}

func (m *Int_Map) Exists(key int) bool {
	_, exists := m.dataMap[key]
	return exists
}

func (m *Int_Map) Insert(key int, data interface{}) bool {
	m.dataMap[key] = data
	index, res := findIndexByBinarySearch(m.keyArray, key)
	if index == -1 {
		return false
	}
	if res == true { //存在则直接返回
		return true
	}
	if len(m.keyArray) == 0 {
		m.keyArray = append(m.keyArray, key)
		return true
	}
	//追加末尾
	if index >= len(m.keyArray) {
		m.keyArray = append(m.keyArray[0:], []int{key}...)
	} else if index == 0 { //追加头部
		m.keyArray = append([]int{key}, m.keyArray[:len(m.keyArray)]...)
	} else { //插入
		rear := append([]int{}, m.keyArray[index:]...)
		m.keyArray = append(m.keyArray[0:index], key)
		m.keyArray = append(m.keyArray, rear...)
	}
	return true
}

func (m *Int_Map) Erase(key int) {
	if !m.Exists(key) {
		return
	}
	index, res := findIndexByBinarySearch(m.keyArray, key)
	if res == false {
		return
	}
	delete(m.dataMap, key)
	if index == 0 {
		m.keyArray = m.keyArray[1:]
	} else if index == len(m.keyArray) {
		m.keyArray = m.keyArray[:len(m.keyArray)-2]
	} else {
		m.keyArray = append(m.keyArray[:index], m.keyArray[index+1:]...)
	}

}

func (m *Int_Map) Size() int {
	return len(m.keyArray)
}

func (m *Int_Map) GetByOrderIndex(index int) (int, interface{}, bool) {
	if index < 0 || index >= len(m.keyArray) {
		return -1, nil, false
	}
	key := m.keyArray[index]
	return key, m.dataMap[key], true
}
