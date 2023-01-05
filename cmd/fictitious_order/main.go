package main

import (
	_ "github.com/jackylee92/rgo"
	"github.com/jackylee92/rgo/core/rgrouter"
	"member-system-server/route/fictitious_order"
)

/*
 * @Content : fictitious_order
 * @Author  : LiJunDong
 * @Time    : 2022-09-11$
 */

func main() {
	r := fictitious_order.GetRouter()
	rgrouter.Run(r)
}
