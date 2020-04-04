package main

import (
    "context"
    gosql "database/sql"
    "fmt"
    "github.com/cockroachdb/cockroach/pkg/server"
    "github.com/cockroachdb/cockroach/pkg/sql/tests"
    "github.com/cockroachdb/cockroach/pkg/testutils/serverutils"
    "testing"
)

func main() {
    var testing testing.T
    t := &testing

    serverutils.InitTestServerFactory(server.TestServerFactory)

    params, _ := tests.CreateTestServerParams()
    s, db, _ := serverutils.StartServer(t, params)
    defer s.Stopper().Stop(context.TODO())

    if _, err := db.Exec(`
		CREATE DATABASE d;
		SET DATABASE = d;
		CREATE TABLE t (c1 INT, c2 INT, c3 INT);
	`); err != nil {
        panic(err)
    }

    testCases := []struct {
        exec   string
        query  string
        expect gosql.NullString
    }{
        {
            `COMMENT ON COLUMN t.c1 IS 'foo'`,
            `SELECT col_description(attrelid, attnum) FROM pg_attribute WHERE attrelid = 't'::regclass AND attname = 'c1'`,
            gosql.NullString{String: `foo`, Valid: true},
        },
        {
            `TRUNCATE t`,
            `SELECT col_description(attrelid, attnum) FROM pg_attribute WHERE attrelid = 't'::regclass AND attname = 'c1'`,
            gosql.NullString{String: `foo`, Valid: true},
        },
        {
            `ALTER TABLE t RENAME COLUMN c1 TO c1_1`,
            `SELECT col_description(attrelid, attnum) FROM pg_attribute WHERE attrelid = 't'::regclass AND attname = 'c1_1'`,
            gosql.NullString{String: `foo`, Valid: true},
        },
        {
            `COMMENT ON COLUMN t.c1_1 IS NULL`,
            `SELECT col_description(attrelid, attnum) FROM pg_attribute WHERE attrelid = 't'::regclass AND attname = 'c1_1'`,
            gosql.NullString{Valid: false},
        },
        {
            `COMMENT ON COLUMN t.c3 IS 'foo'`,
            `SELECT col_description(attrelid, attnum) FROM pg_attribute WHERE attrelid = 't'::regclass AND attname = 'c3'`,
            gosql.NullString{String: `foo`, Valid: true},
        },
        {
            `ALTER TABLE t DROP COLUMN c2`,
            `SELECT col_description(attrelid, attnum) FROM pg_attribute WHERE attrelid = 't'::regclass AND attname = 'c3'`,
            gosql.NullString{String: `foo`, Valid: true},
        },
    }

    for _, tc := range testCases {
        if _, err := db.Exec(tc.exec); err != nil {
            panic(err)
        }

        row := db.QueryRow(tc.query)
        var comment gosql.NullString
        if err := row.Scan(&comment); err != nil {
            panic(err)
        }
        if tc.expect != comment {
            err := fmt.Sprintf("expected comment %v, got %v", tc.expect, comment)
            panic(err)
        }
    }
}
