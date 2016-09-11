package model

import (
	"database/sql"
	"io/ioutil"
	"testing"

	txdb "github.com/achiku/pgtxdb"
	_ "github.com/lib/pq" // postgres
	"github.com/pkg/errors"
)

func init() {
	txdb.Register("txdb", "postgres", "postgres://store_api_test@localhost:5432/pgtest?sslmode=disable")
}

// CREATE USER pgtest; -- superuser
// CREATE USER store_api; -- this user is for development
// CREATE SCHEMA store_api AUTHORIZATION store_api;
// CREATE USER store_api_test; -- this user is for test

// TestingDBSetup set up test schema
func TestingDBSetup(conStr string) error {
	con, err := sql.Open("postgres", conStr)
	if err != nil {
		return errors.Wrap(err, "failed to open connection")
	}
	defer con.Close()

	_, err = con.Exec("CREATE SCHEMA store_api_test AUTHORIZATION store_api_test")
	if err != nil {
		return errors.Wrap(err, "failed to create test schema")
	}
	return nil
}

// TestingTableCreate create test tables
func TestingTableCreate(conStr string) error {
	con, err := sql.Open("postgres", conStr)
	if err != nil {
		return errors.Wrap(err, "failed to open connection")
	}
	defer con.Close()

	ddl, err := ioutil.ReadFile("./ddl.sql")
	if err != nil {
		return errors.Wrap(err, "failed to read ddl.sql")
	}

	_, err = con.Exec(string(ddl))
	if err != nil {
		return errors.Wrap(err, "failed to execute ddl.sql")
	}
	return nil
}

// TestingDBTeardown drop test schema
func TestingDBTeardown(conStr string) error {
	con, err := sql.Open("postgres", conStr)
	if err != nil {
		return errors.Wrap(err, "failed to open connection")
	}
	defer con.Close()

	_, err = con.Exec("DROP SCHEMA store_api_test CASCADE")
	if err != nil {
		return errors.Wrap(err, "failed to drop schema")
	}
	return nil
}

// TestSetupTx create tx and cleanup func for test
func TestSetupTx(t *testing.T) (Txer, func()) {
	db, err := sql.Open("txdb", "dummy")
	if err != nil {
		t.Fatal(err)
	}
	tx, err := db.Begin()
	if err != nil {
		t.Fatal(err)
	}

	cleanup := func() {
		tx.Rollback()
		db.Close()
	}
	return tx, cleanup
}

// TestSetupDB create db and cleanup func for test
func TestSetupDB(t *testing.T) (DBer, func()) {
	db, err := sql.Open("txdb", "dummy")
	if err != nil {
		t.Fatal(err)
	}

	cleanup := func() {
		db.Close()
	}
	return db, cleanup
}
