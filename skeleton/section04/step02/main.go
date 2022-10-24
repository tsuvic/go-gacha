// // STEP01: ファイルを分けよう

// package main

// import (
// 	"fmt"

// 	"gacha.section04.step02/modules"
// )

// func main() {
// 	// TODO: ガチャ券10枚、コイン100枚を持ったプレイヤーを作る
// 	p := modules.NewPlayer(10, 100)D

// 	n := inputN(p)
// 	// TODO: gacha.DrawN関数を呼び、変数resultsとsummaryに結果を代入する
// 	results, summary := modules.DrawN(p, n)

// 	fmt.Println(results)
// 	fmt.Println(summary)
// }

// // TODO: 引数の型をgacha.Playerのポインタにする
// func inputN(p *modules.Player) int {

// 	max := p.DrawableNum()
// 	fmt.Printf("ガチャを引く回数を入力してください（最大:%d回）\n", max)

// 	var n int
// 	for {
// 		fmt.Print("ガチャを引く回数>")
// 		fmt.Scanln(&n)
// 		if n > 0 && n <= max {
// 			break
// 		}
// 		fmt.Printf("1以上%d以下の数を入力してください\n", max)
// 	}

// 	return n
// }

package main

import (
	"fmt"

	"gacha.section04.step02/modules"
)

func main() {
	p := modules.NewPlayer(10, 100)
	n := inputN(p)
	results, summary := modules.DrawN(p, n)
	fmt.Println(results)
	fmt.Println(summary)
}

func inputN(p *modules.Player) int {
	max := p.DrawableNum()
	fmt.Printf("ガチャを引く回数を入力してください（最大：%d回）\n", max)

	var n int
	for {
		fmt.Print("ガチャを引く回数>")
		fmt.Scanln(&n)
		if max >= n && n > 0 {
			break
		}
		fmt.Printf("1以上%d以下の数を入力してください\n", max)
	}
	return n
}
