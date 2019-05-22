package main

import (
	"fmt"
	"log"
)

func main()  {

	age := 27
	rows, err := db.Query("SELECT name FROM users WHERE age=?", age)
	if err != nil {
		log.Fatal(err)
	}
	defer ___ // A
	for ___ { //B
		var name string
		if err := ___; err != nil { //C
			log.Fatal(err)
		}
		fmt.Printf("%s is %d\n", name, age)
	}
	if err := ___; err != nil { //D
		log.Fatal(err)
	}

}

func main2()  {
	age := 27
	rows, err := db.Query("SELECT name FROM users WHERE age=?", age)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s is %d\n", name, age)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

}