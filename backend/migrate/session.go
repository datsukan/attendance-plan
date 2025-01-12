package main

import (
	"time"

	"github.com/guregu/dynamo"
)

const TableNameSession = "AttendancePlan_Session"

type Session struct {
	ID        string    `dynamo:"ID,hash"`
	UserID    string    `dynamo:"UserID" index:"UserID-index,hash"`
	ExpiresAt time.Time `dynamo:"ExpiresAt"`
}

func (s Session) Up(db *dynamo.DB) error {
	db.Table(TableNameSession).DeleteTable().Run()
	return db.CreateTable(TableNameSession, Session{}).Run()
}

func (s Session) Down(db *dynamo.DB) error {
	return db.Table(TableNameSession).DeleteTable().Run()
}
