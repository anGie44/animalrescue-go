package animresc

import (
	"fmt"
	"net/url"
	"net/http"
	"encoding/json"
	"io"
	"bytes"
)

type Client struct {
	BaseURL *url.URL
	UserAgent string

	httpClient *http.Client 
}

type Adoption struct {
	ID int `json:"id"`
	Adopter *Adopter `json:"adopter"`
	Adoptee *Adoptee `json:"adoptee"`
	Date string `json:"date"`
}

type Adopter struct {
	ID             int             `json:"id"`
	FirstName      string          `json:"first_name"`
	LastName       string          `json:"last_name"`
	Phone          string          `json:"phone"`
	Email          string          `json:"email"`
	Gender         string          `json:"gender"`
	Birthdate      string          `json:"birthdate"`
	Address        string          `json:"address"`
	Country        string          `json:"country"`
	State          string          `json:"state"`
	City           string          `json:"city"`
	ZipCode        string          `json:"zip_code"`
	PetPreferences []PetPreference `json:"pet_preferences"`
}

type Adoptee struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Breed  string `json:"breed"`
	Gender string `json:"gender"`
	Age    string `json:"age"`
}

type PetPreference struct {
	ID     int    `json:"id"`
	Breed  string `json:"breed"`
	Age    string `json:"age"`
	Gender string `json:"gender"`
}

func NewClient(opts ...func(*Client)) *Client {
	c := &Client{
		BaseURL: &url.URL{
			Scheme: "http",
			Host: "localhost:3000",
		},
		UserAgent: "animresc-go/0.1.0",
		httpClient: &http.Client{},
	}

	for _, f := range opts {
		f(c)
	}

	return c
}

// WithBaseURL sets the Client base URL.
func WithBaseURL(rawurl string) func(*Client) {
	return func(c *Client) {
		if u, err := url.Parse(rawurl); err == nil {
			c.BaseURL = u
		}
	}
}

func (c *Client) Get(a interface{}, name, id string) (interface{}, error) {
	req, err := c.newRequest("GET", fmt.Sprintf("/%s/%s", name, id), nil)
	if err != nil {
		return a, err
	}
	_, err = c.do(req, &a)
	return a, err
}

func (c *Client) GetAll(a interface{}, name string) (interface{}, error) {
	req, err := c.newRequest("GET", fmt.Sprintf("/%s", name), nil)
	if err != nil {
		return a, err
	}
	_, err = c.do(req, &a)
	return a, err
}

func (c *Client) Create(a interface{}, name string) (interface{}, error) {
	req, err := c.newRequest("POST", fmt.Sprintf("/%s", name), nil)
	if err != nil {
		return a, err
	}
	_, err = c.do(req, &a)
	return a, err
}

func (c *Client) Update(a interface{}, name, id string) (interface{}, error) {
	req, err := c.newRequest("GET", fmt.Sprintf("/%s/%s/update", name, id), nil)
	if err != nil {
		return a, err
	}
	_, err = c.do(req, &a)
	return a, err
}

func (c *Client) Delete(name, id string) error {
	req, err := c.newRequest("GET", fmt.Sprintf("/%s/%s/delete", name, id), nil)
	if err != nil {
		return err
	}
	
	_, err = c.do(req, nil)
	return err
}

// Helper Methods

func (c *Client) newRequest(method, path string, body interface{}) (*http.Request, error) {
	rel := &url.URL{Path: path}
	u := c.BaseURL.ResolveReference(rel)

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)

	return req, nil
}

func (c *Client) do(r *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.httpClient.Do(r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(v)
	return resp, err
}
