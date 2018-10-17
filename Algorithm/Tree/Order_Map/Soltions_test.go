package Order_Map

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func TestNewIntMap(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		t1 := time.Now()
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		testmap := NewIntMap(10000)
		t2 := t1.Second()
		for i := 0; i < 10000; i++ {
			testmap.Insert(r.Intn(10000), r.Intn(10000))
		}

		t1 = time.Now()
		t3 := t1.Second()
		fmt.Println("insert  time  span", t3-t2)
		testmap.Erase(88)
		for i := 0; i < testmap.Size(); i++ {
			k, v, _ := testmap.GetByOrderIndex(i)
			tmp_v := reflect.ValueOf(v)
			fmt.Println("k:", k, "---", "value:", tmp_v)
		}

		t1 = time.Now()
		t4 := t1.Second()
		fmt.Println("range  time  span:", t4-t3)

	})
}
