package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("write task name")
	input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	s := strings.TrimSuffix(strings.Replace(strings.Replace(strings.ToLower(input), ".", "", 1), " ", "_", 5000), "\n")
	_ = os.Mkdir(fmt.Sprintf("./leetcode/%s", s), 0750)
	file, _ := os.Create(fmt.Sprintf("./leetcode/%s/main.go", s))
	defer file.Close()
	_, _ = file.Write([]byte(fmt.Sprintf("package main\n\n//enable for random structs\nimport _ \"algorithms/base/random\"\n\n//%s\nfunc main(){\n\n}", s)))
}
