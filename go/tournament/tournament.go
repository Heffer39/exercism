//Package tournament tallies the results of a small football competition
package tournament

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
)

// team captures the name and stats of a team for the competition
type team struct {
	name                                       string
	matchesPlayed, wins, draws, losses, points int
}

var teams map[string]team

// Tally reads in an input file, calculates the match results, sorts the results, and writes back
// the results to the io.Writer
func Tally(r io.Reader, w io.Writer) error {
	teams = make(map[string]team)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), ";")
		if s[0] == "" || strings.HasPrefix(s[0], "#") {
			continue
		}
		if len(s) != 3 {
			return fmt.Errorf("bad line format %q: expected 'teamA;teamB;result'", s)
		}
		a, b := teams[s[0]], teams[s[1]]
		a.name, b.name = s[0], s[1]
		switch s[2] {
		case "draw":
			a.matchesPlayed, a.draws, a.points = a.matchesPlayed+1, a.draws+1, a.points+1
			b.matchesPlayed, b.draws, b.points = b.matchesPlayed+1, b.draws+1, b.points+1
		case "win":
			a.matchesPlayed, a.wins, a.points = a.matchesPlayed+1, a.wins+1, a.points+3
			b.matchesPlayed, b.losses = b.matchesPlayed+1, b.losses+1
		case "loss":
			a.matchesPlayed, a.losses = a.matchesPlayed+1, a.losses+1
			b.matchesPlayed, b.wins, b.points = b.matchesPlayed+1, b.wins+1, b.points+3
		default:
			return fmt.Errorf("incorrect input")
		}
		teams[a.name], teams[b.name] = a, b
	}
	var sortedTeams []team
	for _, value := range teams {
		sortedTeams = append(sortedTeams, value)
	}
	sort.Slice(sortedTeams, func(i, j int) bool {
		if sortedTeams[i].points == sortedTeams[j].points {
			return sortedTeams[i].name < sortedTeams[j].name
		}
		return sortedTeams[i].points > sortedTeams[j].points
	})
	fmt.Fprintf(w, "%-31s| MP |  W |  D |  L |  P\n", "Team")
	for _, team := range sortedTeams {
		fmt.Fprintf(w, "%-31s| %2d | %2d | %2d | %2d | %2d\n", team.name, team.matchesPlayed, team.wins,
			team.draws, team.losses, team.points)
	}
	return nil
}
