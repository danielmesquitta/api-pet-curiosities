package dto

type LoginRequestDTO struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

type LoginResponseDTO struct {
	AccessToken  string `json:"access_token,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
}
