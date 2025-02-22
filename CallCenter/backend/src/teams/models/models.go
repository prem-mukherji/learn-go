package models

type TeamRequest struct {
	Id          string `json:"id,omitempty" bson:"id,omitempty"`
	DisplayName string `json:"displayName,omitempty" bson:"displayName,omitempty"`
	Area        string `json:"area,omitempty" bson:"area,omitempty"`
	TeamLead    string `json:"teamLead,omitempty" bson:"teamLead,omitempty"`
}

type TeamViewModel struct {
	Id          string `json:"id,omitempty" bson:"id,omitempty"`
	DisplayName string `json:"displayName,omitempty" bson:"displayName,omitempty"`
	Area        string `json:"area,omitempty" bson:"area,omitempty"`
	TeamLead    string `json:"teamLead,omitempty" bson:"teamLead,omitempty"`
}

// Define a complex data structure
type ApiResponse struct {
	Status  string        `json:"status"`
	Message string        `json:"message"`
	Data    []interface{} `json:"data"`
}

// Another struct for a specific type of data
type UserData struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}
