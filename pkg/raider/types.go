package raider

type Encounter struct {
	Id   int
	Slug string
	Name string
}

type Region struct {
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	ShortName string `json:"short_name"`
}

type Guild struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Faction string `json:"faction"`
	Realm   string `json:"realm"`
}

type Realm struct {
	Id               int    `json:"id"`
	ConnectedRealmId int    `json:"connectedRealmId"`
	WowRealmId       int    `json:"wowRealmId"`
	Name             string `json:"name"`
	AltName          string `json:"altName"`
	Slug             string `json:"slug"`
	AltSlug          string `json:"altSlug"`
	Locale           string `json:"locale"`
	IsConnected      bool   `json:"isConnected"`
	RealmType        string `json:"realmType"`
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

type RaidRanking struct {
	Rank               int                  `json:"rank"`
	RegionRank         int                  `json:"regionRank"`
	Guild              Guild                `json:"guild"`
	Region             Region               `json:"region"`
	Path               string               `json:"path"`
	Logo               string               `json:"logo"`
	Color              string               `json:"color"`
	EncountersDefeated []DefeatedEncounters `json:"encountersDefeated"`
	EncountersPulled   []PulledEncounters   `json:"encountersPulled"`
}

type DefeatedEncounters struct {
	Slug          string `json:"slug"`
	LastDefeated  string `json:"lastDefeated"`
	FirstDefeated string `json:"firstDefeated"`
}

type PulledEncounters struct {
	Id            int    `json:"id"`
	Slug          string `json:"slug"`
	NumPulls      int    `json:"numPulls"`
	PullStartedAt string `json:"pullStartedAt"`
	BestPercent   int    `json:"bestPercent"`
	IsDefeated    bool   `json:"isDefeated"`
}
