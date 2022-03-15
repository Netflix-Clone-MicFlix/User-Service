// Package entity defines main entities for business logic (services), data base mapping and
// HTTP response objects if suitable. Each logic group entities in own file.
package entity

import "time"

// Translation -.
type Salt struct {
	Id        string    `json:"_id"          example:"22871fbc8e3ef8"`
	UserId    string    `json:"user_id"      example:"52836fhc8e4ef8"`
	SaltData  []byte    `json:"salt_data"    example:"[7 20 118 194 45 247 47 37 106 90 251 54 144 36 7 124]"`
	CreatedAt time.Time `json:"created_at"   example:"2022-02-17 13:39:03.809450"`
}
