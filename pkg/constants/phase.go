package constants

type Phases map[string]int

type BotActions map[string]Phases

var OshnoBot = BotActions{
	RequestProviderPhase: Phases{
		FullNameInput: 1,
	},
}
