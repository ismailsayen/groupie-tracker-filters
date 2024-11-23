package database

import (
	"sync"
)

type Data struct {
	Art       []Artists
	MinDc     int
	MaxDc     int
	Locations map[string]bool
}

type Artists struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	Loca         Locations
	CongertDates string `json:"concertDates"`
	ConDT        Dates
	Relations    string `json:"relations"`
	Rela         Relation
}
type LocaFltr struct {
	Index []struct {
		ID        int      `json:"id"`
		Locations []string `json:"locations"`
		Dates     string   `json:"dates"`
	} `json:"index"`
}

type Locations struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

type Dates struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

type Relation struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type ErrorPage struct {
	Status int
	Type   string
}

func (d *Data) FindMinMax(a *[]Artists, wg *sync.WaitGroup) {
	defer wg.Done()
	min := 2024
	max := 0
	for _, ele := range *a {
		if min > ele.CreationDate {
			min = ele.CreationDate
		} else if max < ele.CreationDate {
			max = ele.CreationDate
		}
	}
	d.MaxDc = max
	d.MinDc = min
}

func (d *Data) AllLocations(l *LocaFltr, wg *sync.WaitGroup) {
	defer wg.Done()
	if d.Locations == nil {
		d.Locations = make(map[string]bool)
	}
	for _, ele := range l.Index {
		for _, e := range ele.Locations {
			d.Locations[e] = true
		}
	}
}
