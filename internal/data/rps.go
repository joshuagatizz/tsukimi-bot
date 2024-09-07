package data

type RockPaperScissorsGame struct {
	FirstUserId     string
	SecondUserId    string
	FirstUserChoice string
}

// TODO: add mutex
var ActiveRockPaperScissorsGame = make(map[string]*RockPaperScissorsGame)
