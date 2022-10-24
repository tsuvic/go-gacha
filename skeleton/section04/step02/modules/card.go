// package modules

// type Rarity string

// const (
// 	RarityN  Rarity = "N"
// 	RarityR  Rarity = "R"
// 	RaritySR Rarity = "SR"
// 	RarityXR Rarity = "XR"
// )

// func (r Rarity) String() string {
// 	return string(r)
// }

// // TODO: フィールドをエクスポートする
// type Card struct {
// 	Rarity Rarity // レア度
// 	Name   string // 名前
// }

// func (c *Card) String() string {
// 	return c.Rarity.String() + ":" + c.Name
// }

package modules

type Rarity string

const (
	RarityN  Rarity = "N"
	RarityR  Rarity = "R"
	RaritySR Rarity = "SR"
	RarityXR Rarity = "XR"
)

func (r Rarity) String() string {
	return string(r)
}

type Card struct {
	Rarity Rarity
	Name   string
}

func (c *Card) String() string {
	return c.Rarity.String() + ":" + c.Name
}
