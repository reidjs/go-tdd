# Next lesson:
https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/dependency-injection

# View docs
godoc -http=:6060

http://localhost:6060/pkg/

# Run test
go test

# pointers
access memory location using `&` like `&varname`

declare pointers as function params using the `*` like `w *Wallet`

You read as *Wallet as "A pointer to a wallet"

Afterwards, you can us the variable without prefixing the pointer symbol (just use `w`, not `(*w)`). However, you absolutely can use `(*w)` if you want.

# creating types 
You can alias types from existing ones
`type MyName OriginalType`
`type Dollars int`

## maps
key type must be a 'comparable type'
https://golang.org/ref/spec#Comparison_operators

value can be whatever you want, even another map

maps return two values when looking up values, the value, and an error if it can't be found

Map values are pointers to a runtime.hmap structure
when you pass a map to a function you are copying the pointer parts, but not the underlying data structure with the actual data. 

Never initialize an empty map variable
Instead, you can init an empty map using make
`var dictionary = make(map[string]string)`
or 
`var dictionary = map[string]string{}`
