# LearningGo

Learning materials covered for Go.

Common commands for my remembering:

```bash
go mod init x.go
```

This command initializes a new Go module in the current directory. It creates a `go.mod` file that defines the module's path and its dependencies. The `x.go` argument is typically the name of the module, which can be a package path or a simple name.

```bash
go run x.go
```

This command compiles and runs the Go program defined in the `x.go` file. It is a convenient way to execute a Go program without needing to build an executable first. The `go run` command will automatically handle the compilation and execution process for you.

```bash
go build x.go
```

This command compiles the Go source code in the `x.go` file and produces an executable binary. The resulting binary will have the same name as the source file (without the `.go` extension) by default. You can specify a different output name using the `-o` flag, like this: `go build -o myprogram x.go`. After running this command, you can execute the generated binary directly from the command line.

```bash
./x
```

This command is used to execute the compiled binary named `x`.

```bash
go mod tidy
```

This command is used to clean up the `go.mod` file by removing any dependencies that are no longer needed and adding any missing dependencies. It ensures that the `go.mod` file accurately reflects the dependencies required by the project. Running `go mod tidy` can help keep your module's dependencies organized and up to date.

:= initialises a variable and infers the type of it based on its value. really cool.
can also do `var x string` etc.

useful: `go mod edit -replace example.com/greetings=../greetings` to replace a module with a local path for testing purposes. This allows you to use a local version of the `greetings` module instead of fetching it from a remote repository.

A slice is like an array, except that its size changes dynamically as you add and remove items.

`go test` read module's tests. has the `-v` flag for verbose output.

`go build` build correct dir.

`go list -f '{{.Target}}'` to get the path of the built binary.

`go install` once we have it in Path the dir of the module?

and then can just run the module by calling the name of the package. export it to the .bash file to have it on permanently.

`go work init ./x` to initialize a workspace with the `x` module. This allows you to manage multiple modules together in a single workspace, making it easier to work on related projects without needing to manage separate `go.mod` files for each module.
