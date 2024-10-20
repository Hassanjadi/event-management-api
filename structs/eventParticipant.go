package structs

import "time"

type EventParticipant struct {
    ID          int       `json:"id"`
    EventID     int       `json:"event_id"`
    UserID      int       `json:"user_id"`
    RegisteredAt time.Time `json:"registered_at"`
}
