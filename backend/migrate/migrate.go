package main

import "github.com/guregu/dynamo"

type Migrate interface {
	Up(db *dynamo.DB) error
	Down(db *dynamo.DB) error
}
