package animalrescue

import (
	"context"
	"fmt"
)

// AdopteesService provides access to the adoptee-related functions
// in the Animal Rescue API
type AdopteesService service

// Adoptee represents an adoptee within an Animal Rescue organization.
type Adoptee struct {
	ID     int    `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Breed  string `json:"breed,omitempty"`
	Gender string `json:"gender,omitempty"`
	Age    string `json:"age,omitempty"`
}

func (a Adoptee) String() string {
	return Stringify(a)
}

// ListAll lists all of the adoptees for an animal rescue.
func (s *AdopteesService) ListAll(ctx context.Context) ([]*Adoptee, *Response, error) {
	u := "adoptees"
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}
	var adoptees []*Adoptee
	resp, err := s.client.Do(ctx, req, &adoptees)
	if err != nil {
		return nil, resp, err
	}

	return adoptees, resp, nil
}

// GetAdopteeByID fetches an adoptee by ID.
func (s *AdopteesService) GetAdopteeByID(ctx context.Context, adopteeID int64) (*Adoptee, *Response, error) {
	u := fmt.Sprintf("adoptee/%v", adopteeID)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	a := new(Adoptee)
	resp, err := s.client.Do(ctx, req, a)
	if err != nil {
		return nil, resp, err
	}

	return a, resp, nil
}

// NewAdoptee represents an adoptee to be created or modified.
type NewAdoptee struct {
	Name   string `json:"name,omitempty"`
	Breed  string `json:"breed,omitempty"`
	Gender string `json:"gender,omitempty"`
	Age    string `json:"age,omitempty"`
}

// CreateAdoptee creates a new adoptee within an animal rescue.
func (s *AdopteesService) CreateAdoptee(ctx context.Context, adoptee NewAdoptee) (*Adoptee, *Response, error) {
	u := "adoptees"
	req, err := s.client.NewRequest("POST", u, adoptee)
	if err != nil {
		return nil, nil, err
	}
	a := new(Adoptee)
	resp, err := s.client.Do(ctx, req, a)
	if err != nil {
		return nil, resp, err
	}

	return a, resp, nil
}

// EditAdopteeByID edits an adoptee selected by ID.
func (s *AdopteesService) EditAdopteeByID(ctx context.Context, adopteeID int64, adoptee NewAdoptee) (*Adoptee, *Response, error) {
	u := fmt.Sprintf("adoptee/%v", adopteeID)
	req, err := s.client.NewRequest("PATCH", u, adoptee)
	if err != nil {
		return nil, nil, err
	}

	a := new(Adoptee)
	resp, err := s.client.Do(ctx, req, a)
	if err != nil {
		return nil, resp, err
	}

	return a, resp, nil

}

// DeleteAdopteeByID deletes an adoptee referenced by ID.
func (s *AdopteesService) DeleteAdopteeByID(ctx context.Context, adopteeID int64) (*Response, error) {
	u := fmt.Sprintf("/adoptee/%v", adopteeID)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}
