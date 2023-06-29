package interfaces

type InputCertificate struct {
	ID     string `json:"candidateID" binding:"required"`
	Name   string `json:"candidateName" binding:"required"`
	Course string `json:"courseName" binding:"required"`
	Grade  string `json:"courseGrade" binding:"required"`
	Date   string `json:"courseDate" binding:"required"`
}

type ReturnCertificate struct {
	Name   string
	Course string
	Grade  string
	Date   string
}
