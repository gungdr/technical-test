![Go](https://github.com/gungdr/technical-test/workflows/Go/badge.svg)
# Technical Test

This is a repository to hold technical test requirments

## Requirments

This project is written in `go v1.15`

## Installation

I use Makefile to run script for the project
### Dependencies

```bash
make init
```
this script will run `go mod vendor`

```bash
make gen
```
this script generate .go files from .proto

### Runing the project

```bash
make run
```
this script will run `go run main.go`

### Test

```bash
make test
```
this script will run `go test ./...`

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)