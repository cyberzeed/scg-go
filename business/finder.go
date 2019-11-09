package business

import (
	"context"
	"log"

	"googlemaps.github.io/maps"
)

// Finder struct
type Finder struct {
	client *maps.Client
}

// SearchOption type
type SearchOption maps.TextSearchRequest

// NewFinder is factory function for create finder object
func NewFinder(apiKey string) *Finder {
	client, err := maps.NewClient(maps.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("Cannot connect Google Place API, %v\n", err)
	}
	return &Finder{client}
}

// Search business address and information in area
func (pf *Finder) Search(option SearchOption) ([]maps.PlacesSearchResult, error) {
	request := maps.TextSearchRequest(option)
	response, err := pf.client.TextSearch(context.Background(), &request)
	if err != nil {
		return nil, err
	}

	return response.Results, nil
}

// With function will create search object
func With(area string, businessType string) SearchOption {
	return SearchOption{Query: area, Type: maps.PlaceType(businessType), Region: "TH"}
}

// WithPageToken function will create search object
func WithPageToken(pageToken string) SearchOption {
	return SearchOption{PageToken: pageToken}
}
