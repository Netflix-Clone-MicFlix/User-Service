package entity

// Translation -.
type Profile struct {
	Id          string   `json:"id"    example:"6be244a7-25ac-34ce-31e3-04157d3d42e3"`
	UserId      string   `json:"user_id"    example:"6be244a7-25ac-34ce-31e3-04157d3d42e3"`
	Name        string   `json:"name"    example:"Donald goose"`
	MovieTagIds []string `json:"movie_tag_Ids"    example:"6be244a7-25ac-34ce-31e3-04157d3d42e3"`
}
