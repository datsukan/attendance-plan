package main

import (
	"fmt"
	"os"

	"github.com/guregu/dynamo"
)

func main() {
	execType, err := GetArgs()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	db := NewDB()

	switch execType {
	case ExecTypeUp:
		if err := up(db); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	case ExecTypeDown:
		if err := down(db); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	fmt.Println("successful")
}

func up(db *dynamo.DB) error {
	schedule := Schedule{}
	if err := schedule.Up(db); err != nil {
		return err
	}

	user := User{}
	if err := user.Up(db); err != nil {
		return err
	}

	return nil
}

func down(db *dynamo.DB) error {
	schedule := Schedule{}
	if err := schedule.Down(db); err != nil {
		return err
	}

	user := User{}
	if err := user.Down(db); err != nil {
		return err
	}

	return nil
}
