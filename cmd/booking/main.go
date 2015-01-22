package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"

	"github.com/cmdrkeene/booking"
	"github.com/facebookgo/inject"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	flag.Set("logtostderr", "true")
	flag.Parse()

	// Database
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}

	// Domain
	var calendar booking.Calendar
	var guestbook booking.Guestbook
	var ledger booking.Ledger
	var register booking.Register
	var server booking.Server

	// Dependency Injection
	var g inject.Graph
	err = g.Provide(
		&inject.Object{Value: db},
		&inject.Object{Value: &calendar},
		&inject.Object{Value: &guestbook},
		&inject.Object{Value: &ledger},
		&inject.Object{Value: &register},
		&inject.Object{Value: &server},
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if err = g.Populate(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Start
	server.ListenAndServe(":3000")
}
