package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"groupietracker/database"
)

func FetchAPI(url string, s any) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("%d", resp.StatusCode)
	}

	if err := json.NewDecoder(resp.Body).Decode(&s); err != nil {
		return err
	}

	return nil
}
// get all artist details Locations,Relations and dates
func GetForeignData(artist *database.Artists) error {
	err := FetchAPI(artist.Locations, &artist.Loca)
	if err != nil {
		return err
	}

	err = FetchAPI(artist.CongertDates, &artist.ConDT)
	if err != nil {
		return err
	}

	err = FetchAPI(artist.Relations, &artist.Rela)
	if err != nil {
		return err
	}

	return nil
}
