package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

var isUrl bool

func init() {
	flag.BoolVar(&isUrl, "url", false, "use for getting info from leetcode url")
}

func main() {
	if !isUrl {
		fmt.Println("write task name")

		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')

		s := strings.TrimSuffix(strings.Replace(strings.Replace(strings.ToLower(input), ".", "", 1), " ", "_", 5000), "\n")

		_ = os.Mkdir(fmt.Sprintf("./leetcode/%s", s), 0750)

		file, _ := os.Create(fmt.Sprintf("./leetcode/%s/main.go", s))

		defer func() {
			_ = file.Close()
		}()

		_, _ = file.Write([]byte(fmt.Sprintf("package main\n\n//enable for random structs\nimport _ \"algorithms/base/random\"\n\n//%s\nfunc main(){\n\n}", s)))
	} else {
		getInfoFromLeetcodeURL()
	}
}

func getInfoFromLeetcodeURL() {
	fmt.Println("flaged")
}
