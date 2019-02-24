# Go

Golang is a static-typed language developed by Google. Let's learn this language to add more things into our arsenal. Golang can be useful for web service. Need to research further if it is also useful for scripting and data science.

## Tutorial on how to run go in your machine
1. Install Go binary from Go offical website: [Golang DLs](https://golang.org/dl/)
2. Ensure Go is installed
    * **Linux**: `/usr/local/go`
    * **Mac**: `/usr/local/go`
3. *something something about GOPATH* (I didn't set $GOPATH env variable but it still works fine)
4. create `hello` dir, enter it:
    ```
    mkdir $HOME/Codes/go/src/hello
    cd $HOME/Codes/go/src/hello
    ```
5. and create simple hello world code
    ```
    package main

    import "fmt"

    func main() {
        fmt.Printf("hello, world\n")
    }
    ```
6. `go build` in terminal
7. Assuming we are in `hello` dir: `./hello` in terminal to run code.
8. **DONE**