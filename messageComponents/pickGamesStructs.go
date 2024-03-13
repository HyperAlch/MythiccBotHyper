package messageComponents

import "slices"

type GamesDropdown interface {
	Filter(item string, userRoles []string) bool
	GetContent() string
	GetCustomId() string
	GetDefaultMessage() string
}

type AddDropdown struct{}
type RemoveDropdown struct{}

func (_ AddDropdown) Filter(item string, userRoles []string) bool {
	return !slices.Contains(userRoles, item)
}

func (_ AddDropdown) GetContent() string {
	return "Please select the games you're interested in"
}

func (_ AddDropdown) GetCustomId() string {
	return "pick-games-add-execute"
}

func (_ AddDropdown) GetDefaultMessage() string {
	return "You have already selected all available games"
}

func (_ RemoveDropdown) Filter(item string, userRoles []string) bool {
	return slices.Contains(userRoles, item)
}

func (_ RemoveDropdown) GetContent() string {
	return "Please select the games you would like to remove"
}

func (_ RemoveDropdown) GetCustomId() string {
	return "pick-games-remove-execute"
}

func (_ RemoveDropdown) GetDefaultMessage() string {
	return "There are no games to remove"
}
