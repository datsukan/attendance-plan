package main

import (
	"time"

	"github.com/guregu/dynamo"
)

const TableNameUser = "AttendancePlan_User"

type User struct {
	ID        string    `dynamo:"ID,hash"`
	Email     string    `dynamo:"Email" index:"Email-index,hash"`
	Password  string    `dynamo:"Password"`
	Name      string    `dynamo:"Name"`
	Enabled   bool      `dynamo:"Enabled"`
	CreatedAt time.Time `dynamo:"CreatedAt"`
	UpdatedAt time.Time `dynamo:"UpdatedAt"`
}

func (u User) Up(db *dynamo.DB) error {
	tables, err := db.ListTables().All()
	if err != nil {
		return err
	}

	for _, table := range tables {
		if table == TableNameUser {
			return nil
		}
	}

	return db.CreateTable(TableNameUser, User{}).Run()
}

func (u User) Down(db *dynamo.DB) error {
	return db.Table(TableNameUser).DeleteTable().Run()
}
