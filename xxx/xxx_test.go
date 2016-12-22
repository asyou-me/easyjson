package xxx

import (
	"fmt"
	"testing"

	"github.com/asyou-me/easyjson/jlexer"
)

func TestXXX(t *testing.T) {
	r := jlexer.Lexer{Data: []byte(`{"X":"cccc"}`)}
	x := &XXX{}
	err := x.UnmarshalEasyJSON(&r)
	fmt.Println("c:", x.Fields)
	fmt.Println("err:", err)
	fmt.Println("x:", x)
}
