package tournament

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

type team struct {
	name          string
	matchesPlayed int
	wins          int
	draws         int
	losses        int
	points        int
}

var teamsInTournament map[string]*team

func Tally(r io.Reader, w io.Writer) error {
	teamsInTournament = make(map[string]*team)

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		s := strings.SplitN(scanner.Text(), ";", 3)
		if len(s) != 3 && len(s) > 1 {
			//fmt.Println(s)
			//fmt.Println(len(s))
			return errors.New("incorrect size")
		}
		if len(s) == 3 {
			err := calculateMatchResults(s)
			if err != nil {
				return err
			}
		}

		//fmt.Println(scanner.Text()) // Println will add back the final '\n'
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		return err
	}

	w.Write([]byte(printResults()))
	return nil
	//return errors.New("")
}

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
		teamsInTournament[s[0]].matchesPlayed += 1
		teamsInTournament[s[0]].draws += 1
		teamsInTournament[s[0]].points += 1

		teamsInTournament[s[1]].matchesPlayed += 1
		teamsInTournament[s[1]].draws += 1
		teamsInTournament[s[1]].points += 1
	case "win":
		teamsInTournament[s[0]].matchesPlayed += 1
		teamsInTournament[s[0]].wins += 1
		teamsInTournament[s[0]].points += 3

		teamsInTournament[s[1]].matchesPlayed += 1
		teamsInTournament[s[1]].losses += 1
	case "loss":
		teamsInTournament[s[0]].matchesPlayed += 1
		teamsInTournament[s[0]].losses += 1

		teamsInTournament[s[1]].matchesPlayed += 1
		teamsInTournament[s[1]].wins += 1
		teamsInTournament[s[1]].points += 3
	default:
		return errors.New("incorrect input")
	}
	return nil
}

func printResults() string {
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

	var result strings.Builder

	result.WriteString(
		fmt.Sprintf("%-31s| MP |  W |  D |  L |  P\n", "Team"),
	)

	for _, team := range teams {
		s := fmt.Sprintf("%-31s|  %d |  %d |  %d |  %d |  %d\n",
			team.name,
			team.matchesPlayed,
			team.wins,
			team.draws,
			team.losses,
			team.points)
		result.WriteString(s)
	}
	//result.WriteString("\n")

	return result.String()
}
