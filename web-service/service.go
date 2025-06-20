package main

import (
	"errors"
)

func GetLocation(id int, locations []location) location {
	for _, loc := range locations {
		if loc.ID == id {
			return loc
		}
	}

	return location{}
}

func UpdateLocation(id int, updatedLocation location, locations []location) (location, error) {
	for i, loc := range locations {
		if loc.ID == id {
			locations[i] = updatedLocation
			return updatedLocation, nil
		}
	}

	return location{}, errors.New("no location found")
}
