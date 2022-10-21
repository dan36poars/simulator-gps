package routes

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"strconv"
	"strings"
)

/* Structure Data */
type Route struct {
	ID        string     `json:"routeId"`
	ClientID  string     `json:"clientId"`
	Positions []Position `json:"position"`
}

type Position struct {
	Lat  float64 `json:"lat"`
	Long float64 `json:"long"`
}

// structured to send for another system
type PartialRoutePosition struct {
	ID       string    `json:"routeId"`
	ClientID string    `json:"clientID"`
	Position []float64 `json:"position"`
	Finished bool      `json:"finished"`
}

// Create a new Route
func NewRoute() *Route {
	return &Route{}
}

// load positions
func (r *Route) LoadPositions() error {
	if r.ID == "" {
		return errors.New("Route not informed.")
	}

	f, err := os.Open("destination/" + r.ID + ".txt")
	if err != nil {
		return err
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		data := strings.Split(scanner.Text(), ",")

		lat, err := strconv.ParseFloat(data[0], 64)
		if err != nil {
			return err
		}

		long, err := strconv.ParseFloat(data[1], 64)
		if err != nil {
			return err
		}

		r.Positions = append(r.Positions, Position{
			Lat:  lat,
			Long: long,
		})
	}

	return nil
}

// tranform positions string to json
func (r *Route) ExportJsonPositions() ([]string, error) {
	var route PartialRoutePosition
	var result []string

	total := len(r.Positions)

	for k, v := range r.Positions {
		route.ID = r.ID
		route.ClientID = r.ClientID
		route.Position = []float64{v.Lat, v.Long}
		route.Finished = false
		if total-1 == k {
			route.Finished = true
		}

		jsonRoute, err := json.Marshal(route) // cast to byte
		if err != nil {
			return nil, err
		}

		result = append(result, string(jsonRoute)) // cast to string
	}

	return result, nil
}
