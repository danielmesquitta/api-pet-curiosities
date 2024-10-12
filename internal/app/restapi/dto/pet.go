package dto

import (
	"github.com/danielmesquitta/api-pet-curiosities/internal/provider/db/ent/pet"
	"github.com/google/uuid"
)

type Pet struct {
	ID     uuid.UUID  `json:"id,omitempty"`
	Breed  string     `json:"breed,omitempty"`
	Specie pet.Specie `json:"specie,omitempty"`
}

type ListPetsRequestDTO struct {
	Specie pet.Specie `json:"specie,omitempty"`
	Breed  string     `json:"breed,omitempty"`
}

type ListPetsResponseDTO struct {
	Data []Pet `json:"data"`
}
