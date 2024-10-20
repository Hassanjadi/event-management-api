package structs

import "time"

type Event struct {
    ID          int       `json:"id"`
    Title       string    `json:"title"`
    Description string    `json:"description"`
    Date        string    `json:"date"`
    Location    string    `json:"location"`
    CategoryID  int       `json:"category_id"`
    UserID      int       `json:"user_id"`
    CreatedAt   time.Time `json:"created_at"`
}