package types

import "github.com/google/uuid"

type User struct {
	name string
}

// 被控端设备信息
type Device struct {
	IP string    `json:"ip"`
	ID uuid.UUID `json:"id"`
}
