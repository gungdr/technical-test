package movie

import (
	"testing"
)

func TestSearchResult_ToProto(t *testing.T) {
	type fields struct {
		Search       []*MovieResult
		TotalResults string
		Result       *Result
	}
	errTest := "test"
	tests := []struct {
		name         string
		fields       fields
		wantErrEmpty bool
	}{
		{
			"Field error should empty",
			fields{
				Result: &Result{
					Error: nil,
				},
			},
			true,
		},
		{
			"Field error should not empty",
			fields{
				Result: &Result{
					Error: &errTest,
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := SearchResult{
				Search:       tt.fields.Search,
				TotalResults: tt.fields.TotalResults,
				Result:       tt.fields.Result,
			}
			if got := s.ToProto(); got.Error == "" && tt.wantErrEmpty == false {
				t.Errorf("SearchResult.ToProto() = %v, want %v", got.Error, tt.wantErrEmpty)
			}
		})
	}
}

func TestMovieDetailResult_ToProto(t *testing.T) {
	type fields struct {
		Title      string
		Year       string
		Rated      string
		Released   string
		Runtime    string
		Genre      string
		Director   string
		Writer     string
		Actors     string
		Plot       string
		Language   string
		Country    string
		Awards     string
		Poster     string
		Ratings    []Rating
		Metascore  string
		ImdbRating string
		ImdbVotes  string
		ImdbID     string
		Type       string
		DVD        string
		BoxOffice  string
		Production string
		Website    string
		Result     *Result
	}
	errTest := "test"
	tests := []struct {
		name         string
		fields       fields
		wantErrEmpty bool
	}{
		{
			"Field error should empty",
			fields{
				Result: &Result{
					Error: nil,
				},
			},
			true,
		},
		{
			"Field error should not empty",
			fields{
				Result: &Result{
					Error: &errTest,
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := MovieDetailResult{
				Title:      tt.fields.Title,
				Year:       tt.fields.Year,
				Rated:      tt.fields.Rated,
				Released:   tt.fields.Released,
				Runtime:    tt.fields.Runtime,
				Genre:      tt.fields.Genre,
				Director:   tt.fields.Director,
				Writer:     tt.fields.Writer,
				Actors:     tt.fields.Actors,
				Plot:       tt.fields.Plot,
				Language:   tt.fields.Language,
				Country:    tt.fields.Country,
				Awards:     tt.fields.Awards,
				Poster:     tt.fields.Poster,
				Ratings:    tt.fields.Ratings,
				Metascore:  tt.fields.Metascore,
				ImdbRating: tt.fields.ImdbRating,
				ImdbVotes:  tt.fields.ImdbVotes,
				ImdbID:     tt.fields.ImdbID,
				Type:       tt.fields.Type,
				DVD:        tt.fields.DVD,
				BoxOffice:  tt.fields.BoxOffice,
				Production: tt.fields.Production,
				Website:    tt.fields.Website,
				Result:     tt.fields.Result,
			}
			if got := m.ToProto(); got.Error == "" && tt.wantErrEmpty == false {
				t.Errorf("MovieDetailResult.ToProto() = %v, want %v", got.Error, tt.wantErrEmpty)
			}
		})
	}
}
