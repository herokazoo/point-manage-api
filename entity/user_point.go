package entity

import "time"

type UserID int64
type UserPointID int64
type UserPointTotal int32

type UserPoint struct {
	ID      UserPointID    `json:"id"`
	UserId  UserID         `json:"user_id"`
	Total   UserPointTotal `json:"total"`
	Created time.Time      `json:"created"`
	Updated time.Time      `json:"updated"`
}
