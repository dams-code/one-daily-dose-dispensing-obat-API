package migrations

import (
	"database/sql"
	"embed"
	"fmt"

	migrate "github.com/rubenv/sql-migrate"
)

//go:embed sql_migrations/*.sql
var DBMigrations embed.FS

func MigrasiDataObat(db *sql.DB) error {

	migrations := migrate.EmbedFileSystemMigrationSource{
		FileSystem: DBMigrations,
		Root:       "sql_migrations",
	}

	typeDatabase := "postgres"

	n, err := migrate.Exec(db, typeDatabase, migrations, migrate.Up)

	if err != nil {
		return err
	}

	fmt.Printf("Migrasi data obat, total : %d termigrasi ke database.", n)

	return nil
}
