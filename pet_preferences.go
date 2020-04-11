package animalrescue

import (
	"context"
	"fmt"
)

// PetPreferencesService provides access to the pet_preference-related functions
// in the Animal Rescue API.
type PetPreferencesService service

// PetPreference represents a preference made by a prospective adopter in
// an animal rescue.
type PetPreference struct {
	ID     int    `json:"id,omitempty"`
	Breed  string `json:"breed,omitempty"`
	Age    string `json:"age,omitempty"`
	Gender string `json:"gender,omitempty"`
}

func (pp PetPreference) String() string {
	return Stringify(pp)
}

// ListPetPreferences lists all of the pet-preferences within an animal rescue.
func (s *PetPreferencesService) ListPetPreferences(ctx context.Context) ([]*PetPreference, *Response, error) {
	u := "petprefs"
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}
	var pp []*PetPreference
	resp, err := s.client.Do(ctx, req, &pp)
	if err != nil {
		return nil, resp, err
	}
	return pp, resp, nil
}

// GetPetPreferenceByID fetches a pet-preference by ID.
func (s *PetPreferencesService) GetPetPreferenceByID(ctx context.Context, ppID int64) (*PetPreference, *Response, error) {
	u := fmt.Sprintf("petpref/%v", ppID)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}
	pp := new(PetPreference)
	resp, err := s.client.Do(ctx, req, pp)
	if err != nil {
		return nil, resp, err
	}

	return pp, resp, nil
}

// NewPetPreference represents a pet-preference to be created or modified.
type NewPetPreference struct {
	Breed  string `json:"breed,omitempty"`
	Age    string `json:"age,omitempty"`
	Gender string `json:"gender,omitempty"`
}

// CreatePetPreference creates a new pet-preference within an animal rescue.
func (s *PetPreferencesService) CreatePetPreference(ctx context.Context, pp NewPetPreference) (*PetPreference, *Response, error) {
	u := "petprefs"
	req, err := s.client.NewRequest("POST", u, pp)
	if err != nil {
		return nil, nil, err
	}
	p := new(PetPreference)
	resp, err := s.client.Do(ctx, req, p)
	if err != nil {
		return nil, resp, err
	}

	return p, resp, nil

}

// EditPetPreferenceByID edits a pet-preference selected by ID.
func (s *PetPreferencesService) EditPetPreferenceByID(ctx context.Context, ppID int64, pp NewPetPreference) (*PetPreference, *Response, error) {
	u := fmt.Sprintf("petpref/%v", ppID)
	req, err := s.client.NewRequest("PATCH", u, pp)
	if err != nil {
		return nil, nil, err
	}
	p := new(PetPreference)
	resp, err := s.client.Do(ctx, req, p)
	if err != nil {
		return nil, resp, err
	}

	return p, resp, nil
}

// DeletePetPreferenceByID deletes a pet-preference referenced by ID.
func (s *PetPreferencesService) DeletePetPreferenceByID(ctx context.Context, ppID int64) (*Response, error) {
	u := fmt.Sprintf("petpref/%v", ppID)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}
	return s.client.Do(ctx, req, nil)
}
