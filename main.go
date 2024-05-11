package main

import (
	"fmt"
	"gogetter/stores"
	"gogetter/svc"
)

func main() {
	stub := stores.Stub{Path: "datasets.json"}
	svc := svc.New(&stub)
	fmt.Println(svc.Store.GetDatasets())
	fmt.Println(svc.Store.GetDataset("cpih"))
}
