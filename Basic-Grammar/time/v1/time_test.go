package v1

import (
	"fmt"
	"math"
	"testing"
	"time"
)

func TestTime1(t *testing.T) {
	fmt.Println(time.Now())
	fmt.Println(time.Now().Unix())
	fmt.Println(time.Now().Nanosecond())
	fmt.Println(time.Now().Local())
	fmt.Println(time.Now().Local())
	fmt.Println(time.Now().Format(time.RFC850))
	fmt.Println(time.Now().Format(time.RFC3339Nano))
	fmt.Println(int64(time.Second / time.Millisecond))

	seconds := 10
	fmt.Println(time.Duration(seconds) * time.Second) // prints 10s
	fmt.Println(math.MaxInt32)
	fmt.Println(int(time.Duration(time.Hour)))
	fmt.Println(math.MaxInt64 / int(time.Duration(time.Hour)) / 24 / 360)
	fmt.Println(math.Sqrt(math.MaxInt64))

	fmt.Println(time.Unix(int64(0), int64(time.Duration(time.Hour))).Format("2006-01-02 15:04:05"))

	fmt.Println(time.Unix(int64(math.MaxInt64), 0).Format("2006-01-02 15:04:05"))
	fmt.Println(time.Unix(int64(math.MaxInt64/time.Duration(time.Second)), 0).Format("2006-01-02 15:04:05"))
}

var c chan int

func handle(int) {
}
func TestTime2(t *testing.T) {

	select {
	case m := <-c:
		handle(m)
	case <-time.After(10 * time.Second):
		fmt.Println("timed out")
	}
}

func Test_Sleep(t *testing.T) {
	for i := 0; i < 3; i++ {
		fmt.Println("begin",time.Now().Format("2006-01-02 15:04:05"))
		fmt.Println("Do something 1s")
		time.Sleep(time.Second * 1)
		fmt.Println("end", time.Now().Format("2006-01-02 15:04:05"))
		time.Sleep(time.Second * 5)
	}
}

func Test_Tick(t *testing.T) {
	t1 := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-t1.C:
			fmt.Println("begin", time.Now().Format("2006-01-02_15:04:05"))
			fmt.Println("Do something 1s")
			time.Sleep(time.Second * 1)
			fmt.Println("end", time.Now().Format("2006-01-02_15:04:05"))
		}
	}
}

