package raider

import "testing"

func TestGetStaticData(t *testing.T) {
	tests := []struct {
		options  StaticDataOptions
		expected string
	}{
		{
			options: StaticDataOptions{
				Expansion: Expansion.DragonFlight,
			},
			expected: "aberrus-the-shadowed-crucible",
		},
		{
			options: StaticDataOptions{
				Expansion: Expansion.Legion,
			},
			expected: "antorus-the-burning-throne",
		},
	}
	for _, test := range tests {
		res, err := GetStaticData(test.options)
		if err != nil {
			t.Fatalf("GetStaticData(10) failed: %v", err)
		}
		if res.Status != 200 {
			t.Fatalf("GetStaticData(10) status = %q, want 200", res.Status)
		}
		if res.Data[0].Slug != test.expected {
			t.Fatalf(`GetStaticData(10) = %q, want "%q"`, res.Data[0].Slug, test.expected)
		}
	}
}

func TestGetRaidRankings(t *testing.T) {
	tests := []struct {
		options  RaidRankingsOptions
		expected string
	}{
		{
			options: RaidRankingsOptions{
				Raid:       "nerubar-palace",
				Difficulty: Difficulty.Mythic,
				Region:     "world",
				Limit:      1,
				Page:       0,
			},
			expected: "Liquid",
		},
		{
			options: RaidRankingsOptions{
				Raid:       "antorus-the-burning-throne",
				Difficulty: Difficulty.Mythic,
				Region:     "world",
				Limit:      1,
				Page:       2,
			},
			expected: "Memento",
		},
	}

	for _, test := range tests {
		res, err := GetRaidRankings(test.options)
		if err != nil {
			t.Fatalf(`GetRaidRankings("nerubar-palace", "mythic", "world", 1, 0) failed: %v`, err)
		}

		if res.Status != 200 {
			t.Fatalf(`GetRaidRankings("nerubar-palace", "mythic", "world", 1, 0) Status = %q, want 200`, res.Status)
		}
		if res.Data[0].Guild.Name != test.expected {
			t.Fatalf(`GetRaidRankings(%q) Guild Name = %q, want "%q"`, test.options, res.Data[0].Guild.Name, test.expected)
		}
	}

}
