package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/caarlos0/env/v11"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

type Config struct {
	Postgres struct {
		Host     string `env:"HOST"`
		Port     uint   `env:"PORT"`
		User     string `env:"USER"`
		Password string `env:"PASSWORD"`
		Db       string `env:"DB"`
	} `envPrefix:"POSTGRES_"`
}

func main() {
	// migrationModules contains every module that has migrations to run,
	// add your migrations here
	var migrationModules = []string{
		filepath.Join("internal", "module_a", "migrations"),
		filepath.Join("internal", "module_b", "migrations"),
	}

	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		log.Fatal(err)
	}

	db, err := connectToDb(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	migrationDir, err := prepareMigrations(migrationModules)
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(migrationDir)

	if err := execute(db, migrationDir); err != nil {
		log.Fatal(err)
	}
}

// prepareMigrations - gathers migrations from multiple modules into migration temp folder
func prepareMigrations(migrationModules []string) (targetDir string, err error) {
	targetDir, err = os.MkdirTemp("", "ddd-migrations-*")
	if err != nil {
		return "", err
	}

	defer func() {
		if err != nil {
			os.RemoveAll(targetDir)
		}
	}()

	for _, migrationModule := range migrationModules {
		dir, err := os.ReadDir(migrationModule)
		if err != nil {
			return "", err
		}

		for _, entry := range dir {
			if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".sql") {
				continue
			}

			sourcePath := filepath.Join(migrationModule, entry.Name())
			destPath := filepath.Join(targetDir, entry.Name())

			if err := copyFile(sourcePath, destPath); err != nil {
				return "", err
			}
		}
	}

	return targetDir, nil

}

func copyFile(source string, dest string) error {
	input, err := os.ReadFile(source)
	if err != nil {
		return fmt.Errorf("failed to read file '%s': %w", source, err)
	}

	err = os.WriteFile(dest, input, 0644)
	if err != nil {
		return fmt.Errorf("failed to write into file '%s': %w", dest, err)
	}

	return nil
}

// execute - execute the migration tool program
func execute(db *sql.DB, dir string) error {
	if len(os.Args) < 2 {
		return fmt.Errorf("user did not specify the command")
	}

	if err := goose.SetDialect("postgres"); err != nil {
		return fmt.Errorf("failed to set SQL dialect: %w", err)
	}

	switch os.Args[1] {
	case "up":
		return goose.Up(db, dir)
	case "down":
		return goose.Down(db, dir)
	case "redo":
		return goose.Redo(db, dir)
	case "status":
		return goose.Status(db, dir)
	case "version":
		return goose.Version(db, dir)
	default:
		return fmt.Errorf("unknown command: %s, Supported commands: up, down, redo, status, version", os.Args[1])
	}
}

// connectToDb - connects to PostgreSQL.
func connectToDb(cfg Config) (*sql.DB, error) {
	dbURI := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.Db,
	)

	db, err := sql.Open("postgres", dbURI)
	if err != nil {
		return nil, fmt.Errorf("failed to open a connection to the db: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to connect to db: %w", err)
	}

	return db, nil
}
