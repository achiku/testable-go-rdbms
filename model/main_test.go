package model

import (
	"flag"
	"os"
	"testing"

	"log"
)

// TestMain model package setup/teardonw
func TestMain(m *testing.M) {
	flag.Parse()

	// for create/drop schema
	createSchemaCon := "postgres://pgtest@localhost:5432/pgtest?sslmode=disable"
	// for create database objects
	createTableCon := "postgres://store_api_test@localhost:5432/pgtest?sslmode=disable"

	TestingDBTeardown(createSchemaCon)
	if err := TestingDBSetup(createSchemaCon); err != nil {
		log.Fatal(err)
	}
	if err := TestingTableCreate(createTableCon); err != nil {
		log.Fatal(err)
	}
	code := m.Run()
	if err := TestingDBTeardown(createSchemaCon); err != nil {
		log.Fatal(err)
	}
	os.Exit(code)
}
