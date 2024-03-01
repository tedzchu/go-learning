# Go Learnings

Just me trying to learn a new language :)

## Bread and butter

- `go mod init [package]` - create module
- `go mod tidy` - pull deps in package
- `go get .` - update deps (when to use over mod tidy?)
- `go mod edit -replace [package]=[local_package]` - fix local deps

- `:=` means assigning
  - But also `in` in a `for in` loop?
- Return types can be folded i.e. `(string, error)`
- `log.Fatal(err)` exits the program
- What are possible log flags?
- `make` initializes slice, map, or chan
  - I think a slice is an array
- `map` type is declared by key and value type, i.e. `map[key-type]val-type`

- `_test.go` will get picked up with `go test`

## Compiling and building

- `go build` - creates executable you can run with `./[module]` where `module` is folder name
- Install path is `go list -f '{{.Target}}`, updated PATH already to go bin (GOBIN)
- `go install` puts the exe in there, makes it accessible by name (so run `hello` in CLI just works)

## Workspaces

- `go work init ./[module_dir]` makes a `go.work` file that orchestrates the workspace, making modules accessible anywhere in the workspace
- `go work use ./[module_dir]` to add new stuff to the workspace

## RDBMS

- Pretty straightforward here, the MySQL driver is quite nice
- `defer rows.Close()` ? defers execution of a function until the surrounding function returns
