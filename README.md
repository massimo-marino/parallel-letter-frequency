# Parallel Letter Frequency

**Files:** *occurrences.go, occurrences_test.go*

Count the frequency of letters in texts using parallel computation.

Parallelism is about doing things in parallel that can also be done sequentially.  
A common example is counting the frequency of letters.  
Create a function that returns the total frequency of each letter in a text *and* that employs parallelism.  

The implementation depends on the following packages:

- **btree**: from ``github.com/emirpasic/gods/trees/btree`` ([https://github.com/emirpasic/gods](https://github.com/emirpasic/gods) ); you have to ``go get github.com/emirpasic/gods``

- **alphabet**: see [https://github.com/massimo-marino/alphabet.git](https://github.com/massimo-marino/alphabet.git) 
