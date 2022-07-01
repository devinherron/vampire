package txtreader

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

func ReadFile(filename string) [][]string {
	f, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	prompts := make([][]string, 81)

	s := bufio.NewScanner(f)
	start := false
	for s.Scan() {
		if !start && s.Text() == "Prompts" {
			start = true
			continue
		} else if s.Text() == "Appendix One" {
			break
		}

		re := regexp.MustCompile(`^[0-9]{1,2}[a-c]$`)
		prompt := re.FindString(s.Text())
		var num int
		if prompt != "" {
			num, err = strconv.Atoi(prompt[:len(prompt)-1])
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

			if len(prompts[num]) == 0 {
				prompts[num] = make([]string, 3)
			}
			var text string = ""
			for {
				s.Scan()
				if s.Text() == "________________" {
					break
				} else if s.Text() != "" {
					text += s.Text()
				} else {
					break
				}
			}

			prompts[num][entry] = text
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
