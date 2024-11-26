package main

import (
	"time"

	"github.com/guregu/dynamo"
)

const TableNameSchedule = "AttendancePlan_Schedule"

type Schedule struct {
	ID        string    `dynamo:"ID,hash"`
	UserID    string    `dynamo:"UserID" index:"UserID-index,hash"`
	Name      string    `dynamo:"Name"`
	StartsAt  time.Time `dynamo:"StartsAt" index:"UserID-index,range"`
	EndsAt    time.Time `dynamo:"EndsAt"`
	Color     string    `dynamo:"Color"`
	Type      string    `dynamo:"Type"`
	CreatedAt time.Time `dynamo:"CreatedAt"`
	UpdatedAt time.Time `dynamo:"UpdatedAt"`
}

func (s Schedule) Up(db *dynamo.DB) error {
	db.Table(TableNameSchedule).DeleteTable().Run()
	return db.CreateTable(TableNameSchedule, Schedule{}).Run()
}

func (s Schedule) Down(db *dynamo.DB) error {
	return db.Table(TableNameSchedule).DeleteTable().Run()
}
