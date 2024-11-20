package main

import (
	"errors"
	"os"
)

type ExecType string

const (
	ExecTypeUp   ExecType = "up"
	ExecTypeDown ExecType = "down"
)

func GetArgs() (ExecType, error) {
	if len(os.Args) < 2 {
		return "", errors.New("引数を1つ指定してください")
	}

	switch os.Args[1] {
	case "up":
		return ExecTypeUp, nil
	case "down":
		return ExecTypeDown, nil
	default:
		return "", errors.New("引数はupかdownを指定してください")
	}
}
