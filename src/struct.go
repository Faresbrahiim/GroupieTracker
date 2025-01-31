package server

type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

type LocationsResponse struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
}

type Dates struct {
	ID    int      `json:"id"`
	DATES []string `json:"dates"`
}

type Relations struct {
	ID         int                 `json:"id"`
	DATESLOCAT map[string][]string `json:"datesLocations"`
}

type SecondPageData struct {
	ID        int
	Artist    Artist
	Locations []string
}

type ThirdPageData struct {
	ID             int                 `json:"id"`
	Dates          []string            `json:"dates"`
	DatesLocations map[string][]string `json:"datesLocations"`
}
