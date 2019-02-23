package main

func Stop(stop chan<- bool) {
	close(stop)
}
