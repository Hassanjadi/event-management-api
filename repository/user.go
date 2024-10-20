package repository

import (
	"database/sql"
	"event-management-api/structs"
)

func GetAllUser(db *sql.DB) (result []structs.User, err error) {
    sql := "SELECT id, name, email, password, created_at FROM users"
    rows, err := db.Query(sql)
    if err != nil {
        return
    }

    defer rows.Close()
    for rows.Next() {
        var user structs.User

        err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
        if err != nil {
            return
        }

        result = append(result, user)
    }

    return
}

func GetUser(db *sql.DB, id int) (user structs.User, err error) {
    sql := "SELECT id, name, email, password, created_at FROM users WHERE id = $1"

    err = db.QueryRow(sql, id).Scan(
        &user.ID,
        &user.Name,
        &user.Email,
        &user.Password,
        &user.CreatedAt,
    )
    
    if err != nil {
        return
    }

    return
}


func InsertUser(db *sql.DB, user *structs.User) (err error) {
    sql := "INSERT INTO users(name, email, password) VALUES ($1, $2, $3) RETURNING id"

    err = db.QueryRow(sql, user.Name, user.Email, user.Password).Scan(&user.ID)
    return
}

func UpdateUser(db *sql.DB, user structs.User) (err error) {
    sql := "UPDATE users SET name = $1, email = $2, password = $3 WHERE id = $4"

    _, err = db.Exec(sql, user.Name, user.Email, user.Password, user.ID)
    return
}

func DeleteUser(db *sql.DB, user structs.User) (err error) {
    sql := "DELETE FROM users WHERE id = $1"

    _, err = db.Exec(sql, user.ID)
    return
}
