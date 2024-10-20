package structs

import "time"

type Review struct {
    ID        int       `json:"id"`
    Content   string    `json:"content"`
    Rating    int       `json:"rating"`
    UserID    int       `json:"user_id"`
    EventID   int       `json:"event_id"`
    CreatedAt time.Time `json:"created_at"`
}