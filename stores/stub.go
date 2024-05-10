package stores

import (
	"encoding/json"
	"fmt"
	"gogetter/models"
	"io"
	"os"
)

type Stub struct {
	Path string
}

func (s *Stub) GetDataset(id string) models.Dataset {
	datasets := s.GetDatasets()
	var dataset models.Dataset
	items := datasets.Items
	for _, item := range items {
		if item.Id == id {
			dataset = item
		}
	}
	return dataset
}

func (s *Stub) GetDatasets() models.Datasets {
	var datasets models.Datasets
	datasets_file, err := os.Open(s.Path)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer datasets_file.Close()
	datasets_bytes, err := io.ReadAll(datasets_file)
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(datasets_bytes, &datasets)
	if err != nil {
		fmt.Println(err)
	}
	return datasets
}
