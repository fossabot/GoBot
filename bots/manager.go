package bots

import "github.com/Mrs4s/MiraiGo/client"

type Manager struct {
	Bots []BotInfo
}

type BotInfo struct {
	Client *client.QQClient
	Statue bool
}
