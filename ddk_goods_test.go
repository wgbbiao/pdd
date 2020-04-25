package pdd

import (
	"fmt"
	"testing"
)

func TestGoodsAPI_GoodsCatGet(t *testing.T) {
	ClientID := "d"
	ClientSecret := "sdf"
	g := newGoodsAPIWithContext(&Context{
		ClientId:     ClientID,
		ClientSecret: ClientSecret,
	})
	r, err := g.GoodsCatGet(0)
	fmt.Println(err)
	fmt.Printf("%v", r)
	for _, value := range r {
		fmt.Printf("%#v\n", value)
	}
}
