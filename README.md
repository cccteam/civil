# Civil Date and Time types

[![GoDoc](https://godoc.org/github.com/cccteam/civil?status.svg)](https://godoc.org/github.com/cccteam/civil)

Civil provides Date, Time, and DateTime types.

This package was put together to add support for SQL DB interfaces and Json interfaces while maintaining no changes to the upstream package so future updates will flow in easily.

## Source

This civil package was originally copied from `cloud.google.com/go/civil`.
The license remains the same.

Support for support for SQL interfaces (`sql.Scanner` and `driver.Valuer`) and JSON interfaces (`json.Marshal()` and `json.Unmarshal()`) were copied from `github.com/openly-engineering/civil`. A big thankyou for that!
