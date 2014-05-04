trie
====

Golang trie module 

A trie implementation in Golang that you can use to maintain a list of used strings.
For example in a web crawler to keep track of urls allready seen.

By using a trie implementation the list uses a low memory footprint.

        r := NewTrie()
        for i := 1; i < 1000000; i++ {
                r.ExistsOrAdd(strconv.Itoa(i))
        }
        if ! r.ExistsOrAdd("666666") {
                t.Error("ExistsOrAdd: Add failed!")
        } else
        {
                t.Log("ExistsOrAdd: Succesful")
        }

