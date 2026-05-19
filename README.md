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
