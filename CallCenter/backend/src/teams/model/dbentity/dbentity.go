package dbentity

type Team struct {
	Id          string `json:"id,omitempty" bson:"id,omitempty"`
	DisplayName string `json:"displayName,omitempty" bson:"displayName,omitempty"`
	Area        string `json:"area,omitempty" bson:"area,omitempty"`
	TeamLead    string `json:"teamLead,omitempty" bson:"teamLead,omitempty"`
}
