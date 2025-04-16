package models

type UserRank struct {
	Username  string `json:"username"`
	PostCount uint   `json:"postCount"`
	Rank      uint   `json:"rank"`
}
