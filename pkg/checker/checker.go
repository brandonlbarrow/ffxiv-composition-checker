package checker

import (
	"context"
	"fmt"
)

const (
	Tank Kind = iota
	PureHealer
	ShieldHealer
	PhysicalMelee
	PhysicalRanged
	MagicalRanged
)

var (
	Jobs                                                         = JobMap{
		"PLD": Tank,
		"WAR": Tank,
		"DRK": Tank,
		"GNB": Tank,
		"WHM": PureHealer,
		"AST": PureHealer,
		"SCH": ShieldHealer,
		"SGE": ShieldHealer,
		"DRG": PhysicalMelee,
		"NIN": PhysicalMelee,
		"MNK": PhysicalMelee,
		"SAM": PhysicalMelee,
		"RPR": PhysicalMelee,
		"BLM": MagicalRanged,
		"SMN": MagicalRanged,
		"RDM": MagicalRanged,
		"BRD": PhysicalRanged,
		"DNC": PhysicalRanged,
		"MCH": PhysicalRanged,
	}
	FParty = map[string]string{
		"Tank1": "",
		"Tank2": "",
		"PHealer1": "",
		"SHealer1": "",
		"PDPS1": "",
		"PDPS2": "",
		"RDPS1": "",
		"RPDS2": "",
	}
	eMap = JobMapEnum{
		"PLD": Tank,

	}
)

type Kind int

type Player struct {
	Name  string
	Roles []string
}

type CompArgs struct {
	Size   int
	Unique bool
}

type Variant struct {
}

// A LightParty is comprised of one tank, one healer, and two DPS.
type LightParty struct {
	Tank1   map[string]string
	Healer1 map[string]string
	DPS1    map[string]string
	DPS2    map[string]string
}

// A FullParty is comprised of two tanks, two healers, and four DPS.
type FullParty struct {
	Tank1 PlayerAssignment
	Tank2 PlayerAssignment
	Healer1 PlayerAssignment
	Healer2 PlayerAssignment
	DPS1 PlayerAssignment
	DPS2 PlayerAssignment
	DPS3 PlayerAssignment
	DPS4 PlayerAssignment
}

// JobMap maps job names to what kind of job they are.
type JobMap map[string]Kind
type JobMapEnum map[string]Kind
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

type Composition interface{}

func GetComposition(args CompArgs) Composition {
	if args.Size == 8 {
		return FullParty{}
	} else {
		return LightParty{}
	}
}

// CheckCompositionVariants takes a slice of Players and a desired Composition, and determines how many Variants can be made out of the supplied inputs. It returns a slice of Variant, or an error if one is occurred.
func CheckCompositionVariants(ctx context.Context, players []Player, composition CompArgs) ([]Variant, error) {
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
			case Tank:
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
			case PureHealer:
				if !fullParty.Healer1.Assigned() && !contains(assignedPlayers, player.Name) {
					fullParty.Healer1 = AssignRole(player.Name, role)
					assignedPlayers = append(assignedPlayers, player.Name)
					continue
				}
			case ShieldHealer:
				if !fullParty.Healer2.Assigned() && !contains(assignedPlayers, player.Name) {
					fullParty.Healer2 = AssignRole(player.Name, role)
					assignedPlayers = append(assignedPlayers, player.Name)
					continue
				}
			case PhysicalMelee:
				if !fullParty.DPS1.Assigned() && !contains(assignedPlayers, player.Name) {
					fullParty.DPS1 = AssignRole(player.Name, role)
					assignedPlayers = append(assignedPlayers, player.Name)
					continue
				}
				if !fullParty.DPS2.Assigned() && !contains(assignedPlayers, player.Name) {
					fullParty.DPS2 = AssignRole(player.Name, role)
					assignedPlayers = append(assignedPlayers, player.Name)
					continue
				}
			case MagicalRanged:
				if !fullParty.DPS3.Assigned() && !contains(assignedPlayers, player.Name) {
					fullParty.DPS3 = AssignRole(player.Name, role)
					assignedPlayers = append(assignedPlayers, player.Name)
					continue
				}
			case PhysicalRanged: 
				if !fullParty.DPS4.Assigned() && !contains(assignedPlayers, player.Name) {
					fullParty.DPS4 = AssignRole(player.Name, role)
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
