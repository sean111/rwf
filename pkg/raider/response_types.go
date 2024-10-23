package raider

type StaticDataResponse struct {
	Status int
	Data   []RaidData `json:"raids"`
}

type RaidRankingsResponse struct {
	Status int
	Data   []RaidRanking `json:"raidRankings"`
}
