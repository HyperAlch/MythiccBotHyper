package messageComponents

import (
	"MythiccBotHyper/datatype"
)

var (
	MessageComponentHandlers = datatype.InteractionMap{
		"pick-games-add":            pickGamesAdd,
		"pick-games-remove":         pickGamesRemove,
		"pick-games-add-execute":    pickGamesAddExecute,
		"pick-games-remove-execute": pickGamesRemoveExecute,
	}
)
