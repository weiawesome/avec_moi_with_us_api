package movie

type Movies struct {
	TotalPages  int     `json:"total_pages"`
	CurrentPage int     `json:"current_page"`
	MovieList   []Movie `json:"movies"`
}

type Movie struct {
	MovieId     string  `json:"movie_id"`
	Resource    string  `json:"resource"`
	ReleaseYear string  `json:"release_year"`
	Title       string  `json:"title"`
	Score       float32 `json:"score"`
}
