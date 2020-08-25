package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/dogmatiq/dodeca/logging"
	"github.com/dogmatiq/example"
	"github.com/dogmatiq/example/cmd/bank/ui"
	"github.com/dogmatiq/example/database"
	"github.com/dogmatiq/infix"
	infixsql "github.com/dogmatiq/infix/persistence/provider/sql"
	"github.com/dogmatiq/projectionkit/sql/sqlite"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/sync/errgroup"
)

func main() {
	rand.Seed(time.Now().Unix())

	if err := run(); err != nil {
		if !errors.Is(err, context.Canceled) {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

func run() error {
	// Create a new UI.
	u := ui.New()

	// Run the engine and the UI in parallel.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		// Ensure engine is always stopped when UI stops, even if it exited
		// cleanly.
		defer cancel()
		return u.Run(ctx)
	})

	g.Go(func() error {
		return runEngine(ctx, u)
	})

	return g.Wait()
}

func runEngine(ctx context.Context, u *ui.UI) error {
	// Open the SQLite database.
	db, err := sql.Open("sqlite3", "file:artifacts/bank.sqlite?mode=rwc")
	if err != nil {
		return err
	}
	defer db.Close()

	// Initialize the example application.
	app, err := example.NewApp(db)
	if err != nil {
		return err
	}

	// Determine the database driver to use (always SQLite in this case).
	driver, err := infixsql.NewDriver(db)
	if err != nil {
		return err
	}

	// Create schema elements for infix, projection kit, and the bank example.
	createSchema(ctx, db, driver, u.Logger)

	// Initialize the Infix engine itself.
	e := infix.New(
		app,
		infix.WithPersistence(
			&infixsql.Provider{
				DB:     db,
				Driver: driver,
			},
		),
		infix.WithLogger(u.Logger),
	)

	// Run the engine until the context is canceled or an error occurs.
	err = e.Run(ctx)
	logging.Log(u.Logger, "infix has stopped: %s", err)

	return err
}

// createSchema creates the SQL schema needed by the dogmatiq/infix,
// dogmatiq/projectionkit, and the bank example itself.
//
// Schema creation failures are logged without causing the UI to exit.
func createSchema(
	ctx context.Context,
	db *sql.DB,
	d infixsql.Driver,
	l logging.Logger,
) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// Create the schema for dogmatiq/infix.
	if err := d.CreateSchema(ctx, db); err != nil {
		logging.DebugString(l, err.Error())
	}

	// Create the schema for dogmatiq/projectionkit.
	if err := sqlite.CreateSchema(ctx, db); err != nil {
		logging.DebugString(l, err.Error())
	}

	// Create the schema for dogmatiq/example.
	if err := database.CreateSchema(ctx, db); err != nil {
		logging.DebugString(l, err.Error())
	}
}
