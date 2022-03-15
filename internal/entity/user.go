// Package entity defines main entities for business logic (services), data base mapping and
// HTTP response objects if suitable. Each logic group entities in own file.
package entity

// Translation -.
type User struct {
	Id       string `json:"_id"            example:"1"`
	Username string `json:"username"      example:"poeky"`
	Password string `json:"password"      example:"dsafsd111fa"`
	Email    string `json:"email"         example:"example@hotmail.com"`
}
