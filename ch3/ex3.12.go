package main

import (
    "fmt"
)

func isAnagram(word1, word2 string) bool {
    wordMap := make(map[string]int)
    for i:=0; i<len(word1); i++ {
        wordMap[word1[i:i+1]]++
    }
    for i:=0; i<len(word2); i++ {
        wordMap[word2[i:i+1]]--
    }

    for _, v := range wordMap {
        if v !=0 {
            return false
        }
    }

    return true
}

func main() {
    fmt.Printf("%s and %s is anagram? %t\n", "live", "evil", isAnagram("live", "evil"))
    fmt.Printf("%s and %s is anagram? %t\n", "north", "thorn", isAnagram("north", "thorn"))
    fmt.Printf("%s and %s is anagram? %t\n", "north", "Thorn", isAnagram("north", "Thorn"))
    fmt.Printf("%s and %s is anagram? %t\n", "wet", "ttew", isAnagram("wet", "ttew"))
}
