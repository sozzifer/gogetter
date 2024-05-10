package svc

import "gogetter/models"

type datasetStore interface {
	GetDatasets() models.Dataset
	GetDataset(id string) models.Dataset
}
