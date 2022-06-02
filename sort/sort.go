package sort

type HeroKind int

const (
	None HeroKind = iota
	Tank
	Assassin
	Mage
)

type Hero struct {
	Name string
	Kind HeroKind
}

type Heros []*Hero

func (s Heros) Len() int {
	return len(s)
}

func (s Heros) Less(i, j int) bool {
	if s[i].Kind != s[j].Kind {
		return s[i].Kind < s[j].Kind
	}
	return s[i].Name < s[j].Name
}

func (s Heros) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type MyStringList []string

func (p MyStringList) Len() int           { return len(p) }
func (p MyStringList) Less(i, j int) bool { return p[i] < p[j] }
func (p MyStringList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
