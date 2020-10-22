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

var teamsInTournament map[string]*team

// Tally reads in an input file, calculates the match results, and writes back
// the results to the io.Writer
func Tally(r io.Reader, w io.Writer) error {
	teamsInTournament = make(map[string]*team)
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		s := strings.SplitN(scanner.Text(), ";", 3)
		if len(s) != 3 && len(s) > 1 {
			return errors.New("incorrect size")
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

	teamsList := sortedTeamList()
	if _, err := w.Write([]byte(printResults(teamsList))); err != nil {
		return err
	}
	return nil
}

// calculateMatchResults updates the stats for each team, based on the provided match information
func calculateMatchResults(s []string) error {
	if _, ok := teamsInTournament[s[0]]; !ok {
		teamsInTournament[s[0]] = &team{name: s[0]}
	}
	if _, ok := teamsInTournament[s[1]]; !ok {
		teamsInTournament[s[1]] = &team{name: s[1]}
	}

	if s[0] == s[1] {
		return errors.New("team cannot play itself")
	}

	switch s[2] {
	case "draw":
		teamsInTournament[s[0]].matchesPlayed++
		teamsInTournament[s[0]].draws++
		teamsInTournament[s[0]].points++

		teamsInTournament[s[1]].matchesPlayed++
		teamsInTournament[s[1]].draws++
		teamsInTournament[s[1]].points++
	case "win":
		teamsInTournament[s[0]].matchesPlayed++
		teamsInTournament[s[0]].wins++
		teamsInTournament[s[0]].points += 3

		teamsInTournament[s[1]].matchesPlayed++
		teamsInTournament[s[1]].losses++
	case "loss":
		teamsInTournament[s[0]].matchesPlayed++
		teamsInTournament[s[0]].losses++

		teamsInTournament[s[1]].matchesPlayed++
		teamsInTournament[s[1]].wins++
		teamsInTournament[s[1]].points += 3
	default:
		return errors.New("incorrect input")
	}
	return nil
}

// sortedTeamList returns a sorted list of all the teams in the tournament
// Teams are sorted by points
// If teams are tied on points, they are instead sorted alphabetically
func sortedTeamList() []team {
	var teams []team
	for _, value := range teamsInTournament {
		teams = append(teams, *value)
	}

	sort.Slice(teams, func(i, j int) bool {
		if teams[i].points == teams[j].points {
			return teams[i].name < teams[j].name
		}
		return teams[i].points > teams[j].points
	})

	return teams
}

// printResults returns a formatted string, displaying the tabular data of each team and their stats
func printResults(teams []team) string {
	var result strings.Builder

	result.WriteString(
		fmt.Sprintf("%-31s| MP |  W |  D |  L |  P\n", "Team"),
	)

	for _, team := range teams {
		s := fmt.Sprintf("%-31s| %2d | %2d | %2d | %2d | %2d\n",
			team.name,
			team.matchesPlayed,
			team.wins,
			team.draws,
			team.losses,
			team.points)
		result.WriteString(s)
	}

	return result.String()
}
