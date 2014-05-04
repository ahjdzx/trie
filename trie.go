package trie

type Link struct {
	value rune
	link *Trie
}

type Trie struct {
	childNodes []Link
	stringEnd bool
}

func findRuneLink(links []Link, value rune) (*Trie, bool) {
	for _, link := range links {
		if link.value == value {
			return link.link, true
		}
	}
	return nil, false
}

func (r *Trie) ExistsOrAdd(s string) bool {
	check := true
	i := r
	for _, runeValue := range s {
		ti, ok := findRuneLink(i.childNodes,runeValue)
		if !ok {
			ti = new(Trie)
			ti.childNodes = make([]Link, 0)
			i.childNodes = append(i.childNodes,Link{value: runeValue, link: ti})
		}
		i = ti
	}
	if !i.stringEnd {
		i.stringEnd = true
		check = false
	}
	return check
}

func NewTrie() *Trie {
	return &Trie{stringEnd: true, childNodes: make([]Link, 0)}
}

