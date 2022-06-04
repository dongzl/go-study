package sort

import (
	"fmt"
	"sort"
	"testing"
)

func TestHeros_sort(t *testing.T) {
	heros := Heros{
		&Hero{"吕布", Tank},
		&Hero{"李白", Assassin},
		&Hero{"妲己", Mage},
		&Hero{"貂蝉", Assassin},
		&Hero{"关羽", Tank},
		&Hero{"诸葛亮", Mage},
	}

	sort.Sort(heros)

	for _, v := range heros {
		fmt.Printf("%+v\n", v)
	}
}

func TestHeros_sort2(t *testing.T) {
	heros := Heros{
		&Hero{"吕布", Tank},
		&Hero{"李白", Assassin},
		&Hero{"妲己", Mage},
		&Hero{"貂蝉", Assassin},
		&Hero{"关羽", Tank},
		&Hero{"诸葛亮", Mage},
	}

	sort.Slice(heros, func(i, j int) bool {
		if heros[i].Kind != heros[j].Kind {
			return heros[i].Kind < heros[j].Kind
		}
		return heros[i].Name < heros[j].Name
	})

	for _, v := range heros {
		fmt.Printf("%+v\n", v)
	}
}

func TestString_sort(t *testing.T) {
	strs := MyStringList{
		"a",
		"b",
		"z",
		"t",
	}

	sort.Sort(strs)

	for _, v := range strs {
		fmt.Printf("%+v\n", v)
	}
}

func TestString_IsSorted(t *testing.T) {
	strs := MyStringList{"a", "b", "z", "t"}
	ret := sort.IsSorted(strs)
	fmt.Printf("%+v\n", ret)

	strs2 := MyStringList{"a", "b", "z"}
	fmt.Printf("%+v\n", sort.StringsAreSorted(strs2))
}
