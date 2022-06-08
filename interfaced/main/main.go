package main

import (
	"fmt"
	"goworkspase/interfaced"
	"sort"
)

func main() {
	var heroes interfaced.HeroByAge
	heroes = append(heroes, interfaced.Hero{"zs", 20})
	heroes = append(heroes, interfaced.Hero{"ls", 15})
	heroes = append(heroes, interfaced.Hero{"ww", 18})
	sort.Sort(heroes)
	fmt.Println(heroes)
}
