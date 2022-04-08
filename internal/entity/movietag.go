package entity

// Translation -.
type MovieTag struct {
	Id         string `json:"id"             example:"6be244a7-25ac-34ce-31e3-04157d3d42e3"`
	GenreId    string `json:"genre_id"       example:"6be244a7-25ac-34ce-31e3-04157d3d42e3"`
	WatchCount int    `json:"watch_count"    example:"43"`
}
