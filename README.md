# abcsort

[![Build Status](https://travis-ci.org/icza/abcsort.svg?branch=master)](https://travis-ci.org/icza/abcsort)
[![GoDoc](https://godoc.org/github.com/icza/abcsort?status.svg)](https://godoc.org/github.com/icza/abcsort)
[![Go Report Card](https://goreportcard.com/badge/github.com/icza/abcsort)](https://goreportcard.com/report/github.com/icza/abcsort)
[![codecov](https://codecov.io/gh/icza/abcsort/branch/master/graph/badge.svg)](https://codecov.io/gh/icza/abcsort)

Go string sorting library that uses a custom, user-defined alphabet.

`abcsort` provides the essence of sorting: the implementation of a `less()` function
required by the standard lib's sort package.

Implementation does not convert the input strings into byte or rune slices, so
performance is rather good.
