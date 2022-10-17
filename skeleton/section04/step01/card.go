package main

// TODO: rarityやcardに関する処理をここに移す
type rarity string

const (
	rarityN  rarity = "N"
	rarityR  rarity = "R"
	raritySR rarity = "SR"
	rarityXR rarity = "XR"
)

func (r rarity) String() string {
	return string(r)
}

type card struct {
	rarity rarity
	name   string
}

func (c *card) String() string {
	return c.rarity.String() + ":" + c.name
}
