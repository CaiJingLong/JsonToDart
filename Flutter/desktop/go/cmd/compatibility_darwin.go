package main

// the file compatibility_darwin.go is used to set custom compile fags
 
// #cgo darwin CFLAGS: -mmacosx-version-min=10.13
// #cgo darwin LDFLAGS: -mmacosx-version-min=10.13
// #cgo darwin CXXFLAGS: -mmacosx-version-min=10.13
import "C"