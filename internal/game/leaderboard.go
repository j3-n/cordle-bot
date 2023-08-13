package game

import (
	"cordle/internal/pkg/util"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func GetLeaderboard(s *discordgo.Session) string {
	t, err := db.ReadTop()
	util.PrintErr(err)

	o := "``\n"
	for index, user := range t {
		i, err := s.User(user.Id)
		util.PrintErr(err)

		o += fmt.Sprintf(
			"%d : %s (%s)\n",
			index+1,
			i,
			user.ToLeaderboard(),
		)
	}
	return o + "``"
}
