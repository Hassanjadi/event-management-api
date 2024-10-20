package repository

import (
	"database/sql"
	"event-management-api/structs"
)

func GetAllReview(db *sql.DB) (result []structs.Review, err error) {
    sql := "SELECT id, content, rating, user_id, event_id, created_at FROM reviews"
    rows, err := db.Query(sql)
    if err != nil {
        return
    }

    defer rows.Close()
    for rows.Next() {
        var review structs.Review

        err = rows.Scan(&review.ID, &review.Content, &review.Rating, &review.UserID, &review.EventID, &review.CreatedAt)
        if err != nil {
            return
        }

        result = append(result, review)
    }

    return
}

func GetReview(db *sql.DB, id int) (review structs.Review, err error) {
	sql := "SELECT id, content, rating, user_id, event_id, created_at FROM reviews WHERE id = $1"

	err = db.QueryRow(sql, id).Scan(&review.ID, &review.Content, &review.Rating, &review.UserID, &review.EventID, &review.CreatedAt)
	return
}

func InsertReview(db *sql.DB, review *structs.Review) (err error) {
	sql := "INSERT INTO reviews(content, rating, user_id, event_id, created_at) VALUES ($1, $2, $3, $4, $5) RETURNING id"

	err = db.QueryRow(sql, review.Content, review.Rating, review.UserID, review.EventID, review.CreatedAt).Scan(&review.ID)
	return
}

func UpdateReview(db *sql.DB, review structs.Review) (err error) {
	sql := "UPDATE reviews SET content = $1, rating = $2, user_id = $3, event_id = $4, created_at = $5 WHERE id = $6"

	_, err = db.Exec(sql, review.Content, review.Rating, review.UserID, review.EventID, review.CreatedAt, review.ID)
	return
}

func DeleteReview(db *sql.DB, id int) (err error) {
	sql := "DELETE FROM reviews WHERE id = $1"

	_, err = db.Exec(sql, id)
	return
}
