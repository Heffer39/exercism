//Package tournament tallies the results of a small football competition
package tournament

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

// team captures the name and stats of a team for the competition
type team struct {
	name          string
	matchesPlayed int
	wins          int
	draws         int
	losses        int
	points        int
}

var teams map[string]team

// Tally reads in an input file, calculates the match results, and writes back
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
		if len(s) == 3 {
			if err := calculateMatchResults(s); err != nil {
				return err
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	sortedTeams := sortedTeamList()
	printResults(sortedTeams, w)
	return nil
}

// calculateMatchResults updates the stats for each team, based on the provided match information
func calculateMatchResults(s []string) error {
	a, b := teams[s[0]], teams[s[1]]
	a.name, b.name = s[0], s[1]

	switch s[2] {
	case "draw":
		a.matchesPlayed++
		a.draws++
		a.points++

		b.matchesPlayed++
		b.draws++
		b.points++
	case "win":
		a.matchesPlayed++
		a.wins++
		a.points += 3

		b.matchesPlayed++
		b.losses++
	case "loss":
		a.matchesPlayed++
		a.losses++

		b.matchesPlayed++
		b.wins++
		b.points += 3
	default:
		return errors.New("incorrect input")
	}

	teams[a.name], teams[b.name] = a, b
	return nil
}

// sortedTeamList returns a sorted list of all the teams in the tournament
// Teams are sorted by points
// If teams are tied on points, they are instead sorted alphabetically
func sortedTeamList() []team {
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

	return sortedTeams
}

// printResults writes the results to the io.Writer, displaying the tabular data of each team and their stats
func printResults(teams []team, w io.Writer) {
	fmt.Fprintf(w, "%-31s| MP |  W |  D |  L |  P\n", "Team")

	for _, team := range teams {
		fmt.Fprintf(w, "%-31s| %2d | %2d | %2d | %2d | %2d\n",
			team.name,
			team.matchesPlayed,
			team.wins,
			team.draws,
			team.losses,
			team.points)
	}
}
