package movie

import "omdb/proto"

type Result struct {
	Response string  `json:"Response"`
	Error    *string `json:"Error"`
}
type SearchResult struct {
	Search       []*MovieResult `json:"Search"`
	TotalResults string         `json:"totalResults"`
	*Result
}

func (s SearchResult) ToProto() *proto.SearchResult {
	movieResultProto := []*proto.MovieResult{}
	for _, item := range s.Search {
		movieResultProto = append(movieResultProto, item.ToProto())
	}
	var err string
	if s.Error != nil {
		err = *s.Error
	}
	return &proto.SearchResult{
		Search:       movieResultProto,
		TotalResults: s.TotalResults,
		Response:     s.Response,
		Error:        err,
	}
}

type IMDBID string

func (i IMDBID) ToMap() map[string]string {
	return map[string]string{
		"i": string(i),
	}
}

type MovieResult struct {
	Title  string `json:"title"`
	Year   string `json:"year"`
	IMDBID IMDBID `json:"imdbID"`
	Type   string `json:"type"`
	Poster string `json:"poster"`
}

func (m MovieResult) ToProto() *proto.MovieResult {
	return &proto.MovieResult{
		Title:  m.Title,
		Year:   m.Year,
		ImdbId: string(m.IMDBID),
		Type:   m.Type,
		Poster: m.Poster,
	}
}

type Rating struct {
	Source string `json:"Source"`
	Value  string `json:"Value"`
}

func (r Rating) ToProto() *proto.Rating {
	return &proto.Rating{
		Source: r.Source,
		Value:  r.Value,
	}
}

type MovieDetailResult struct {
	Title      string   `json:"Title"`
	Year       string   `json:"Year"`
	Rated      string   `json:"Rated"`
	Released   string   `json:"Released"`
	Runtime    string   `json:"Runtime"`
	Genre      string   `json:"Genre"`
	Director   string   `json:"Director"`
	Writer     string   `json:"Writer"`
	Actors     string   `json:"Actors"`
	Plot       string   `json:"Plot"`
	Language   string   `json:"Language"`
	Country    string   `json:"Country"`
	Awards     string   `json:"Awards"`
	Poster     string   `json:"Poster"`
	Ratings    []Rating `json:"Ratings"`
	Metascore  string   `json:"Metascore"`
	ImdbRating string   `json:"imdbRating"`
	ImdbVotes  string   `json:"imdbVotes"`
	ImdbID     string   `json:"imdbID"`
	Type       string   `json:"Type"`
	DVD        string   `json:"DVD"`
	BoxOffice  string   `json:"BoxOffice"`
	Production string   `json:"Production"`
	Website    string   `json:"Website"`
	*Result
}

func (m MovieDetailResult) ToProto() *proto.MovieDetailResult {
	ratingProto := []*proto.Rating{}
	for _, rating := range m.Ratings {
		ratingProto = append(ratingProto, rating.ToProto())
	}
	var err string
	if m.Error != nil {
		err = *m.Error
	}
	return &proto.MovieDetailResult{
		Title:      m.Title,
		Year:       m.Year,
		Rated:      m.Rated,
		Released:   m.Released,
		Runtime:    m.Runtime,
		Genre:      m.Genre,
		Director:   m.Director,
		Writer:     m.Writer,
		Actors:     m.Actors,
		Plot:       m.Plot,
		Language:   m.Language,
		Country:    m.Country,
		Awards:     m.Awards,
		Poster:     m.Poster,
		Ratings:    ratingProto,
		Metascore:  m.Metascore,
		ImdbRating: m.ImdbRating,
		ImdbVotes:  m.ImdbVotes,
		ImdbId:     m.ImdbID,
		Type:       m.Type,
		Dvd:        m.DVD,
		BoxOffice:  m.BoxOffice,
		Production: m.Production,
		Website:    m.Website,
		Response:   m.Response,
		Error:      err,
	}
}
