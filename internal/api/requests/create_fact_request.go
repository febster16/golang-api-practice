package requests

type CreateFactPayload struct {
	Question string `json:"question" example:"Question x?"`
	Answer   string `json:"answer" example:"Answer x."`
}
