package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Player struct {
	DiceCount int
	Points    int
}

func main() {
	// Input jumlah pemain dan jumlah dadu
	var numPlayers, numDice int
	fmt.Print("Masukkan jumlah pemain: ")
	fmt.Scan(&numPlayers)
	fmt.Print("Masukkan jumlah dadu per pemain: ")
	fmt.Scan(&numDice)

	players := make([]Player, numPlayers)
	for i := range players {
		players[i] = Player{DiceCount: numDice, Points: 0}
	}

	winnerIndex := playGame(players, numDice)
	fmt.Printf("Pemain %d adalah pemenang dengan %d poin!\n", winnerIndex+1, players[winnerIndex].Points)
}

func playGame(players []Player, numDice int) int {
	activePlayers := len(players)

	for activePlayers > 1 {
		for i := range players {
			player := &players[i]
			if player.DiceCount <= 0 {
				continue
			}

			fmt.Printf("Pemain %d melempar dadu: ", i+1)
			dicePerRound := []int{}
			for j := 0; j < player.DiceCount; j++ {
				dice := rollDice()
				dicePerRound = append(dicePerRound, dice)

			}
			fmt.Print(dicePerRound)
			evaluate(i, dicePerRound, players)
			fmt.Println()

			if player.DiceCount <= 0 {
				activePlayers--
			}
		}
	}

	return winner(players)
}

func rollDice() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(6) + 1
}

func winner(players []Player) int {
	winnerIndex := 0
	maxPoints := -1

	for i, player := range players {
		if player.Points > maxPoints {
			winnerIndex, maxPoints = i, player.Points
		}
	}

	return winnerIndex
}

// evaluate adalah fungsi untuk mengevaluasi hasil lemparan dadu dan mengubah jumlah dadu pemain
func evaluate(index int, dicePerRound []int, players []Player) {
	var nextPlayer int
	for _, dice := range dicePerRound {
		switch dice {
		case 6:
			players[index].Points++    // Pemain mendapatkan poin
			players[index].DiceCount-- // Kurangi jumlah dadu pemain
		case 1:
			nextPlayer = index + 1
			if index == len(players)-1 {
				nextPlayer = 0
			}
			players[index].DiceCount--      // Kurangi jumlah dadu pemain
			players[nextPlayer].DiceCount++ // Pemain berikutnya mendapatkan dadu
		}
	}
}
