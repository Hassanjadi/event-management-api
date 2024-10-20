package repository

import (
	"database/sql"
	"event-management-api/structs"
)

func GetAllParticipant(db *sql.DB) (result []structs.EventParticipant, err error) {
    sql := "SELECT id, event_id, user_id, registered_at FROM event_participants"
    rows, err := db.Query(sql)
    if err != nil {
        return
    }

    defer rows.Close()
    for rows.Next() {
        var participant structs.EventParticipant

        err = rows.Scan(&participant.ID, &participant.EventID, &participant.UserID, &participant.RegisteredAt)
        if err != nil {
            return
        }

        result = append(result, participant)
    }

    return
}

func GetParticipant(db *sql.DB, id int) (participant structs.EventParticipant, err error) {
	sql := "SELECT id, user_id, event_id, registered_at FROM event_participants WHERE id = $1"

	err = db.QueryRow(sql, id).Scan(&participant.ID, &participant.UserID, &participant.EventID, &participant.RegisteredAt)
	return
}

func InsertParticipant(db *sql.DB, participant *structs.EventParticipant) (err error) {
	sql := "INSERT INTO event_participants(user_id, event_id) VALUES ($1, $2) RETURNING id"

	err = db.QueryRow(sql, participant.UserID, participant.EventID).Scan(&participant.ID)
	return
}

func UpdateParticipant(db *sql.DB, participant structs.EventParticipant) (err error) {
	sql := "UPDATE event_participants SET user_id = $1, event_id = $2, registered_at = $3 WHERE id = $4"

	_, err = db.Exec(sql, participant.UserID, participant.EventID, participant.RegisteredAt, participant.ID)
	return
}

func DeleteParticipant(db *sql.DB, id int) (err error) {
	sql := "DELETE FROM event_participants WHERE id = $1"

	_, err = db.Exec(sql, id)
	return
}
