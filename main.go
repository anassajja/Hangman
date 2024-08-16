package main

import (
	"bufio" // bufio est un package utilisé pour les E/S tamponnées
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

var score int

func drawHangman(chances int) {

	colorRed := "\033[31m"
	fmt.Println(string(colorRed))
	switch chances {
	case 9:
		println("         ")
		println("         ")
		println("         ")
		println("         ")
		println("         ")
		println("_____    ")
		break

	case 8:
		println("         ")
		println(" |       ")
		println(" |       ")
		println(" |       ")
		println(" |       ")
		println("_|___    ")
		break

	case 7:
		println("  ____   ")
		println(" |       ")
		println(" |       ")
		println(" |       ")
		println(" |       ")
		println("_|___    ")
		break
	case 6:
		println("  ____   ")
		println(" |    |  ")
		println(" |       ")
		println(" |       ")
		println(" |       ")
		println("_|___    ")
		break

	case 5:
		println("  ____   ")
		println(" |    |  ")
		println(" |    O  ")
		println(" |       ")
		println(" |       ")
		println("_|___    ")
		break

	case 4:
		println("  ____   ")
		println(" |    |  ")
		println(" |    O_ ")
		println(" |       ")
		println(" |       ")
		println("_|___    ")
		break

	case 3:
		println("  ____   ")
		println(" |    |  ")
		println(" |   _O_ ")
		println(" |       ")
		println(" |       ")
		println("_|___    ")
		break
	case 2:
		println("  ____   ")
		println(" |    |  ")
		println(" |   _O_ ")
		println(" |    |  ")
		println(" |       ")
		println("_|___    ")
		break
	case 1:
		println("  ____   ")
		println(" |    |  ")
		println(" |   _O_ ")
		println(" |    |  ")
		println(" |     \\  ")
		println("_|___    ")
		break

	case 0:
		println("  ____   ")
		println(" |    |  ")
		println(" |   _O_ ")
		println(" |    |  ")
		println(" |   / \\ ")
		println("_|___    ")
		gamover()
		colorReset := "\033[0m"
		fmt.Println(string(colorReset))
		fmt.Println("You want to replay ? (YES / NO)")
		replay()
	}

}

func checkWin(answer, word string) bool {
	if answer != word {
		colorReset := "\033[0m"
		fmt.Println(string(colorReset))
		return false
	}
	return true
}

func gamePlay1(themes string) {
	var theme string
	switch themes {
	case "cars1":
		theme = "themes/cars1.txt"
		break
	case "food1":
		theme = "themes/food1.txt"
		break
	case "country1":
		theme = "themes/pays1.txt"
		break
	}

	var words []string
	var word []string
	var answer []string
	var letterused []string

	var lifes int
	lifes = 10

	readFile1, err := os.Open(theme)
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile1) //Use bufio.NewScanner() function to create the file scanner.
	fileScanner.Split(bufio.ScanLines)         //Use bufio.ScanLines() function with the scanner to read a file line by line.

	for fileScanner.Scan() { //use the filescanner.Scan() function in a for loop to get each line and process it.
		words = append(words, fileScanner.Text())
	}
	readFile1.Close()

	rand.Seed(time.Now().UnixNano()) // This function allow to pick a randomly word
	randomElement := words[rand.Intn(len(words))]

	for i := range randomElement {
		word = append(word, string(randomElement[i]))
	}

	/* Print word */
	//fmt.Println(word)

	for i := 0; i < len(randomElement); i++ {
		answer = append(answer, ".")
	}
	fmt.Println(answer)
	for lifes > 0 {
		colorReset := "\033[0m"
		fmt.Println(string(colorReset))
		var inputAnswer string
		var checkAnswer = false

		fmt.Print("Type a letter : ")
		fmt.Scan(&inputAnswer)

		fmt.Println("----------------------")

		for i := range word {
			if inputAnswer == word[i] {
				checkAnswer = true
				answer[i] = inputAnswer
				word[i] = "."
				fmt.Println(answer)
				break
			}
		}

		if checkAnswer {
			score = score + 20
			fmt.Println()
			fmt.Println("Your score is : ", score)
			if checkWin(strings.Join(answer, ""), randomElement) {
				colorGreen := "\033[32m"
				fmt.Println(string(colorGreen))
				fmt.Println("You Win")
				fmt.Println()
				fmt.Println("Your score is : ", score)
				colorReset := "\033[0m"
				fmt.Println(string(colorReset))
				fmt.Println("You want to replay ? (YES / NO)")
				replay()
				break
			}

		} else {
			lifes--
			drawHangman(lifes)
			letterused = append(letterused, inputAnswer)

			if lifes >= 1 {
				fmt.Println("Lettres Used : ", letterused)
				fmt.Println("Attetion ! You still have ", lifes, "lifes ")
			}
		}
	}
}

func gamePlay2(themes string) {
	var theme string
	switch themes {
	case "cars2":
		theme = "themes/cars2.txt"
		break
	case "food2":
		theme = "themes/food2.txt"
		break
	case "country2":
		theme = "themes/pays2.txt"
		break
	}

	var words []string
	var word []string
	var answer []string
	var letterused []string

	var lifes int
	lifes = 10

	readFile, err := os.Open(theme)
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		words = append(words, fileScanner.Text())
	}
	readFile.Close()

	rand.Seed(time.Now().UnixNano())
	randomElement := words[rand.Intn(len(words))]

	for i := range randomElement {
		word = append(word, string(randomElement[i]))
	}

	/* Print word */
	//fmt.Println(word)

	for i := 0; i < len(randomElement); i++ {
		answer = append(answer, ".")
	}
	fmt.Println(answer)
	for lifes > 0 {
		colorReset := "\033[0m"
		fmt.Println(string(colorReset))
		var inputAnswer string
		var checkAnswer = false

		fmt.Print("Type a letter : ")
		fmt.Scan(&inputAnswer)

		fmt.Println("----------------------")

		for i := range word {
			if inputAnswer == word[i] {
				checkAnswer = true
				answer[i] = inputAnswer
				word[i] = "."
				fmt.Println(answer)
				break
			}
		}
		if checkAnswer {
			score = score + 20
			fmt.Println()
			fmt.Println("Your score is : ", score)
			if checkWin(strings.Join(answer, ""), randomElement) {
				colorGreen := "\033[32m"
				fmt.Println(string(colorGreen))
				fmt.Println("You Win")
				fmt.Println()
				fmt.Println("Your score is : ", score)
				colorReset := "\033[0m"
				fmt.Println(string(colorReset))
				fmt.Println("You want to replay ? (YES / NO)")
				replay()
				break
			}

		} else {
			lifes--
			drawHangman(lifes)
			letterused = append(letterused, inputAnswer)
			if lifes >= 1 {
				fmt.Println("Lettres Used : ", letterused)
				fmt.Println("Attetion ! You still have ", lifes, "lifes ")
			}
		}
	}
}

func gamePlay3(themes string) {
	var theme string
	switch themes {
	case "cars3":
		theme = "themes/cars3.txt"
		break
	case "food3":
		theme = "themes/food3.txt"
		break
	case "country3":
		theme = "themes/pays3.txt"
		break
	}

	var words []string
	var word []string
	var answer []string
	var letterused []string

	var lifes int
	lifes = 10

	readFile, err := os.Open(theme)
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		words = append(words, fileScanner.Text())
	}
	readFile.Close()

	rand.Seed(time.Now().UnixNano())
	randomElement := words[rand.Intn(len(words))]

	for i := range randomElement {
		word = append(word, string(randomElement[i]))
	}

	/* Print word */
	//fmt.Println(word)

	for i := 0; i < len(randomElement); i++ {
		answer = append(answer, ".")
	}
	fmt.Println(answer)
	for lifes > 0 {
		colorReset := "\033[0m"
		fmt.Println(string(colorReset))
		var inputAnswer string
		var checkAnswer = false

		fmt.Print("Type a letter : ")
		fmt.Scan(&inputAnswer)

		fmt.Println("----------------------")

		for i := range word {
			if inputAnswer == word[i] {
				checkAnswer = true
				answer[i] = inputAnswer
				word[i] = "."
				fmt.Println(answer)
				break
			}
		}
		if checkAnswer {
			score = score + 20
			fmt.Println()
			fmt.Println("Your score is : ", score)
			if checkWin(strings.Join(answer, ""), randomElement) {
				colorGreen := "\033[32m"
				fmt.Println(string(colorGreen))
				fmt.Println("You Win")
				fmt.Println()
				fmt.Println("Your score is : ", score)
				colorReset := "\033[0m"
				fmt.Println(string(colorReset))
				fmt.Println("You want to replay ? (YES / NO)")
				replay()
				break
			}

		} else {
			lifes--
			drawHangman(lifes)
			letterused = append(letterused, inputAnswer)
			if lifes >= 1 {
				fmt.Println("Lettres Used : ", letterused)
				fmt.Println("Attetion ! You still have ", lifes, "lifes ")
			}
		}
	}
}

func replay() bool {
	response := ""
	fmt.Scanln(&response)
	if response == "YES" || response == "yes" {
		return true
	}
	return false
}
func gamover() {
	fmt.Println("GAME OVER !")
}

func main() {

	println("----------------------------")
	println("----------HANGMAN-----------")
	println("----------------------------")
	fmt.Println("<--- Bienvenue dans notre jeu Hangman ! ---> ")
	fmt.Println("****You have 10 lifes !****")

	var menuChoice int
	var difficulty string

	for ok := true; ok; ok = menuChoice != 4 { //Boucle do while
		colorReset := "\033[0m"
		fmt.Println(string(colorReset))
		println("*** Main Menu ***")
		println("Choose difficulty (F / M / D) : ")
		fmt.Scanln(&difficulty)
		println("1- Cars Theme")
		println("2- Food Theme")
		println("3- Country Theme")
		println("4- Exit")
		fmt.Print("Choose element: ")
		fmt.Scan(&menuChoice)

		switch difficulty {

		case "F":
			switch menuChoice {
			case 1:
				fmt.Println("Please enter lettres with uppercase !")
				gamePlay1("cars1")
				break
			case 2:
				fmt.Println("Please enter the letters in French !")
				gamePlay1("food1")
				break
			case 3:
				fmt.Println("Please enter the letters in French !")
				gamePlay1("country1")
				break
			case 4:
				break
			default:
				break
			}

		case "M":
			switch menuChoice {
			case 1:
				fmt.Println("Please enter lettres with uppercase !")
				gamePlay2("cars2")
				break
			case 2:
				fmt.Println("Please enter the letters in French !")
				gamePlay2("food2")
				break
			case 3:
				fmt.Println("Please enter the letters in French !")
				gamePlay2("country2")
				break
			case 4:
				break
			default:
				break
			}

		case "D":
			switch menuChoice {
			case 1:
				fmt.Println("Please enter lettres with uppercase !")
				gamePlay3("cars3")
				break
			case 2:
				fmt.Println("Please enter the letters in French !")
				gamePlay3("food3")
				break
			case 3:
				fmt.Println("Please enter the letters in French !")
				gamePlay3("country3")
				break
			case 4:
				break
			default:
				break
			}

		default:
			break
		}

		if replay() {
			continue
		}
		break
	}
}
