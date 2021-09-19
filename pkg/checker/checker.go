package checker

import (
	"context"
	"fmt"
)

const (
	Light    = 4
	Full     = 8
	Alliance = 24
)

var (
	Jobs                                                         = JobMap{
		"PLD": "Tank",
		"WAR": "Tank",
		"DRK": "Tank",
		"GNB": "Tank",
		"WHM": "PHealer",
		"AST": "PHealer",
		"SCH": "SHealer",
		"SGE": "SHealer",
		"DRG": "PDPS",
		"NIN": "PDPS",
		"MNK": "PDPS",
		"SAM": "PDPS",
		"RPR": "PDPS",
		"BLM": "MDPS",
		"SMN": "MDPS",
		"RDM": "MDPS",
		"BRD": "RDPS",
		"DNC": "RDPS",
		"MCH": "RDPS",
	}
)

type Tank Kind 
type PHealer Kind

type Player struct {
	Name  string
	Roles []string
}

type Composition struct {
	Size   int
	Unique bool
}

type Variant struct {
	LightParty LightParty
	FullParty  FullParty
}

type LightParty struct {
	Tank1   map[string]string
	Healer1 map[string]string
	DPS1    map[string]string
	DPS2    map[string]string
}

// A FullParty is comprised of two tanks, one pure healer, one shield healer, one ranged physical DPS, one magical DPS, and two physical melee DPS. Each type is a Slot. 
type FullParty struct {
	Tank1 PlayerAssignment
	Tank2  PlayerAssignment
	PHealer PlayerAssignment
	SHealer PlayerAssignment
	RDPS1  PlayerAssignment
	MDPS1    PlayerAssignment
	PDPS1    PlayerAssignment
	PDPS2    PlayerAssignment
}

type Kind string
type Slot Kind

// JobMap maps job names to what kind of job they are.
type JobMap map[string]string
// PlayerAssignment contains the map of player name to job it is playing.
type PlayerAssignment map[string]string

func (p *PlayerAssignment) Assigned() bool {
	return len(*p) != 0
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func AssignRole(player, role string) map[string]string {
	return map[string]string{player: role}
}

// CheckCompositionVariants takes a slice of Players and a desired Composition, and determines how many Variants can be made out of the supplied inputs. It returns a slice of Variant, or an error if one is occurred.
func CheckCompositionVariants(ctx context.Context, players []Player, composition Composition) ([]Variant, error) {
	return nil, nil
}

// we need to accept a slice of values belonging to only one of many types, which are then assigned a slot in the returning data structure. only return the first data structure.
// for each player, what job do they have? for each job, what role is it? we need to know our available role list and match that against our desired composition.
// if we do not have enough players with the right roles to fill our composition, return an error.
// how do we figure out what combination of input satisfies our desired composition?
// a full party is comprised of two tanks, one shield healer and one pure healer, two physical dps and two ranged dps. We need to look up the jobs each player has and figure out what kind they are. If there is an open slot in our party for that kind,
// their job qualifies and the player is Assigned for that FullParty (cannot be used again to fill another slot).
func Allocate(players []Player) (FullParty, error) {
	var assignedPlayers []string
	fullParty := FullParty{}
	for _, player := range players {
		fmt.Printf("Player %s has the following roles: %s\n", player.Name, player.Roles)
		for _, role := range player.Roles {
			switch Jobs[role] {
			case "Tank":
				if !fullParty.Tank1.Assigned() && !contains(assignedPlayers, player.Name) {
				    fullParty.Tank1 = AssignRole(player.Name, role)
					assignedPlayers = append(assignedPlayers, player.Name)
					continue
				}
				if !fullParty.Tank2.Assigned() && !contains(assignedPlayers, player.Name) {
					fullParty.Tank2 = AssignRole(player.Name, role)
					assignedPlayers = append(assignedPlayers, player.Name)
					continue
				}
			case "PHealer":
				if !fullParty.PHealer.Assigned() && !contains(assignedPlayers, player.Name) {
					fullParty.PHealer = AssignRole(player.Name, role)
					assignedPlayers = append(assignedPlayers, player.Name)
					continue
				}
			case "SHealer":
				if !fullParty.SHealer.Assigned() && !contains(assignedPlayers, player.Name) {
					fullParty.SHealer = AssignRole(player.Name, role)
					assignedPlayers = append(assignedPlayers, player.Name)
					continue
				}
			case "PDPS":
				if !fullParty.PDPS1.Assigned() && !contains(assignedPlayers, player.Name) {
					fullParty.PDPS1 = AssignRole(player.Name, role)
					assignedPlayers = append(assignedPlayers, player.Name)
					continue
				}
				if !fullParty.PDPS2.Assigned() && !contains(assignedPlayers, player.Name) {
					fullParty.PDPS2 = AssignRole(player.Name, role)
					assignedPlayers = append(assignedPlayers, player.Name)
					continue
				}
			case "MDPS":
				if !fullParty.MDPS1.Assigned() && !contains(assignedPlayers, player.Name) {
					fullParty.MDPS1 = AssignRole(player.Name, role)
					assignedPlayers = append(assignedPlayers, player.Name)
					continue
				}
			case "RDPS": 
				if !fullParty.RDPS1.Assigned() && !contains(assignedPlayers, player.Name) {
					fullParty.RDPS1 = AssignRole(player.Name, role)
					assignedPlayers = append(assignedPlayers, player.Name)
					continue
				}
			default:
				continue
			}
		}
	}
	fmt.Println(assignedPlayers, "were assigned")
	return fullParty, nil
}
