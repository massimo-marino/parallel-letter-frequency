package occurrences

import (
	"fmt"
	"testing"
)

const checkMark = "\u2714"
const ballotX = "\u2718"

func init() {
	fmt.Println("*** Starting Tests for occurrences ***")
	fmt.Println()
}

var ft = func(t *testing.T, alphabetString string, text string, totalOccurrencesExpected counter) {

	_, totalOccurrences := Fco(alphabetString, text)

	if totalOccurrencesExpected == totalOccurrences {
		t.Log("Total occurrences:", totalOccurrences, "Expected: ", totalOccurrencesExpected, ": OK", checkMark)
	} else {
		t.Error("Total occurrences:", totalOccurrences, "Expected: ", totalOccurrencesExpected, ": NOT OK", ballotX)
	}
}

func TestO(t *testing.T) {
	alphabetString := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	text := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

	ft(t, alphabetString, text, counter(len(text)))
}

func Test1(t *testing.T) {
	alphabetString := "9876543210"
	text := "0123456789"

	ft(t, alphabetString, text, counter(len(text)))
}

func Test2(t *testing.T) {
	alphabetString := ""
	text := ""

	ft(t, alphabetString, text, counter(len(text)))
}

func Test3(t *testing.T) {
	alphabetString := "abcdefghijklmnopqrstuvwxyz"
	text := ""

	ft(t, alphabetString, text, counter(len(text)))
}

func Test4(t *testing.T) {
	alphabetString := "abcdefghijklmnopqrstuvwxyz"
	text := "9876543210"

	ft(t, alphabetString, text, counter(0))
}

func Test5(t *testing.T) {
	exampleOfUse()
}

func TestLastTest(t *testing.T) {
	fmt.Println()
	fmt.Println("*** Ended Tests for occurrences ***")
	fmt.Println()

}
