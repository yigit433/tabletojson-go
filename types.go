package tabletojson

type Column struct {
	ParentName  string `json:"parentName,omitempty"`
	ParentValue string `json:"parentValue,omitempty"`
}

type Row []Column

type Table []Row
