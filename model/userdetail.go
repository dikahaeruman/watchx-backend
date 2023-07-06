package model

type UserDetail struct {
	Watched_movies []int32 `json:"watched_movies" bson:"watched_movies"`
}
