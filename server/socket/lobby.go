package socket

import (
	"bomberman-dom/models"
	"log"
)

type Lobby struct {
	timer    *Timer
	userList OnlineList
}

func (l *Lobby) startLobby() {
	for {
		if l.getAmountOfPlayers() < 1 {
			return
		}

		l.timer.C = make(chan models.LobbyState)
		go l.timer.startTimer(0)
		for lobbyState := range l.timer.C {
			log.Println("aboba", lobbyState)
			l.broadcastMessage(SerializeData("ama_boy_next_door", lobbyState))
		}
	}
}

func (l *Lobby) isUsernameExists(username string) bool {
	_, exists := l.userList[username]
	return exists
}

func (l *Lobby) getAmountOfPlayers() int {
	return len(l.userList)
}

func (l *Lobby) addPlayer(c *Client) {
	c.lobby = l
	l.userList[c.username] = c

	if l.getAmountOfPlayers() == 1 {
		go l.startLobby()
	}
}

func (l *Lobby) removePlayer(username string) {
	delete(l.userList, username)
}

func (l *Lobby) broadcastMessage(data Event) {
	for _, client := range l.userList {
		client.egress <- data
	}
}
