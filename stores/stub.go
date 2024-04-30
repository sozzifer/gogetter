package stub

import "models"

type Stub struct {
	Path string
}

func (stub Stub) get_dataset() models.Dataset {
	dataset := models.Dataset{Id: "cpih", PublicationDate: "24/4/2024T00:00:00"}
	return dataset
}
