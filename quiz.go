package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func (quiz quiz) Ask(gameState *GameState) quizResult {

	fmt.Fprintln(writer, quiz.question)
	time.Sleep(time.Second)

	all := []string{
		quiz.correct, quiz.incorrect[0], quiz.incorrect[1], quiz.incorrect[2],
	}
	rand.Shuffle(len(all), func(i, j int) {
		all[i], all[j] = all[j], all[i]
	})

	var selections map[string]string
	var userSelection string
	for {
		selections = map[string]string{}
		validSelections := []string{}
		for i, possibleAnswer := range all {
			selection := strconv.Itoa(i + 1)
			validSelections = append(validSelections, selection)
			selections[selection] = possibleAnswer

			fmt.Fprintf(writer, "%v - %v\n", selection, possibleAnswer)
			time.Sleep(time.Second)
		}

		if gameState.canAskAudience {
			validSelections = append(validSelections, "7")
			fmt.Fprintln(writer, "7 - Publiczność")
		}
		if gameState.canUseHalfByHalf {
			validSelections = append(validSelections, "8")
			fmt.Fprintln(writer, "8 - Pół na pół")
		}
		if gameState.canAskFrind {
			validSelections = append(validSelections, "9")
			fmt.Fprintln(writer, "9 - Telefon do przyjaciela")
		}

		validSelections = append(validSelections, "0")
		fmt.Fprintln(writer, "0 - Rezygnacja")

		userSelection = getSelection(validSelections)

		if userSelection == "7" {
			fmt.Fprintln(writer, "Droga publiczności, prosimy o Waszą pomoc.")
			time.Sleep(3 * time.Second)

			no := 100
			ch := make(chan string, no)
			for i := 0; i < no; i++ {
				go func() {
					if chance := rand.Float64(); chance < 0.15 {
						ch <- quiz.correct
					} else {
						ch <- all[rand.Intn(len(all))]
					}
				}()
			}

			votesCount := map[string]int{}
			for i := 0; i < no; i++ {
				votesCount[<-ch]++
			}

			fmt.Fprintln(writer, "Tak się przedstawiają wyniki głosowania publiczności:")
			for key, procentage := range votesCount {
				fmt.Fprintf(writer, "%v - %d%%\n", key, procentage)
			}

			gameState.canAskAudience = false
		} else if userSelection == "8" {
			all = []string{
				quiz.correct,
				quiz.incorrect[rand.Intn(len(quiz.incorrect))],
			}
			fmt.Fprintln(writer, "Wybrałeś pół na pół, dlatego odrzucam dwie błędne odpowiedzi.")
			gameState.canUseHalfByHalf = false
		} else if userSelection == "9" {
			fmt.Fprintln(writer, "Dzwonimy do Twojego przyjaciela Gabika. Dawno z nim nie rozmawialiśmy.")
			time.Sleep(time.Second)
			fmt.Fprintln(writer, "dryn, dryn...")
			time.Sleep(time.Second)
			fmt.Fprintln(writer, "dryn, dryn...")
			time.Sleep(time.Second)
			fmt.Fprintln(writer, "ADAM: Hallo")
			time.Sleep(time.Second)
			fmt.Fprintln(writer, "HUBERT: Cześć Adamie tutaj Hubert Urbański [...]")
			time.Sleep(3 * time.Second)
			if chance := rand.Float64(); chance < 0.1 {
				fmt.Fprintf(writer, "ADAM: Na pewno poprawna odpowiedź to: %v\n", quiz.correct)
			} else if chance < 0.3 {
				fmt.Fprintf(writer, "ADAM: Wydaje mi się, że poprawna odpowiedź to: %v\n", quiz.correct)
			} else if chance < 0.8 {
				fmt.Fprintf(writer, "ADAM: Wydaje mi się, że poprawna odpowiedź to: %v\n", quiz.incorrect[rand.Intn(len(quiz.incorrect))])
			} else {
				fmt.Fprintln(writer, "ADAM: Nie mam pojęcia. Radź se sam.")
			}
			gameState.canAskFrind = false
		} else {
			break
		}
	}

	userAnswer, ok := selections[userSelection]
	return quizResult{
		hasNotResigned: userSelection != "0",
		isCorrect:      ok && userAnswer == quiz.correct,
	}
}

func getSelection(validSelections []string) string {
	for {
		line, err := reader.ReadString('\n')
		check(err)
		userSelection := strings.TrimSpace(line)
		for _, validSelecton := range validSelections {
			if userSelection == validSelecton {
				return userSelection
			}
		}
	}
}

type quiz struct {
	question  string
	correct   string
	incorrect [3]string
}

type quizResult struct {
	hasNotResigned bool
	isCorrect      bool
}
