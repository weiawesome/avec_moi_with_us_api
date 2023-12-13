package movie

type SpecificMovie struct {
	MovieId       string      `json:"movie_id"`
	Resource      string      `json:"resource"`
	ReleaseYear   string      `json:"release_year"`
	Title         string      `json:"title"`
	OriginalTitle string      `json:"original_title"`
	Titles        []string    `json:"titles"`
	Categories    []string    `json:"categories"`
	Introduction  string      `json:"introduction"`
	RankScores    []RankScore `json:"rank_scores"`
	Actors        []Celebrity `json:"actors"`
	Directors     []Celebrity `json:"directors"`
}

type RankScore struct {
	Score        float32 `json:"score"`
	Organization string  `json:"organization"`
}
type Celebrity struct {
	Name     string `json:"name"`
	Resource string `json:"resource"`
	Gender   string `json:"gender"`
}
