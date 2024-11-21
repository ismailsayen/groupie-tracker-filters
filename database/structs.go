package database

type Data struct {
	Art   []Artists
	MinDc int
	MaxDc int
	FrstAlbm int
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

func (d *Data) FindMinMax(a *[]Artists) {
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
