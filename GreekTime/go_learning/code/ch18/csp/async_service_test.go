package concurrency

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond * 50)
	return "Done"
}

func otherTask() {
	fmt.Println("working on something else")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("Task is done.")
}

func TestService(t *testing.T) {
	fmt.Println(service())
	otherTask()
}

func AsyncService() chan string {
	retCh := make(chan string, 1)
	//retCh := make(chan string, 1)
	go func() {
		ret := service()
		fmt.Println("returned result.")
		retCh <- ret
		fmt.Println("service exited.")
	}()
	return retCh
}

//
func TestAsynService(t *testing.T) {
	retCh := AsyncService()
	otherTask()
	fmt.Println(<-retCh)
	time.Sleep(time.Second * 1)
}

func S() string {
	time.Sleep(time.Millisecond * 500)
	return "OK"
}

func AsyncS() chan string {
	//resChan := make(chan string)
	resChan := make(chan string, 1)
	res := S()
	go func() {
		fmt.Println("start do srv")
		resChan <- res
		fmt.Println("end do srv")

	}()
	return resChan
}

func task() {
	fmt.Println("start do task")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("end do task")

}

func TestS(t *testing.T) {
	ret := AsyncS()
	fmt.Println(<-ret)
	task()

	for {
		select {
		case res := <-ret:
			fmt.Println(res)
		case <-time.After(time.Second * 1):
			t.Error("time out")
		}
	}
}
