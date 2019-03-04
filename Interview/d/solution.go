package d

import (
	"fmt"
	"strconv"
	"strings"
)

type User struct {
	Id      int
	StrTime string
	IntTime int
}

func Solution(str [][]string) []int {
	if len(str) <= 5 {
		return []int{}
	}
	users := generateStringToUsers(str)
	users = InsertSort(users)

	//json_byte, _ := json.MarshalIndent(users, "", "  ")
	//fmt.Println(string(json_byte))
	ans := []int{}

	//	判断时间差
	for i := 4; i < len(users); i++ {
		if users[i-4].Id == users[i].Id && (users[i].IntTime-users[i-4].IntTime) > 3600 {
			fmt.Println(users[i], users[i-4])
			ans = append(ans, users[i].Id)
		}
	}
	fmt.Println(ans)
	return RemoveDuplicatesAndEmpty(ans)
}

//	数组去重
func RemoveDuplicatesAndEmpty(a []int) (ret []int) {
	a_len := len(a)
	for i := 0; i < a_len; i++ {
		if i > 0 && a[i-1] == a[i] {
			continue
		}
		ret = append(ret, a[i])
	}
	return
}

//	排序
func InsertSort(users []User) []User {
	var i, j int
	tmp := User{}
	//	根据ID排序
	for i = 1; i < len(users); i++ {
		tmp = users[i]
		for j = i; j > 0 && users[j-1].Id > tmp.Id; j-- {
			users[j] = users[j-1]
		}
		users[j] = tmp
	}
	//	根据时间排序
	for i = 1; i < len(users); i++ {
		tmp = users[i]
		//	需要特殊处理保证ID的稳定性
		for j = i; j > 0 && users[j-1].IntTime > tmp.IntTime && users[j-1].Id > tmp.Id; j-- {
			users[j] = users[j-1]
		}
		users[j] = tmp
	}

	return users
}

//	将字符串转用户对象
func generateStringToUsers(s [][]string) []User {
	users := []User{}
	for _, v := range s {
		id, _ := strconv.Atoi(v[0])
		users = append(users, User{
			Id:      id,
			StrTime: v[1],
			IntTime: generateTime(v[1]),
		})
	}
	return users
}

//	将字符串时间转化成秒
func generateTime(s string) int {
	arr := strings.Split(s, ":")

	hour, _ := strconv.Atoi(arr[0])
	minute, _ := strconv.Atoi(arr[1])
	second, _ := strconv.Atoi(arr[2])

	return hour*3600 + minute*60 + second
}
