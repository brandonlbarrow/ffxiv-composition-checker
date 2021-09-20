package checker

import (
	"fmt"
	"strings"
)

const (
	CompUndefined CompType = iota
	Any
	Balanced
	Unique

	KindUndefined Kind = iota
	Tank
	PureHealer
	ShieldHealer
	PhysicalMelee
	PhysicalRanged
	MagicalRanged

	KindTypeUndefined KindType = iota
	KindTank
	KindHealer
	KindDPS

	FullPartyWeight int = 18
)

var (
	Jobs = JobMap{
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
	eMap = map[Kind]KindType{
		Tank:           KindTank,
		PureHealer:     KindHealer,
		ShieldHealer:   KindHealer,
		PhysicalMelee:  KindDPS,
		PhysicalRanged: KindDPS,
		MagicalRanged:  KindDPS,
	}
	weight = map[KindType]int{
		KindTank:   1,
		KindHealer: 2,
		KindDPS:    3,
	}
)

type Kind int
type KindType int
type CompType int

func (k Kind) TypeOf() KindType {
	return eMap[k]
}

type Player struct {
	Name  string
	Roles []string
}

type CompArgs struct {
	Format CompType
}

// A FullParty is comprised of two tanks, two healers, and four DPS.
type FullParty struct {
	comp     RoleAssignments
	compType CompType
	keys     []int
}

// JobMap maps job names to what kind of job they are.
type JobMap map[string]Kind

type RoleAssignments map[KindOf]PlayerAssignment

func (r RoleAssignments) Composition() string {
	comp := []string{}
	for k, v := range r {
		comp = append(comp, fmt.Sprintf("%v %d: %s", k.KindType, k.RoleNumber, v))
	}
	return strings.Join(comp, "\n")
}

type KindOf struct {
	Kind       `json:"kind"`
	KindType   `json:"kind_type"`
	RoleNumber int `json:"role_number"`
}

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

func (f FullParty) Size() int {
	return 8
}

func generateRoleAssignmentsAny() RoleAssignments {
	r := RoleAssignments{}
	for i := 0; i < 8; i++ {
		if i < 2 {
			k := KindOf{
				Kind:       KindUndefined,
				KindType:   KindTank,
				RoleNumber: i + 1,
			}
			r[k] = PlayerAssignment{}
		}
		if i >= 2 && i < 4 {
			k := KindOf{
				Kind:       KindUndefined,
				KindType:   KindHealer,
				RoleNumber: i + 1,
			}
			r[k] = PlayerAssignment{}
		} 
		if i >= 4 && i <=8 {
			k := KindOf{
				Kind:       KindUndefined,
				KindType:   KindDPS,
				RoleNumber: i + 1,
			}
			r[k] = PlayerAssignment{}
		}
	}
	return r
}

// Initialize takes the desired composition and builds an internal RoleAssignments that matches the user's desire.
func (f *FullParty) Initialize(args *CompArgs) {
	if args == nil {
		return
	}
	switch args.Format {
	case Any:
		f.comp = generateRoleAssignmentsAny()
		f.compType = Any
	}
}

func (f *FullParty) Composition() string {
	return f.comp.Composition()
}

// we need to accept a slice of values belonging to only one of many types, which are then assigned a slot in the returning data structure. only return the first data structure.
// for each player, what job do they have? for each job, what role is it? we need to know our available role list and match that against our desired composition.
// if we do not have enough players with the right roles to fill our composition, return an error.
// how do we figure out what combination of input satisfies our desired composition?
// a full party is comprised of two tanks, one shield healer and one pure healer, two physical dps and two ranged dps. We need to look up the jobs each player has and figure out what kind they are. If there is an open slot in our party for that kind,
// their job qualifies and the player is Assigned for that FullParty (cannot be used again to fill another slot).
func Allocate(players []Player, args CompArgs) (FullParty, error) {
	var assignedPlayers []string
	fullParty := FullParty{}
	fullParty.Initialize(&args)
	for _, player := range players {
		// for k, v := range fullParty.comp {
		// 	switch k.KindType {
		// 	case KindTank:
		// 	}
		// 	fmt.Println(v)
		// }
		for _, role := range player.Roles {
			for k, v := range fullParty.comp {
				switch k.KindType {
				case KindTank:
					if eMap[Jobs[role]] == KindTank {
						switch fullParty.compType {
						case Any:
						}
						if notYetAssigned(assignedPlayers, player.Name, v) {
							v[player.Name] = role
							assignedPlayers = append(assignedPlayers, player.Name)
							continue
						}
					}
				case KindHealer:
					if eMap[Jobs[role]] == KindHealer {
						switch fullParty.compType {
						case Any:
						}
						if notYetAssigned(assignedPlayers, player.Name, v) {
							v[player.Name] = role
							assignedPlayers = append(assignedPlayers, player.Name)
							continue
						}
					}
				case KindDPS:
					if eMap[Jobs[role]] == KindDPS {
						switch fullParty.compType {
						case Any:
						}
						if notYetAssigned(assignedPlayers, player.Name, v) {
							v[player.Name] = role
							assignedPlayers = append(assignedPlayers, player.Name)
						}
					}
				}
			}
		}
	}
	fmt.Println(assignedPlayers, "were assigned")
	fmt.Println(IsValidComp(players, args))
	fmt.Println(len(fullParty.comp))

	return fullParty, nil
}

// IsValidComp determines if a combined []Player and CompArgs results in at least one valid party and returns true if so. If not, this function returns false.
func IsValidComp(players []Player, args CompArgs) bool {
	// we will calculate the possibility of each variant using the weight of the roles, which should add up to a specific number
	// TODO after implementing this, merge logic into the main checker func logic
	w := 0
	variants := make([]int, len(players))
	fmt.Println(variants)
	return w == FullPartyWeight
}

func notYetAssigned(players []string, name string, assignment PlayerAssignment) bool {
	return !contains(players, name) && len(assignment) == 0
}
