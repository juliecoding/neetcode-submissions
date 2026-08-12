class Solution {
    /**
     * @param {string} s
     * @return {boolean}
     */
    isValid(str) {
        const complements = {
            "(": ")",
            "{": "}",
            "[": "]"
        }
        const stack = []

        for (const char of str) {
            if (Object.hasOwn(complements, char)) {
                stack.push(char)
            } else {
                const topOfStack = stack.pop()
                if (!topOfStack || char !== complements[topOfStack]) {
                    return false
                }
            }
        }
        return stack.length === 0
    }
}
