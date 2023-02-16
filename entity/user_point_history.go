package entity

import "time"

type UserPointHistoryID int64
type PointOperation int8
type PointOperationDate time.Time
type PointExpiryDate time.Time
type PointAmount int32
type PointRemaining int32

type UserPointHistory struct {
	ID            UserPointHistoryID `json:"id"`
	UserPointID   UserPointID        `json:"user_point_id"`
	Operation     PointOperation     `json:"operation"`
	OperationDate PointOperationDate `json:"operation_date"`
	ExpiryDate    PointExpiryDate    `json:"expiry_date"`
	Amount        PointAmount        `json:"amount"`
	Remaining     PointRemaining     `json:"remaining"`
	Created       time.Time          `json:"created"`
	Updated       time.Time          `json:"updated"`
}

type UserPointHistories []*UserPointHistory
