package ajson

import (
	"fmt"
	"testing"
)

const json = `{"name":{"first":"Janet","last":"Prichard"},"age":47}`

func TestAJSON(t *testing.T) {
	r := Parse(json)
	r.Set("name.first", "Jospeh")

	fmt.Println(r.Raw)
}
