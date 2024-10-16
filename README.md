# Tetris Optimizer

## Table of Content
- [Description](#description)
- [Features](#features)
- [FileFormat](#file-format)
- [Usage](#usage)
- [Licence](#license)
- [Author](#author)

### Description

`Tetris Optimizer` a simple Tetris-like tetromino puzzle solver in Go. It includes functionality for reading tetromino data from a file, validating the tetrominoes, and solving a given board by placing the tetrominoes.

### Features
- Read Tetrominoes from a File: Reads a list of tetrominoes from a file, where tetromino shapes are represented with # and empty spaces with ..
- Validate Tetrominoes: Ensures that the tetrominoes meet the necessary conditions (valid shapes, sizes, and connections).
- Solver: Attempts to place all the tetrominoes on a board without overlap.

### File Format

The tetrominoes file should contain tetrominoes where each block is represented as # and empty spaces are represented as .. Each tetromino is 4 lines tall, and the tetrominoes are separated by an empty line.
Example File:

bash

```
####
.#..
.#..
....

..##
##..
....
....

#...
###.
....
....

```

### In this example:

- The first tetromino is a square with all blocks (#), represented as A.
- The second tetromino is an L shape, represented as B.
- The third tetromino is a right-angle (T) shape, represented as C.

### Usage

``` 
go run main.go

Then Pass a file containing tetrominoes as an argument:
    file.txt
```
Testing

Unit tests are included to test the file reading, validation, and solving functions. You can run the tests using Go's testing package:
```
bash

go test -v
```
### License

This project is licensed under the MIT License.

### Author

- [Malika](https://learn.zone01kisumu.ke/git/masman)