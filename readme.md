# List [![Build Status](https://travis-ci.org/markelog/release.svg)](https://travis-ci.org/markelog/release) [![GoDoc](https://godoc.org/github.com/markelog/release?status.svg)](https://godoc.org/github.com/markelog/release) [![Go Report Card](https://goreportcard.com/badge/github.com/markelog/release)](https://goreportcard.com/report/github.com/markelog/release)

> Get OS type/name/version

The thing is – golang standard library doesn't have methods/constants which could provide os info :/

```go

// Depending on the system

// "mac"
release.Type()

// "El Capitan"
release.Name()

// "10.11.6"
release.Version()

// "mac", "El Capitan", "10.11.6"
release.All()
```



