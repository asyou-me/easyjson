package xxx

import (
	"fmt"
	"testing"

	"github.com/asyou-me/easyjson/jlexer"
)

func TestXXX(t *testing.T) {
	r := jlexer.Lexer{Data: []byte(`{"X":111}`)}
	x := &XXX{}
	err := x.UnmarshalEasyJSON(&r)
	fmt.Println("err:", err)
	fmt.Println("c:", x.Fields)
	fmt.Println("x:", x)
}
