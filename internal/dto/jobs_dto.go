package dto

type CreateJobRequestBody struct {
	Position string  `json:"position"`
	Salary   float64 `json:"salary"`
	UserID   int64   `json:"userId"`
}
type UpdateJobRequestBody struct {
	Position string  `json:"position"`
	Salary   float64 `json:"salary"`
	UserID   int64   `json:"userId"`
}
type CreateJobService struct {
	Position string
	Salary   float64
	UserID   int64
}
type GetJobResponse struct {
	Position string
	Salary   float64
	UserID   int64
}
