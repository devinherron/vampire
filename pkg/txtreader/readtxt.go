package txtreader

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

var (
	re = regexp.MustCompile(`^[0-9]{1,2}[a-c]$`)
)

func ReadFile(filename string) [][]string {
	const numPrompts int = 81

	f, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	prompts := make([][]string, numPrompts)

	s := bufio.NewScanner(f)
	start := false
	for s.Scan() {
		if !start && s.Text() == "Prompts" {
			start = true
			continue
		} else if s.Text() == "Appendix One" {
			break
		}

		prompt, entry := GetPrompt(s.Text())
		if prompt != 0 {

			var text string = ""
			for {
				s.Scan()
				newPrompt, newEntry := GetPrompt(s.Text())

				if newPrompt != 0 && newPrompt != prompt {
					prompt = newPrompt
					entry = newEntry
					text = s.Text()
				} else if s.Text() == "________________" {
					break
				} else if s.Text() != "" {
					text += s.Text()
				} else {
					break
				}
			}

			if len(prompts[prompt]) == 0 {
				prompts[prompt] = make([]string, 3)
			}

			prompts[prompt][entry] = text
		}

		if err != nil {
			log.Fatal(err)
		}
	}
	err = s.Err()
	if err != nil {
		log.Fatal(err)
	}

	return prompts
}

func GetPrompt(text string) (int, int) {
	prompt := re.FindString(text)
	if prompt == "" {
		return 0, 0
	}

	num, err := strconv.Atoi(prompt[:len(prompt)-1])

	if err != nil {
		log.Fatal(err)
	}

	letter := prompt[len(prompt)-1:]
	var entry int
	switch letter {
	case "a":
		entry = 0
	case "b":
		entry = 1
	case "c":
		entry = 2
	default:
		log.Fatal("Invalid prompt entry.")
	}

	return num, entry
}
