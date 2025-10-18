package main

import (
    "fmt"
    "strings"
    "unicode"

    "golang.org/x/text/cases"
    "golang.org/x/text/language"
)

// Constants for configuration
const (
    defaultSplitLimit = 3
    repeatCount       = 3
    unlimitedSplit   = -1
)

// demonstrateBasicComparisons shows basic string comparison operations
func demonstrateBasicComparisons(s1, s2 string) {
    fmt.Println("\n=== Basic String Comparisons ===")
    fmt.Printf("Compare: %d\n", strings.Compare(s2, s1))
    fmt.Printf("Contains 'Exciting': %t\n", strings.Contains(s1, "Exciting"))
    fmt.Printf("ContainsAny 'F65432': %t\n", strings.ContainsAny(s1, "F65432"))
}

// demonstrateStringSplitting shows different ways to split strings
func demonstrateStringSplitting(s string) {
    fmt.Println("\n=== String Splitting ===")
    words := strings.Fields(s)
    for i, word := range words {
        fmt.Printf("Word %d: %s\n", i, word)
    }
}

// demonstrateStringJoining shows string joining operations
func demonstrateStringJoining(elements []string) string {
    fmt.Println("\n=== String Joining ===")
    return strings.Join(elements, " <--> ")
}

// demonstrateStringSearching shows various string search operations
func demonstrateStringSearching(s string) {
    fmt.Println("\n=== String Searching ===")
    fmt.Printf("Index of 'Fun': %d\n", strings.Index(s, "Fun"))
    fmt.Printf("LastIndex of 'Fun': %d\n", strings.LastIndex(s, "Fun"))
    fmt.Printf("IndexAny of 'is': %d\n", strings.IndexAny(s, "is"))
    fmt.Printf("LastIndexAny of 'is': %d\n", strings.LastIndexAny(s, "is"))
}

// demonstrateCaseOperations shows string case transformation
func demonstrateCaseOperations(s string) {
    fmt.Println("\n=== Case Operations ===")
    fmt.Printf("Original: %s", s)
    fmt.Printf("Upper: %s", strings.ToUpper(s))
    fmt.Printf("Lower: %s", strings.ToLower(s))
    
    englishCaser := cases.Title(language.English)
    fmt.Printf("Title (English): %s", englishCaser.String(s))
}

// demonstrateUnicodeOperations shows operations with Unicode characters
func demonstrateUnicodeOperations(s string) {
    fmt.Println("\n=== Unicode Operations ===")
    isChineseCharacter := func(r rune) bool {
        return unicode.Is(unicode.Han, r)
    }
    
    fmt.Printf("First Chinese character at: %d\n", 
        strings.IndexFunc("Hello, 你好", isChineseCharacter))
}

func main() {
    // Test strings
    stringVar1 := "Learning Go is Fun and Exciting\n"
    stringVar2 := "Learning"
    
    // Demonstrate different string operations
    demonstrateBasicComparisons(stringVar1, stringVar2)
    demonstrateStringSplitting(stringVar1)
    
    mySlice := []string{"Here", "Iam", "Learning", "Go", "Language"}
    joined := demonstrateStringJoining(mySlice)
    fmt.Printf("Joined result: %s\n", joined)
    
    demonstrateStringSearching(stringVar1)
    demonstrateCaseOperations(stringVar1)
    
    // Special character handling
    specialString := "São Paulo is a large city"
    demonstrateUnicodeOperations(specialString)
    
    // Example of string replacement
    fmt.Println("\n=== String Replacement ===")
    fmt.Println(strings.ReplaceAll(stringVar1, "Fun", "Exciting"))
}
