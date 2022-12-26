package graph_grock_the_algo

import (
    "fmt"
)

func main() {
    graph := make(map[string][]string)
    graph["you"] = []string{"alice", "bob", "claire"}
    graph["bob"] = []string{"anuj", "peggy"}
    graph["alice"] = []string{"peggy"}
    graph["claire"] = []string{"thom", "jonny"}
    graph["anuj"] = []string{}
    graph["peggy"] = []string{}
    graph["thom"] = []string{}
    graph["jonny"] = []string{}

    lookFor := "peggy1"

    qu := make([]string, 0)
    qu = append(qu, graph["you"]...)
    for len(qu) > 0 {
        currMan := qu[0]
        qu = qu[1:]
        if currMan == lookFor {
            fmt.Println("Finded!")
            return
        } else {
            qu = append(qu, graph[currMan]...)
        }
    }

    fmt.Println("Not finded(")
}
