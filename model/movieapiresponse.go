package model

type MovieApiResponse struct {
	Page          int     `json:"page,omitempty" bson:"page,omitempty"`
	Results       []Movie `json:"results,omitempty" bson:"results,omitempty"`
	Total_pages   int     `json:"total_pages,omitempty" bson:"total_pages,omitempty"`
	Total_results int     `json:"total_results,omitempty" bson:"total_results,omitempty"`
}
