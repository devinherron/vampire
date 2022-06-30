package txtreader

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

func ReadFile(filename string) {
	f, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	s := bufio.NewScanner(f)
	for s.Scan() {
		re := regexp.MustCompile(`^[0-9]{1,2}[a-c]$`)
		promptNumber := re.FindString(s.Text())
		if promptNumber != "" {
			var promptText string = ""
			for {
				s.Scan()
				if s.Text() != "" {
					promptText += s.Text()
				} else {
					break
				}
			}
			fmt.Printf("%s: %s\n", promptNumber, promptText)
		}
		if err != nil {
			log.Fatal(err)
		}
	}
	err = s.Err()
	if err != nil {
		log.Fatal(err)
	}
}
