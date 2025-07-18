/*
Given a string text, you want to use the characters of text to form as many instances of the word "balloon" as possible.

You can use each character in text at most once. Return the maximum number of instances that can be formed.

Example 1:

Input: text = "nlaebolko"
Output: 1

Example 2:

Input: text = "loonbalxballpoon"
Output: 2

Example 3:

Input: text = "leetcode"
Output: 0

Constraints:

	1 <= text.length <= 104
	text consists of lower case English letters only.
*/
package main

import (
	"fmt"
)

func maxNumberOfBalloons(text string) int {

	balloonsText := "balloon"
	countBalloons := map[rune]int{}
	for _, value := range balloonsText {
		countBalloons[value]++
	}

	countLetters := map[rune]int{}
	for _, value := range text {
		countLetters[value]++
	}
	// fmt.Println(countLetters)

	result := 1000000

	for key, value := range countBalloons {
		if countLetters[key] == 0 {
			return 0
		}

		if countLetters[key]/value < result {
			result = countLetters[key] / value
		}
	}

	return result

}

func main() {
	text := "loonbalxballpoon"
	fmt.Println(maxNumberOfBalloons(text))
}
