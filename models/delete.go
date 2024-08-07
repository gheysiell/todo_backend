package models

import "github.com/gheysiell/todo_backend/db"

func Delete(id int64) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	res, err := conn.Exec(`DELETE FROM todo WHERE id=?`, id)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
