package main

import "fmt"

func BasicJoin(strs []string) string {
	strs = append(strs, string(" "))
	k := strs[0] + strs[4] + strs[1] + strs[4] + strs[2] + strs[4] + strs[3]

	return k
}

func main() {
	toConcat := []string{"Hello!", "How", "are", "you?"}
	fmt.Println(BasicJoin(toConcat))
}
