package ajson

import (
	"fmt"
	"testing"
)

const json = `{"name":{"first":"Janet","last":"Prichard"},"age":47}`

func TestAJSON(t *testing.T) {
	r := Parse(json)
	r.Create("dob", "4/12/1984")

	v := r.Get("dob")
	fmt.Println(v.String())
}
