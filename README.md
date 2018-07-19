# Go Programming Workshop

> Introductory workshop for the Go Programming language

## Prerequisites
Ensure you've completed the [getting started](https://golang.org/doc/install) and [how to write go code](https://golang.org/doc/code.html) tutorials

## Set up
 * Ensure you've **completed the prerequisites**
 * Ensure you GOPATH is set e.g GOPATH=/Users/YOUR_USERNAME/workspace/go. If not set, it defaults to $HOME/go
 * `mkdir $GOPATH/src/github.com/YOUR_USER_NAME`
 * `cd $GOPATH/src/github.com/YOUR_USER_NAME`
 * `git clone git@github.com:bbc/go-workshop.git`
 * If you're using `Visual Studio code` you may wish to install the offical Go plugin.

## Workshop
This is a **self paced** workshop which allows developers to learn the fundamentals 
of the [Go Programming language](https://golang.org/) by solving coding challenges 
and writing small applications. 

**This is not a tutorial or lecture based workshop**. However, the workshop will provide
references to any required resources. 

## Instructions 
    * Read the questions carefully. 
    * Some exercises have tests but not all. Typically you don't need to worry about the tests unless instructed.  
    * DO NOT change the tests. Report any issues.
    * DO NOT cheat the tests (pointless) i.e return hardcoded values to make them pass. 
    * Use referenced resources for assistance.
    * Answer the questions in order (optional but recommended)
    * You can skip questions
    * Work individually (recommended) unless instructed.
 
## Useful commands
 * `go test ./...` - runs all tests recursively
 * `go test -v ./...` - as above but with verbose output
 * `go test FULLY_QUALIFIED_PACKAGE_NAME --run TEST_NAME` - executes the `TEST_NAME` test in `FULLY_QUALIFIED_PACKAGE_NAME`
 * `go run prog.go` - executes `prog.go` (compile and execute `prog.go` in one step). Note - `prog.go` has a main function.

### Examples:
Execute any tests matching `^TestIsUnique$` in the `github.com/bbc/go-workshop/workshop/slices` package:
```
go test github.com/bbc/go-workshop/workshop/slices -run ^TestIsUnique$
```

Execute the reverse tests from project root. NOTE - for compilation, the source files must be supplied:
```
go test slices/reverse*.go
```

## Resources
* [Geting started with Go](https://golang.org/doc/install)
* [How to write Go code](https://golang.org/doc/code.html)
* [Go documentation](https://golang.org/doc/)
* [Go packages](https://golang.org/pkg/)
* [Go playground](https://play.golang.org/)

## Have you read the instructions?
If so [_start workshop_](workshop/questions.md)


