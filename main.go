package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"
	"time"
)

var (
	writer io.Writer
	reader *bufio.Reader
)

func main() {
	rand.Seed(time.Now().UnixNano())
	stages := getStages()
	Play(os.Stdout, os.Stdin, stages)
}

// Play starts the game
func Play(iowriter io.Writer, ioreader io.Reader, stages []stage) {
	writer = iowriter
	reader = bufio.NewReader(ioreader)

	basicStep("Tutaj Hubert Urbański! Witamy w Milionerach\n")
	name := openQuestion("Proszę przedstaw się. Jak masz na imię?\n")
	basicStep("Witamy Ciebie %v bardzo serdecznie w naszym programie!\n", name)

	gameState := NewGameState()

	for _, stage := range stages {
		fmt.Fprintf(writer, "Oto pytanie za %d zł:\n", stage.money)
		if answer := stage.Ask(gameState); answer.isCorrect {
			gameState.currentMoney = stage.money
			if stage.isGuaranteed {
				gameState.guaranteedMoney = stage.money
				basicStep("To jest poprawna odpowiedź! Zdobyłeś GWARANTOWANE %d zł.\n", stage.money)
			} else {
				basicStep("To jest poprawna odpowiedź! Zdobyłeś %d zł.\n", stage.money)
			}
		} else if answer.hasNotResigned {
			basicStep("Przykro mi. To jest niestety błędna odpowiedź.\n")
			gameState.currentMoney = gameState.guaranteedMoney
			break
		} else {
			basicStep("Postanowiłeś zrezygnować.\n")
			break
		}
	}

	basicStep("To już koniec. Wygrałeś %d zł.\nDziękuję bardzo za wspólną grę. Pozdrawiam i do zobaczenia w następnym odcinku!\n", gameState.currentMoney)
}

func basicStep(text string, a ...interface{}) {
	fmt.Fprintf(writer, text, a...)
	time.Sleep(time.Second)
}

func openQuestion(text string) string {
	fmt.Fprint(writer, text)
	result, err := reader.ReadString('\n')
	check(err)
	return strings.TrimSpace(result)
}
