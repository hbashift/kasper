package models

type Specialization struct {
	SpecializationID int32  `json:"specialization_id,omitempty"`
	Name             string `json:"name,omitempty"`
}

type Group struct {
	GroupID int32  `json:"group_id,omitempty"`
	Name    string `json:"name,omitempty"`
}
