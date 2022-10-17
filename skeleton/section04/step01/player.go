package main

import (
	"fmt"
)

type player struct {
	tickets int
	coin    int
}

func (p *player) drawableNum() int {
	return p.tickets + p.coin/10
}

func (p *player) draw(n int) {
	if p.drawableNum() < n {
		fmt.Println("ガチャ券またはコインが不足しています")
		return
	}

	if p.tickets > n {
		p.tickets -= n
		return
	}

	p.tickets = 0
	p.coin -= n * 10 // 1回あたり10枚消費する
}
