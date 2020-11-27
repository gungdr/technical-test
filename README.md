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


## Usage

```python
import foobar

foobar.pluralize('word') # returns 'words'
foobar.pluralize('goose') # returns 'geese'
foobar.singularize('phenomena') # returns 'phenomenon'
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)