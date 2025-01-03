package raider

type StaticDataOptions struct {
	Expansion int
}

type RaidRankingsOptions struct {
	Raid       string
	Difficulty string
	Region     string
	Limit      int
	Page       int
}

func NewRaidRankingsOptions() *RaidRankingsOptions {
	return &RaidRankingsOptions{
		Raid:       "",
		Difficulty: "",
		Region:     "",
		Limit:      10,
		Page:       0,
	}
}
