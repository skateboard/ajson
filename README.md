# AllJson
This package brings together [gJSON](https://github.com/tidwall/gjson) and [sJSON](https://github.com/tidwall/sjson) in 1 library allow you
to read and set json values quickly and effectively.

# Usage
```go
package main

import "github.com/skateboard/ajson"

const json = `{"name":{"first":"Janet","last":"Prichard"},"age":47}`

func main() {
	r := ajson.Parse(json)
	
	beforeValue := r.Get("name.last")
	println(beforeValue.String())
	
	r.Set("name.first", "Jospeh")

	afterValue := r.Get("name.last")
	println(afterValue.String())
	
	r.Create("dob", "4/12/1984")

	bobV := r.Get("dob")
	println(bobV.String())
}
```
