package animalrescue

import (
	"context"
	"fmt"
)

// AdoptersService provides access to the adopter-related functions
// in the Animal Rescue API
type AdoptersService service

// Adopter represents an adopter within an Animal Rescue organization.
type Adopter struct {
	ID             *int64           `json:"id,omitempty"`
	FirstName      *string          `json:"first_name,omitempty"`
	LastName       *string          `json:"last_name,omitempty"`
	Phone          *string          `json:"phone,omitempty"`
	Email          *string          `json:"email,omitempty"`
	Gender         *string          `json:"gender,omitempty"`
	Birthdate      *string          `json:"birthdate,omitempty"`
	Address        *string          `json:"address,omitempty"`
	Country        *string          `json:"country,omitempty"`
	State          *string          `json:"state,omitempty"`
	City           *string          `json:"city,omitempty"`
	ZipCode        *string          `json:"zip_code,omitempty"`
	PetPreferences []*PetPreference `json:"pet_preferences,omitempty"`
}

func (a Adopter) String() string {
	return Stringify(a)
}

// ListAdopters lists all of the adopters for an animal rescue.
func (s *AdoptersService) ListAdopters(ctx context.Context) ([]*Adopter, *Response, error) {
	u := "adopters"
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}
	var adopters []*Adopter
	resp, err := s.client.Do(ctx, req, &adopters)
	if err != nil {
		return nil, resp, err
	}

	return adopters, resp, nil
}

// GetAdopterByID fetches an adopter by ID.
func (s *AdoptersService) GetAdopterByID(ctx context.Context, adopterID int64) (*Adopter, *Response, error) {
	u := fmt.Sprintf("adopter/%v", adopterID)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	a := new(Adopter)
	resp, err := s.client.Do(ctx, req, a)
	if err != nil {
		return nil, resp, err
	}

	return a, resp, nil
}

// NewAdopter represents an adopter to be created or modified.
type NewAdopter struct {
	FirstName      *string          `json:"first_name,omitempty"`
	LastName       *string          `json:"last_name,omitempty"`
	Phone          *string          `json:"phone,omitempty"`
	Email          *string          `json:"email,omitempty"`
	Gender         *string          `json:"gender,omitempty"`
	Birthdate      *string          `json:"birthdate,omitempty"`
	Address        *string          `json:"address,omitempty"`
	Country        *string          `json:"country,omitempty"`
	State          *string          `json:"state,omitempty"`
	City           *string          `json:"city,omitempty"`
	ZipCode        *string          `json:"zip_code,omitempty"`
	PetPreferences []*PetPreference `json:"pet_preferences,omitempty"`
}

// CreateAdopter creates a new adopter within an animal rescue.
func (s *AdoptersService) CreateAdopter(ctx context.Context, adopter NewAdopter) (*Adopter, *Response, error) {
	u := "adopters"
	req, err := s.client.NewRequest("POST", u, adopter)
	if err != nil {
		return nil, nil, err
	}
	a := new(Adopter)
	resp, err := s.client.Do(ctx, req, a)
	if err != nil {
		return nil, resp, err
	}

	return a, resp, nil
}

// EditAdopterByID edits an adopter by ID.
func (s *AdoptersService) EditAdopterByID(ctx context.Context, adopterID int64, adopter NewAdopter) (*Adopter, *Response, error) {
	u := fmt.Sprintf("adopter/%v", adopterID)
	req, err := s.client.NewRequest("PATCH", u, adopter)
	if err != nil {
		return nil, nil, err
	}

	a := new(Adopter)
	resp, err := s.client.Do(ctx, req, a)
	if err != nil {
		return nil, resp, err
	}

	return a, resp, nil

}

// DeleteAdopterByID deletes an adopter referenced by ID
func (s *AdoptersService) DeleteAdopterByID(ctx context.Context, adopterID int64) (*Response, error) {
	u := fmt.Sprintf("/adopter/%v", adopterID)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}
