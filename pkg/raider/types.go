package raider

type Encounter struct {
	Id   int
	Slug string
	Name string
}

type RaidData struct {
	Id         int         `json:"id"`
	Slug       string      `json:"slug"`
	Name       string      `json:"name"`
	ShortName  string      `json:"short_name"`
	Icon       string      `json:"icon"`
	StartTimes RegionTime  `json:"starts"`
	EndTimes   RegionTime  `json:"ends"`
	Encounters []Encounter `json:"encounters"`
}

type RegionTime struct {
	US string `json:"us"`
	EU string `json:"eu"`
	TW string `json:"tw"`
	KR string `json:"kr"`
	CN string `json:"cn"`
}
