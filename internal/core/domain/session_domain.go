package domain

type GenerateSessionRespose struct {
	SessionToken string `json:"session_token"`
}

type ValidateSession struct {
	UUID string
}
