package scraper

type Comic struct {
	Name        string             `json:"name"`
	Rate        string             `json:"rate"`
	RatingCount string             `json:"rating_count"`
	Views       string             `json:"views"`
	URL         string             `json:"url"`
	AltName     string             `json:"alt_name"`
	Status      string             `json:"status"`
	Summary     string             `json:"summary"`
	Publisher   map[string]string  `json:"publisher"`
	Authors     map[string]string  `json:"authors"`
	Artists     map[string]string  `json:"artists"`
	Genres      map[string]string  `json:"genres"`
	Tags        map[string]string  `json:"tags"`
	Chapters    map[string]Chapter `json:"chapters"`
}

type Chapter struct {
	Name          string `json:"name"`
	URL           string `json:"url"`
	PublishedDate string `json:"published_date"`
}
