package request_models

type FirstStudentRegistry struct {
	FullName         string `json:"full_name,omitempty"`
	Department       string `json:"department,omitempty"`
	SpecializationID int32  `json:"specialization_id,omitempty"`
	ActualSemester   int32  `json:"actual_semester,omitempty"`
	NumberOfYears    int32  `json:"number_of_years,omitempty"`
	StartDate        string `json:"start_date,omitempty"`
	GroupID          int32  `json:"group_number,omitempty"`
}
