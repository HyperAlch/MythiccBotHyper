package messageComponents

import (
	"MythiccBotHyper/datatype"
)

var (
	MessageComponentHandlers = datatype.InteractionMap{
		"pick-games-add":    pickGamesAdd,
		"pick-games-remove": pickGamesRemove,
	}
)
