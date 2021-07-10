package vampire

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"
	"regexp"
	"strings"
)

func readSourceFile(file multipart.File) {
	bio := bufio.NewReader(os.Stdin)
	for {
		line, err := bio.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		sline := strings.TrimRight(string(line), "\n")
		re := regexp.MustCompile(`^[0-9]{1,2}[a-c]$`)
		promptNumber := re.FindString(sline)
		if promptNumber != "" {
			var promptText string = ""
			for {
				line, err := bio.ReadBytes('\n')
				if err == io.EOF {
					break
				}
				if err != nil {
					panic(err)
				}
				sline := strings.TrimRight(string(line), "\n")

				if sline != "" {
					promptText += sline
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
}

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
