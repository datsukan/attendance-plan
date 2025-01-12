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
	CreatedAt time.Time `dynamo:"CreatedAt"`
	UpdatedAt time.Time `dynamo:"UpdatedAt"`
}

func (u User) Up(db *dynamo.DB) error {
	db.Table(TableNameUser).DeleteTable().Run()
	return db.CreateTable(TableNameUser, User{}).Run()
}

func (u User) Down(db *dynamo.DB) error {
	return db.Table(TableNameUser).DeleteTable().Run()
}
