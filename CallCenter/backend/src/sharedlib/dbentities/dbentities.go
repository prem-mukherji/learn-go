package dbentities

type Team struct {
	ID          string `json:"id,omitempty" bson:"id,omitempty"`
	DisplayName string `json:"displayName,omitempty" bson:"displayName,omitempty"`
	Area        string `json:"area,omitempty" bson:"area,omitempty"`
	TeamLead    string `json:"teamLead,omitempty" bson:"teamLead,omitempty"`
	UpdatedBy   string `json:"updatedBy,omitempty" bson:"updatedBy,omitempty"`
	UpdatedDate string `json:"updatedDate,omitempty" bson:"updatedDate,omitempty"`
}
