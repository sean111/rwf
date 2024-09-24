package raider

type StaticDataResponse struct {
	Raid []RaidData `json:"raids"`
}

type RaidRankingsResponse struct {
	RaidRanking []RaidRanking `json:"raidRankings"`
}
