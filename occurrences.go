// occurrences.go
// an example of counting letter occurrences in a text using parallelism
// used a btree just for learning purposes
package occurrences

import (
	"alphabet"
	"fmt"
	"sort"
	"strings"
	"sync"
	// you have to 'go get github.com/emirpasic/gods'
	"github.com/emirpasic/gods/trees/btree"
)

var wg sync.WaitGroup

type workerFun func(in string, c string, inch chan data, reschan resultch)

// data structure exchanged with workers
type data struct {
	msg string
}

type counter uint64
type resultch chan counter
type result map[string]counter

// elements of the btree
type elem struct {
	c       string
	inchan  chan data
	reschan resultch
	wf      workerFun
}

// the worker function runs as a goroutine
func worker(in string, c string, inch chan data, reschan resultch) {
	defer wg.Done()
	var counter counter = 0

	// loop forever until the channel is closed: the loop is terminated ONLY
	// when the channel is closed, otherwise it deadlocks
	for inData := range inch {
		if inData.msg == c {
			counter++
			// fmt.Println(in, "Received:", inData, " - Received so far:", counter)
		}
	}

	// send back the result and close the sending channel
	reschan <- counter
	close(reschan)

	if counter != 0 {
		fmt.Println(in, "Ended: Received:", counter, "of", in)
	}
}

// count occurrences of alphabet letters in a text using parallelism
func CountOccurrencesInText(a alphabet.Alphabet, text string) result {

	fmt.Printf("Alphabet: '%s'\nCounting on text: '%s'\n\n", a, text)

	// empty tree, keys are of type string
	tree := btree.NewWithStringComparator(3)

	// start a worker for every element of the alphabet
	for _, c := range a {
		cs := string(c)
		inchan := make(chan data, 10)
		reschan := make(resultch)
		tree.Put(cs, elem{cs, inchan, reschan, worker})
		wg.Add(1)
		go worker(cs, cs, inchan, reschan)
	}

	// send characters to their workers
	for _, c := range text {
		if v, found := tree.Get(string(c)); found {
			v.(elem).inchan <- data{string(c)}
		} else {
			fmt.Printf("Character '%s' not found in alphabet: Not counted\n", string(c))
		}
	}

	// prepare the map with results
	r := make(result, len(a))

	// iterate through the tree and: 1. close all channels; 2. receive the results
	it := tree.Iterator()
	for it.Next() {
		//index, value := it.Key(), it.Value()

		// close the channels used for sending
		close(it.Value().(elem).inchan)

		// receive the results
		for s := range it.Value().(elem).reschan {
			r[it.Key().(string)] = s
		}
	}

	wg.Wait()

	return r
}

var Fco = func(alphabetString string, text string) (result, counter) {
	// create the alphabet
	alphabet := alphabet.MakeNewAlphabet(alphabetString)
	// count
	r := CountOccurrencesInText(alphabet, text)
	// compute the total number of occurrences that were found in the alphabet
	totalOccurrences := counter(0)
	for _, v := range r {
		totalOccurrences = totalOccurrences + v
	}

	return r, totalOccurrences
}

func exampleOfUseSimple() {
	alphabetString := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	text := "How long has this been going on?"

	alphabet := alphabet.MakeNewAlphabet(alphabetString)
	r := CountOccurrencesInText(alphabet, text)

	fmt.Println()
	fmt.Println("exampleOfUse_Simple: Results:")
	totalOccurrences := counter(0)
	for k, v := range r {
		totalOccurrences = totalOccurrences + v
		fmt.Println("exampleOfUse_Simple:", k, "->", v)
	}
	fmt.Println("exampleOfUse_Simple: Total occurrences:", totalOccurrences)
}

func exampleOfUse() {
	alphabetString := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!, "
	text := "Hello, World!"

	results, totalOccurrences := Fco(alphabetString, text)

	fmt.Println()
	fmt.Println("ExampleOfUse: Results:")

	// sort the alphabet string, if not sorted
	keys := strings.Split(alphabetString, "")
	sort.Strings(keys)
	// print results sorted since results as returned from Fco() is not sorted
	for _, k := range keys {
		v := results[k]
		fmt.Println("ExampleOfUse:", k, "->", v)
	}

	fmt.Println("ExampleOfUse: Total occurrences:", totalOccurrences)
}
