package db

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	_ "modernc.org/sqlite"
)

var db *sql.DB

const schema = `
CREATE TABLE IF NOT EXISTS scheduler (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    date CHAR(8) NOT NULL DEFAULT "",
    title VARCHAR(255) NOT NULL DEFAULT "",
    comment TEXT NOT NULL DEFAULT "",
    repeat VARCHAR(128) NOT NULL DEFAULT ""
);

CREATE INDEX IF NOT EXISTS scheduler_date ON scheduler(date);
`

func Init(dbFile string) error {

	_, err := os.Stat(dbFile)
	install := errors.Is(err, os.ErrNotExist)
	db, err = sql.Open("sqlite", dbFile)
	if err != nil {
		return fmt.Errorf("Ошибка в открытии SQL %v", err)
	}

	_, err = db.Exec(schema)
	if err != nil {
		return fmt.Errorf("Ошибка в выполнении SQL-Запроса %v", err)
	}

	if install {
		fmt.Println("Создана новая база данных:", dbFile)
	}
	_, err = db.Exec(`DELETE FROM scheduler WHERE date = ''`)
	if err != nil {
		return fmt.Errorf("ошибка очистки базы: %v", err)
	}

	return nil
}
