package main

// the file compatibility_darwin.go is used to set custom compile fags
 
// #cgo darwin CFLAGS: -mmacosx-version-min=10.14
// #cgo darwin LDFLAGS: -mmacosx-version-min=10.14
// #cgo darwin CXXFLAGS: -mmacosx-version-min=10.14
import "C"