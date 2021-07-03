# Next lesson:
https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/maps

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