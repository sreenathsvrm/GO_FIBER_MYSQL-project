# Clean Architecture using Golang with Fiber framework

## 

-Fiber is a fast, efficient, and developer-friendly web framework for Go. It provides a robust API and is designed to have low memory allocation and high performance. Fiber is an ideal choice for projects that require both performance and productivity.

-GORM is a fantastic ORM library for Go, designed to be developer-friendly and powerful. It allows you to interact with databases using Go struct and provides a simple and intuitive API for database operations.

- [Wire](https://github.com/google/wire) is a code generation tool that automates connecting components using dependency injection.
- [Viper](https://github.com/spf13/viper) is a complete configuration solution for Go applications including 12-Factor apps. It is designed to work within an application, and can handle all types of configuration needs and formats.

## Using `go-FIBER-clean-arch` project

To use `go-fiber-clean-arch` project, follow these steps:
# Install dependencies
make deps

# Generate wire_gen.go for dependency injection
# Please make sure you are export the env for GOPATH
make wire

# Run the project in Development Mode
make run
```

Additional commands:

```bash
➔ make help
build                          Compile the code, build Executable File
run                            Start application
test                           Run tests
test-coverage                  Run tests and generate coverage file
deps                           Install dependencies
deps-cleancache                Clear cache in Go module
wire                           Generate wire_gen.go
help                           Display this help screen
```


