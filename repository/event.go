package repository

import (
	"database/sql"
	"event-management-api/structs"
)

func GetAllEvent(db *sql.DB) (result []structs.Event, err error) {
    sql := "SELECT id, title, description, date, location, category_id, user_id, created_at FROM events"
    rows, err := db.Query(sql)
    if err != nil {
        return
    }

    defer rows.Close()
    for rows.Next() {
        var event structs.Event

        err = rows.Scan(&event.ID, &event.Title, &event.Description, &event.Date, &event.Location, &event.CategoryID, &event.UserID, &event.CreatedAt)
        if err != nil {
            return
        }

        result = append(result, event)
    }

    return
}

func GetEvent(db *sql.DB, id int) (event structs.Event, err error) {
    sql := "SELECT id, title, description, date, location, category_id, user_id, created_at FROM events WHERE id = $1"

    err = db.QueryRow(sql, id).Scan(
        &event.ID,
        &event.Title,
        &event.Description,
        &event.Date,
        &event.Location,
        &event.CategoryID,
        &event.UserID,
        &event.CreatedAt,
    )
    
    if err != nil {
        return
    }

    return
}


func InsertEvent(db *sql.DB, event *structs.Event) (err error) {
    sql := "INSERT INTO events(title, description, date, location, category_id, user_id, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id"

    err = db.QueryRow(sql, event.Title, event.Description, event.Date, event.Location, event.CategoryID, event.UserID, event.CreatedAt).Scan(&event.ID)
    return
}

func UpdateEvent(db *sql.DB, event structs.Event) (err error) {
    sql := "UPDATE events SET title = $1, description = $2, date = $3, location = $4, category_id = $5, user_id = $6, created_at = $7 WHERE id = $8"

    _, err = db.Exec(sql, event.Title, event.Description, event.Date, event.Location, event.CategoryID, event.UserID, event.CreatedAt, event.ID)
    return
}

func DeleteEvent(db *sql.DB, event structs.Event) (err error) {
    sql := "DELETE FROM events WHERE id = $1"

    _, err = db.Exec(sql, event.ID)
    return
}
