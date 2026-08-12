func isValid(str string) bool {
    stack := []rune{}
    complements := map[rune]rune{
        '(': ')',
        '{': '}',
        '[': ']',
    }

    for _, char := range str {
        if _, ok := complements[char]; ok {
            stack = append(stack, char)
        } else {
            if len(stack) < 1 {
                return false
            }
            target := stack[len(stack) - 1]
            if complements[target] != char {
                return false
            }
            stack = stack[0:len(stack) - 1]
        }
    }
    return len(stack) == 0
}
