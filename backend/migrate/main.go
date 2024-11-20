package main

import (
	"fmt"
	"os"
)

func main() {
	execType, err := GetArgs()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	db := NewDB()

	s := Schedule{}

	switch execType {
	case ExecTypeUp:
		if err := s.Up(db); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	case ExecTypeDown:
		if err := s.Down(db); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	fmt.Println("successful")
}
