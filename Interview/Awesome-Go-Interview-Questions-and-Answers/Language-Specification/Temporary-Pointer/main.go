package main

const N = 3

func main0() {
	m := make(map[int]*int)

	for i := 0; i < N; i++ {
		m[i] = &i //A
	}

	for _, v := range m {
		print(*v)
	}
}

func main1() {
	m := make(map[int]*int)

	for i := 0; i < N; i++ {
		j := i
		m[i] = &j //A
	}

	for _, v := range m {
		print(*v)
	}
}

func main() {
	m := make(map[int]int)

	for i := 0; i < N; i++ {
		m[i] = i //A
	}

	for _, v := range m {
		print(v)
	}
}
