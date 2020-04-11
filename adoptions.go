package animalrescue

import (
	"context"
	"fmt"
	"time"
)

// AdoptionsService provides access to adoption-related funcions
// in the Animal Rescue API
type AdoptionsService service

// Adoption represents an adoption event within an animal rescue. Adoptions
// are used to store adopter and adoptee relationships.
type Adoption struct {
	ID        int      `json:"id,omitempty"`
	Adopter   *Adopter `json:"adopter,omitempty"`
	Adoptee   *Adoptee `json:"adoptee,omitempty"`
	CreatedAt string   `json:"created_at,omitempty"`
}

func (a Adoption) String() string {
	return Stringify(a)
}

// ListAll lists all of the adoptions for an animal rescue.
func (s *AdoptionsService) ListAll(ctx context.Context) ([]*Adoption, *Response, error) {
	u := "adoptions"
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}
	var adoptions []*Adoption
	resp, err := s.client.Do(ctx, req, &adoptions)
	if err != nil {
		return nil, resp, err
	}

	return adoptions, resp, nil
}

// GetAdoptionByID fetches an adoption by ID.
func (s *AdoptionsService) GetAdoptionByID(ctx context.Context, adoptionID int64) (*Adoption, *Response, error) {
	u := fmt.Sprintf("adoption/%v", adoptionID)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	a := new(Adoption)
	resp, err := s.client.Do(ctx, req, a)
	if err != nil {
		return nil, resp, err
	}

	return a, resp, nil
}

// NewAdoption represents a team to be created.
type NewAdoption struct {
	Adopter   *Adopter   `json:"adopter,omitempty"`
	Adoptee   *Adoptee   `json:"adoptee,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
}

// CreateAdoption creates a new adoption within an animal rescue.
func (s *AdoptionsService) CreateAdoption(ctx context.Context, adoption NewAdoption) (*Adoption, *Response, error) {
	u := "adoptions"
	req, err := s.client.NewRequest("POST", u, adoption)
	if err != nil {
		return nil, nil, err
	}
	a := new(Adoption)
	resp, err := s.client.Do(ctx, req, a)
	if err != nil {
		return nil, resp, err
	}

	return a, resp, nil
}

// DeleteAdoptionByID delets an adoption referenced by ID.
func (s *AdoptionsService) DeleteAdoptionByID(ctx context.Context, adoptionID int64) (*Response, error) {
	u := fmt.Sprintf("/adoption/%v", adoptionID)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}
