package models

type Datasets struct {
	Items []Dataset `json:"items"`
	Count int `json:"count"`
}

type Dataset struct {
	Id string `json:"id"`
	PublicationDate string `json:"publication_date"`
}