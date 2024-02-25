package utils

/*

TODO: This is a temporary link dump for now and experimenting with some potential refactorings

https://go-proverbs.github.io/
https://go.dev/blog/context
https://go.dev/blog/context-and-structs
https://medium.com/@parikshit/understanding-the-context-package-in-golang-b1392c821d14
https://www.digitalocean.com/community/tutorials/how-to-use-contexts-in-go

https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/concurrency
https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/sync#when-to-use-locks-over-channels-and-goroutines

https://github.com/cloudflare/tableflip
*/

// Structs define data, interfaces defines behaviour
// Functions should accept interfaces as arguments and return structs
// Given an interface and a type that should implement the interface:
type MyCustomInterface interface{}
type MyCustomStruct struct{}

// Verify that MyCustomType implements MyCustomInterface.
var _ MyCustomInterface = MyCustomStruct{}

// Verify that *MyCustomType implements MyCustomInterface.
var _ MyCustomInterface = (*MyCustomStruct)(nil)
