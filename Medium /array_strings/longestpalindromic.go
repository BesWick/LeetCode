//Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.

/**
Example:

Input: "babad"
Output: "bab"
Note: "aba" is also a valid answer.

Input: "cbbd"
Output: "bb"


Solution: 0(n^2)
*/

func longestPalindrome(s string) string {
    n := len(s)
    pali := make([][]bool, n)
    for i:= 0; i<n; i++{
        pali[i] = make([]bool, n)
        pali[i][i] = true
    }
    
    resultStr:= ""
    if len(s) <=1 {return s}
    
    for i:= n-1; i>=0; i--{
        for j:= i+1; j < n; j++{
            if s[i] == s[j] {
                if j - i == 1{ //fiybd a pair
                    pali[i][j] = true
                }else{
                    pali[i][j] = pali[i+1][j-1]
                }
                if pali[i][j] && (j-i +1) >= len(resultStr){
                    resultStr = s[i:j+1]
                }
            }
        }
    }
    if len(resultStr) == 0{ //if there's no pali substring
        return string(s[0])
    }
    
    return resultStr
    
}