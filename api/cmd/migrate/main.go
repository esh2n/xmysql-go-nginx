package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

const (
	Source = "file://./db/migrations/"
)

var (
	Command  = flag.String("e", "", "up or down")
	Database string
)

var ExecCommands = map[string]string{
	"up":   "up sql",
	"down": "down sql",
}

func init() {
	dbName := os.Getenv("DB_NAME")
	if os.Getenv("APP_MODE") == "test" {
		dbName = os.Getenv("DB_NAME_TEST")
	}

	Database = fmt.Sprintf("mysql://%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		dbName)
}

func main() {
	flag.Parse()

	if len(*Command) < 1 {
		fmt.Println("expected more than 1 argument")
		os.Exit(1)
		return
	}

	m, err := migrate.New(Source, Database)
	if err != nil {
		fmt.Println("err", err)
		return
	}
	version, dirty, err := m.Version()

	fmt.Println("command: exec", *Command)
	applyQuery(m, version, dirty)
}

func applyQuery(m *migrate.Migrate, version uint, dirty bool) {
	var err error
	switch *Command {
	case "up":
		err = m.Up()
	case "down":
		err = m.Down()
	default:
		fmt.Println("\nerror: invalid command '" + *Command + "'\n")
		os.Exit(1)
	}

	if err != nil {
		fmt.Println("err", err)
		os.Exit(1)
	} else {
		fmt.Println("success:", *Command+"\n")
	}
}
