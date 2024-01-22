package response

import (
	"database/sql"
	"time"
)

type TodoDataResponse struct {
	PersonID  string       `json:"personid"`
	Title     string       `json:"title"`
	CreatedAt time.Time    `json:"createdAt"`
	DoneAt    sql.NullTime `json:"doneAt"`
}
