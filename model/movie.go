package model

type Movie struct {
	ID                int     `json:"id" bson:"id"`
	Adult             bool    `json:"adult" bson:"adult"`
	Backdrop_path     string  `json:"backdrop_path" bson:"backdrop_path"`
	Genre_ids         []int16 `json:"genre_ids" bson:"genre_ids"`
	Original_language string  `json:"original_language" bson:"original_language"`
	Original_title    string  `json:"original_title" bson:"original_title"`
	Overview          string  `json:"overview" bson:"overview"`
	Release_date      string  `json:"release_date" bson:"release_date"`
	Poster_path       string  `json:"poster_path" bson:"poster_path"`
	Vote_average      float32 `json:"vote_average" bson:"vote_average"`
	Vote_count        int     `json:"vote_count" bson:"vote_count"`
}
