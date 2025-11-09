// fields, structs, types, structures, lists, urls, names, subdomains... we configure all here.

package config

// all those structs are main models of our json structure, example in structure.json
type Skin struct {
	Name       string `json:"name"`
	Weapon     string `json:"weapon"`
	Rarity     string `json:"rarity"`
	Collection string `json:"collection"`
	Price      Price  `json:"price"`
	URL        string `json:"url"`
}

type Price struct {
	PriceString         string     `json:"price_string"`
	PriceStattrakString string     `json:"price_stattrak_string"`
	Currency            string     `json:"currency"`
	Min                 PriceValue `json:"min"`
	Max                 PriceValue `json:"max"`
	UpdatedAt           string     `json:"updated_at"`
}

type PriceValue struct {
	Value         float64 `json:"value"`
	StattrakValue float64 `json:"stattrak_value"`
	Unit          string  `json:"unit"`
}

// here you should add subdomains (in our case name of collection/case)
// list := []string{
// 	"kilowatt-case", "revolution-case", "recoil-case", "dreams-nightmares-case", "sealed-genesis-terminal",
// 	"snakebite-case", "fracture-case", "prisma-2-case", "cs20-case", "prisma-case", "danger-zone-case",
// 	"horizon-case", "clutch-case", "spectrum-2-case", "operation-hydra-case", "spectrum-case", "glove-case",
// }
