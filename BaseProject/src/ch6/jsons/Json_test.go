package jsons

import (
	"encoding/json"
	"fmt"
	"structs"
	"testing"
)

func TestJSON(t *testing.T)  {

	p := structs.Person{Age:1,Tall:2}
	bytes, e := json.Marshal(p)
	fmt.Println(e,string(bytes))
	fmt.
}
