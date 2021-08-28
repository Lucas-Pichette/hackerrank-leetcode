// Lucas Pichette's LeetCode submission for the problem: Palindrome Number
// 8/28/21

func isPalindrome(x int) bool{
    reversedStr := ""
    normalStr := strconv.Itoa(x)
    for i := len([]rune(normalStr)) - 1; i >= 0; i-- {
        reversedStr += string(normalStr[i])
    }
    if reversedStr == normalStr {
        return true
    }
    return false
}
