package entity

// Translation -.
type User struct {
	Id         string   `json:"id"    example:"6be244a7-25ac-34ce-31e3-04157d3d42e3"`
	KeycloakId string   `json:"keycloak_id"    example:"6be244a7-25ac-34ce-31e3-04157d3d42e3"`
	ProfileIds []string `json:"profile_ids"    example:"[6be244a7-25ac-34ce-31e3-04157d3d42e3,6be244a7-25ac-34ce-31e3-04157d3d42e3]"`
}
