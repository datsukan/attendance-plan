package main

import (
	"time"

	"github.com/guregu/dynamo"
)

const TableNameSubject = "AttendancePlan_Subject"

type Subject struct {
	ID        string    `dynamo:"ID,hash"`
	UserID    string    `dynamo:"UserID" index:"UserID-index,hash"`
	Name      string    `dynamo:"Name"`
	Color     string    `dynamo:"Color"`
	CreatedAt time.Time `dynamo:"CreatedAt"`
	UpdatedAt time.Time `dynamo:"UpdatedAt"`
}

func (s Subject) Up(db *dynamo.DB) error {
	tables, err := db.ListTables().All()
	if err != nil {
		return err
	}

	for _, table := range tables {
		if table == TableNameSubject {
			return nil
		}
	}

	return db.CreateTable(TableNameSubject, Subject{}).Run()
}

func (s Subject) Down(db *dynamo.DB) error {
	return db.Table(TableNameSubject).DeleteTable().Run()
}
