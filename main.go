package main

import (
	"fmt"
	"gogetter/stores"
)

func main() {
	stub := stores.Stub{Path: "datasets.json"}
	datasets := stub.GetDatasets()
	fmt.Println(datasets.Items)
	fmt.Println(datasets.Count)
	dataset := stub.GetDataset("cpih")
	fmt.Println(dataset.Id)
	fmt.Println(dataset.PublicationDate)
}
