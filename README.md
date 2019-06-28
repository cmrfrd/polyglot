# Polyglot

<img align="left" width="169" height="400" src="./pics/polyglot.png">

[![CircleCI](https://circleci.com/gh/cmrfrd/polyglot.svg?style=svg)](https://circleci.com/gh/cmrfrd/polyglot)

This repo is a collection of solutions for common programming problems in multiple programming languages.

Learning a new programming language is hard. Just when you think you are "fluent" there is some caveat you might have never thought of or seen. It seems it is a necessity that in order to expand your mind as a programmer you must proficient in multiple languages. Without doing so might cause you to be unknowingly pigeonholed.

> Everyone has a vested interest in the languages they already know, because (a) it is a lot of work to learn a new language, and (b) programming languages dictate how you think about programs, so it is hard even to conceive of a language more powerful than whatever you're used to.  -- Paul Graham

|

|


## Requirements

- [Docker](https://docs.docker.com/install/) version 18.09.1

## Running this project

To run all problems in all languages run

```
$ make build run
```

To run a particular problem in a particular language, fill in the following paramaters

``` shell
$ PROBLEM=<problem name> LANGUAGE=<language name> make build run
```

## Examples

Run `hello_world` in python

``` shell
$ PROBLEM=hello_world LANGUAGE=python make build run
```

Run all `hello_world` language solutions

```
$ PROBLEM=hello_world make build run
```

## Supported languages

- Python
- Scala
- Javascript
- Golang
- Common Lisp
