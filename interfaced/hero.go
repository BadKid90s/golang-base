package interfaced

type Hero struct {
	Name string
	Age  int
}
type HeroByAge []Hero

func (a HeroByAge) Len() int {
	return len(a)
}
func (a HeroByAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a HeroByAge) Less(i, j int) bool {
	return a[i].Age < a[j].Age
}

