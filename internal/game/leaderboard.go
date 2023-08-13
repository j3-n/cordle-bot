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
		u, err := s.User(user.Id)
		util.PrintErr(err)
		i := u.Username

		o += fmt.Sprintf(
			"%d : %s (%d)\n",
			index+1,
			i,
			user.Elo,
		)
	}
	return o + "``"
}
