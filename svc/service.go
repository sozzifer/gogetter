package svc

import "gogetter/models"

type datasetStore interface {
	GetDatasets() models.Datasets
	GetDataset(id string) models.Dataset
}

type Service struct {
	Store datasetStore
}

func New(store datasetStore) *Service {
	svc := &Service{Store: store}
	return svc
}
