# Installation and setup 

If you have not installed and completed setup Go (workspace etc), please read [how to get started](https://golang.org/doc/install)

⚠️ Before preceding, ensure you can `build` and `run` code from the getting started guide ⚠️<br>

# Fundamentals

* _Refactor [ShortVars](fundamentals/short.go) to use `short variable declarations`. The test for this should already pass._

* _Refactor [Constants](fundamentals/constants.go) to extract the variable declarations as constants. The test for this should already pass._ 

* _Extract the `fToC` and `convert` functions from [ShortVars](fundamentals/short.go) to a new package called `converter` (within the `fundamentals` package). Ensure `fToC` is private but make the `convert` function public (exported). Refactor any dependant code to use the new converter package_ 

* _Complete the implementation of [Fib](fundamentals/fib.go) to return the nth Fibonacci number. Your final implementation should use `tuple assignments`._ 

* _Complete the implementation of [BigFloat](fundamentals/big_float.go) to create and return a pointer to the largest (max) float64. Do not use the `new` operator. You'll need to uncomment the [tests](fundamentals/big_float_test.go)_

* _Refactor [BigFloat](fundamentals/big_float.go) to create the float64 pointer using `new`._

* _Complete the implementation of [Swap](fundamentals/swap.go) to swap the values of two given integers. The values must be swapped in place. You'll need to uncomment the [tests](fundamentals/swap_test.go)__  

* _Complete the implementation of [Basename](fundamentals/basename.go) to return the file or directory portion of a pathname. An optional `suffix` parameter should be specified to remove the a file suffix (if matched).<br> For example:
<br>`Basename("a/b/c.go", ".go") // c` <br>`Basename("a/b/c.go") // c.go`<br><br>You'll need to uncomment the [tests](fundamentals/basename_test.go)_

* _Complete the implementation of [isSameColour](fundamentals/chess.go) to determine if two locations (cells) on a chessboard are the same colour.<br> For example:
<br>`isSameColour("a1", "c3") // true` <br>`isSameColour("a1", "h3") // false`<br><br>See grid chess board [here](https://en.wikipedia.org/wiki/Grid_chess) to help visualise cells_

# Composite Data types

## Arrays 

⚠️ 🔥 Slices are not arrays. Arrays are fixed size. Do not use slices for the exercises. 🔥 ⚠️<br>

* _Complete the implementation of [CreateArray](arrays/create_array.go) to return an array of 1..10 integers_

* _Complete the implementation of [ReverseWithPointers](arrays/reverse_with_pointers_test.go) that reverses an array of string using pointers. NOTE - you'll need to correct function signiture of the [test](arrays/reverse_with_pointers_test.go) to accept a pointer to an array of string_

* _Complete the implementation of [FirstNonRepeated](arrays/non_repeated.go) to return the first non-repeated character. **You MUST USE AN ARRAY (not a slice or map)**<br>
     For example: <br>`FirstNonRepeated("aabbcddeee") // c` <br> `FirstNonRepeated("aabbcc") // -1`_

## Slices 

⚠️ 🔥 Don't use standard libraries for these questions 🔥 ⚠️<br>

* _Complete the implementation of [StrToSlice](slices/str_to_slice.go) which accepts a string and returns a slice of the strings characters_<br>
    For example:<br> `StrToSlice("filehound")// ['f','i','l','e','h','o','u','n','d']` 

* _Complete the implementation of [Reverse](slices/reverse.go) to reverse a given slice in place_

* _Complete the implementation of [Compare](slices/compare.go) to determine if two slices are equal_

* _Complete the implementation of [IsUnique](slices/unique.go) to determine if a given string contains unique characters_<br>
     For example: <br>`IsUnique("abc")// true` <br> `IsUnique("abbc")// false` 

* _Complete the implementation of [LeftRotate](slices/left_rotate.go) to rotate elements of a slice left by n positions._<br>
    For example: <br>`[1, 2, 3] becomes [2, 3, 1] where n = 1`

* _Complete the implementation of [IsAnagram](slices/anagram.go) to determine if two given strings are an anagram. Spaces should be excluded and the function should be case insensitive._<br>
    For example: <br>`IsAnagram("anagram", "naga ram") // true` 

## Maps 

* _Complete the implementation of [MakeMap](maps/make_map.go) to return a map of string to int using the built in `make` function. Each key is the name of a Greek letter and the value is the length of the name. Use the first five names of the greel alphabet: `alpha`, `beta`, `gamma`, `delta`, `epsilon`<br>_ For example: `MakeMap() // {"alpha: 5", "epsilon: 7" ...}`

* _Complete the implementation of [MapLiteral](maps/map_literal.go) that returns the same map as `MakeMap` but refactored to use a `map literal` (not `make`)_

* _Complete the implementation of [IterateMap](maps/iterate_map.go) to call `MakeMap` and return the sum of all the values._

* _Complete the implementation of [DeleteEntries](maps/map_delete.go) to call `MakeMap` and return the sum of all **uniq** values. Delete the duplicate entries (determine by value) from the map. NOTE - the 2nd return value is the map_

* _Complete the implementation of [CompareMap](maps/map_compare.go) to determine if two maps are equal._

* Rewrite `FirstNonRepeated` to use a `map`. Implement your solution [here](maps/not_repeated.go)

# Structs 

* _Complete the implementation of [CreateBundle](structs/create.go) to return a named `struct` type called `Bundle`. The struct should contain the following fields:<br>_

__Bundle__:

| Name | Field |
|----------------------|--------------|
| ID                   | string       |
| ShouldBeRemoved      | boolean      |
| IsExpected           | boolean      |
| Score                | int          |
| Error                | error        |

You'll need to uncomment the [tests](structs/create_test.go)

# Testing and debugging 

* _[Magic9](testing/magic9.go) is a function that always return `9` given an integer between 20 and 99 (inclusive). Write tests (within the testing package) for this function. This implemation of `Magic9` should not be changed._

* The [BestSongsClient](testing/songs_client.go) returns a list of popular songs and performs basic songs searches. It uses a songs `Database access object (DAO)` to retrieve song metadata from the music database. Calls to the songs DAO are very expensive and can be unpredictable. Consequently, invoking DAOs is **NOT suitable for unit tests**.<br><br>Write `unit tests` for the [BestSongsClient](testing/songs_client.go). Ensure the tests avoid these expensive DAO calls.

* Refactor the [BestSongsClient](testing/songs_client.go) tests to setup any test data before each test runs. i.e like `beforeEach` in Node `@before` in Java.  

* _[PadRight](testing/pad.go) pads (with space) a given string `s` to `n` characters. Complete the [unit tests](testing/pad_test.go) by writing a test that exhausively invokes [PadRight](testing/pad.go) with randomly generated test data. Does [PadRight](testing/pad.go) have a bug?<br><br>Standard libs can be used for this test._


* _Benchmark [PadRight](testing/pad.go) by adding a `benchmark` test to the [unit tests](testing/pad_test.go)_

# Functions 

* _Complete the implementation of [IsPalindrome](functions/palindrome.go) to assert if a given string is a [Palindrome](https://en.wikipedia.org/wiki/Palindrome). This function **MUST BE RECURSIVE**.Spaces should be excluded and the function should be case insensitive._<br>
    For example: <br>`IsPalindrome("Anna") // true`<br>`IsPalindrome("jonty") // false` 

* _Write a function called [Sum](functions/sum.go) that recursively returns the sum of any given integers. This function should support variadic arguments. **Uncomment** [the sum test](functions/sum_test.go)<br>For example:<br> `Sum(99, 98, 97) // 294`._

* _Complete the implementation of [CovertMobile](functions/mobile.go) to convert a given uk mobile number to its international representation. **If the mobile number is invalid, return and error with the message 'Invalid format for: mobileNumber'**_<br> For example: <br>CovertMobile("07704334455") // "+447704334455"<br> number, err := CovertMobile("bad") // err != nil

* _Complete the implementation of [ReadFile](functions/read_file.go) to return the contents of a file of a given file as a slice. Return an error if the file does not exist. Ensure the file is closed **AFTER** [ReadFile](functions/read_file.go) has finished executing_

* _Complete the implemention of [StrMapper](functions/mapper.go) to return a transformed string using a given `map function` and a `string`. The `map function` (1st arg to `StrMapper`) should accept a `rune` and return a `rune`. The `map function` should return `-1` to exclude a `rune` from the transformed `string`._<br>For example: <br>`StrMapper(toLowerCase, "STR") // str`

    You'll need to uncomment the [tests](functions/mapper_test.go)

* [Mutate](functions/mutate.go) accepts a slice of int and removes any numbers (in place) not divisble by 3 and 5. However, the test is failing. **Fix** the failing [test](functions/mutate_test.go). You can modify the `Mutate` and the test, if required._

# Objects, interfaces and methods 

* _Complete the implemention of [Number](objects/number.go). `Number` should be constructed like `x := Number(10)`, accepting an `int32`, `n`. Write methods `Add`, `Sub`, `Div`, `Mul` to perform for required mathematical operation against `n`. Each method should return a `Number`. You'll need to uncomment the [tests](objects/number_test.go) accordying._

* _Fix the database object [db](objects/db.go). Items are being saved with `.saveAll` but the `get` method is returning `unexpected not found` errors._

* _Write a `Set` interface with the following defination:_

| Name | param types | return type     |desc|
|------|-------------|-----------------|----|
| Add       | string | void            |Add string to set|
| Remove    | string | void            |deletes string from set|
| Has       | string | boolean         |`true` is string is present else `false`|
| unionWith | Set    | Set             |merge sets|
| Values    | -      | []string        |slice of values|
| String    | -      | string  | e.g {1,2,3} |


* _Complete the implemention of a [Set data structure](objects/set_imp.go) called `BasicSet`. The implementation should satisfy the `Set` interface. Create the Set via `NewSet`. You'll need to uncomment the [tests](objects/set_imp_test.go) accordying._

* _Write a function `unique` that removes duplicates from a given slice (of int) using a set. Try switching the set implemetations._

* _Refactor `unique` and the set to accept values of any type._

# Concurrency 

⚠️ 🔥 For simplicity this section does may not have tests 🔥 ⚠️<br>
⚠️ 🔥 Programs have been defined in the `main` package to allow them to be executed using `go run someProg.go`🔥 ⚠️<br>

* _Complete the implementation of [echo](concurrency/echo/echo.go) which echos (copies) data from stdin to stdout. The copying logic should be written in a **separate go routine**. The main function should sleep for 30 seconds then exit without any errors (zero exit status)._

* Refactor [compress](concurrency/compress/compress.go) to compress files concurrency using a `wait group`_

* [Count](concurrency/words/count.go) accepts one or more files and writes the count of each word to stdout. However, there is a serious fault. Find and fix the bug using a `mutex`

* _Complete the implementation of [Pipeline](concurrency/pipeline/pipeline.go) using `unbuffered channels`. The Pipeline accepts any number of strings and processes the strings concurrently using different `stages` (functions). e.g create channel => clean => capitalise().<br>`Pipeline("forname1_surname1", "forename2_surname2") // ["Forename1 Surname1", "Forename2 Surname2"]`.You'll need to uncomment the [tests](concurrency/pipeline/pipeline_test.go) accordying._

* _Write a function called `Request` that accepts a slice of *URL (pointer to URL). `Request` should execute a HTTP GET for each URL in parallel. `Request` should return a slice of responses. ⚠️ 🔥 Be carful. Unlimited parallel http requests can be dangerous 🔥 ⚠️_

* _Refactor `Request` to accept a `limit` argument. The number of parallel requests should not exceed `limit` at any given time._

* _Refactor `Request` to timeout requests via a `timeout` parameter._

* _Implement a solution for the [Dinning philosophers problem](https://en.wikipedia.org/wiki/Dining_philosophers_problem)_

# Real world applications and techniques 

* _Write a function to `unmarshall` (JSON decode) an [episode](./fixtures/episode.json) into a Go data structure._

* _Write a simple web server to start up on port 8081 and read [episode fixture](./fixtures/episode.json), unmarshall it (previous exercise) and serve the response._

* _Write a simple (skeleton) command line app. eg: mimic a unix cmd. It must:_
    * Displays a help message
    * Parses commandline options
    * Parses commandline arguments

# Misc

⚠️ 🔥 Implement all solutions using TDD 🔥 ⚠️<br>

* _Complete the implementation of [Reorder](misc/reorder.go), to determine if a given array of strings can be reordered so that each consecutive string only differs by one character<br>`reorder(["xx", "xy", "yy"]) // true`<br>_

* _Complete the implementation of [PopCount](misc/pop.go) to return the number of set bits in a given integer.<br> For example:<br> `PopCount(128) // 8`_

* Complete the implementation of [Translate](misc/translations/translate.go) to substitute [translation keys](./fixtures/translations.txt)_for the corresponding value in a given str. Ensure all tests pass, including basic error handling.<br> For example:<br> `Translate("#{availability.mon}") // Mon`_

* _Write a version of `Translate` that can be executed from a `node` script._

* _Benchmark the above solution (running from `Node`). How does it compare against native `Node` and `golang` solutions?_

* _Write a function called [Sum](misc/sum.go) to add two integers **WITHOUT** using arithemic operators. **A Test suite has already been written**._

* _Write a function to return the shortest palindrome. You can convert the string to a palindrome by adding characters to the beginning (if required)`_

* _Write a data structure for a `LRUCache`. Implemented `.get(key)` and `set(key, value)` methods. When the data structure reaches capacity, the least recently used item should be invalidated before another item is created in the cache._
