package repository

import (
	"database/sql"
	"event-management-api/structs"
)

func GetAllCategory(db *sql.DB) (result []structs.Category, err error) {
    sql := "SELECT id, name, description, created_at FROM categories"
    rows, err := db.Query(sql)
    if err != nil {
        return
    }

    defer rows.Close()
    for rows.Next() {
        var category structs.Category

        err = rows.Scan(&category.ID, &category.Name, &category.Description, &category.CreatedAt)
        if err != nil {
            return
        }

        result = append(result, category)
    }

    return
}

func GetCategory(db *sql.DB, id int) (category structs.Category, err error) {
    sql := "SELECT id, name, description, created_at FROM categories WHERE id = $1"

    err = db.QueryRow(sql, id).Scan(
        &category.ID,
        &category.Name,
        &category.Description,
        &category.CreatedAt,
    )
    
    if err != nil {
        return
    }

    return
}


func InsertCategory(db *sql.DB, category *structs.Category) (err error) {
    sql := "INSERT INTO categories(name, description) VALUES ($1, $2) RETURNING id"

    err = db.QueryRow(sql, category.Name, category.Description).Scan(&category.ID)
    return
}

func UpdateCategory(db *sql.DB, category structs.Category) (err error) {
    sql := "UPDATE categories SET name = $1, description = $2 WHERE id = $3"

    _, err = db.Exec(sql, category.Name, category.Description, category.ID)
    return
}

func DeleteCategory(db *sql.DB, category structs.Category) (err error) {
    sql := "DELETE FROM categories WHERE id = $1"

    _, err = db.Exec(sql, category.ID)
    return
}