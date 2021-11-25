package Module

import "time"

type Recipes struct {
	RecipeId     string    `gorm:"recipe_id" json:"RecipeId, omitempty"`
	Title        string    `gorm:"title" json:'Title, omitempty"`
	Introduction string    `gorm:"introduction" json:'Introduction, omitempty"`
	Serving      string    `gorm:"serving" json:'Serving, omitempty"`
	Image        string    `gorm:"image" json:'Image, omitempty"`
	CreatedAt    time.Time `gorm:"created_at" json:'CreatedAt'`
	UpdatedAt    time.Time `gorm:"updated_at" json:'UpdatedAt'`
}

type Ingredients struct {
	RecipeId  string    `gorm:"recipe_id" json:"RecipeId, omitempty"`
	OrderNo   int       `gorm:"order_no	" json:'OrderNo, omitempty"`
	Name      string    `gorm:"name	" json:'Name, omitempty"`
	Quantity  string    `gorm:"quantity" json:'Quantity, omitempty"`
	CreatedAt time.Time `gorm:"created_at" json:'CreatedAt'`
	UpdatedAt time.Time `gorm:"updated_at" json:'UpdatedAt'`
}

type Instructions struct {
	RecipeId  string    `gorm:"recipe_id" json:"RecipeId, omitempty"`
	OrderNo   int       `gorm:"order_no	" json:'OrderNo, omitempty"`
	Detail    string    `gorm:"detail	" json:'Detail	, omitempty"`
	CreatedAt time.Time `gorm:"created_at" json:'CreatedAt'`
	UpdatedAt time.Time `gorm:"updated_at" json:'UpdatedAt'`
}

type Data struct {
	Word string `json:"word"`
}

type Models struct {
	Status       int
	Recipes      []Recipes
	Ingredients  []Ingredients
	Instructions []Instructions
}

func GetMaterial() []Ingredients {

	// 定義その1＆要素追加
	var sample []Ingredients
	sample = append(sample, Ingredients{
		OrderNo:  1,
		Name:     "きゅうり",
		Quantity: "6本",
	})
	sample = append(sample, Ingredients{
		OrderNo:  2,
		Name:     "塩",
		Quantity: "小さじ2",
	})
	sample = append(sample, Ingredients{
		OrderNo:  3,
		Name:     "生姜",
		Quantity: "2片",
	})
	sample = append(sample, Ingredients{
		OrderNo:  4,
		Name:     "(B)しょうゆ",
		Quantity: "大さじ6",
	})
	sample = append(sample, Ingredients{
		OrderNo:  5,
		Name:     "(B)みりん",
		Quantity: "大さじ4",
	})
	sample = append(sample, Ingredients{
		OrderNo:  6,
		Name:     "(B)酢",
		Quantity: "大さじ2",
	})
	sample = append(sample, Ingredients{
		OrderNo:  7,
		Name:     "(B)砂糖",
		Quantity: "大さじ1",
	})
	sample = append(sample, Ingredients{
		OrderNo:  8,
		Name:     "(B)鷹の爪輪切り",
		Quantity: "小さじ1",
	})
	sample = append(sample, Ingredients{
		OrderNo:  9,
		Name:     "にんにく",
		Quantity: "2片",
	})

	return sample

}

func GetRecipe() Recipes {

	// 定義その1＆要素追加
	var sample Recipes
	sample = Recipes{
		RecipeId:     "1",
		Title:        "ポリポリ食感 きゅうりの佃煮　レシピ・作り方",
		Introduction: "ごはんを食べる箸が止まらなくなるかも！？しょっぱい味が癖になる、ポリポリ食感きゅうりの佃煮です。簡単に作れる上に、食べ方はお好みで無限大に広がります。大量のきゅうりの消費にもオススメですよ。ぜひ作ってみて下さいね。",
		Serving:      "4",
		Image:        "/image/food_sample.jpg",
		//mgae: "../Go/rakushiru/image/food_sample.jpg",
	}
	return sample
}

func GetInstructions() []Instructions {

	var sample []Instructions
	sample = append(sample, Instructions{
		OrderNo: 1,
		Detail:  "きゅうりのヘタを取り、薄い輪切りにします。",
	})
	sample = append(sample, Instructions{
		OrderNo: 1,
		Detail:  "1をボウルに入れ塩を揉み込み、10分置き、出てきた水分を絞ります。",
	})
	sample = append(sample, Instructions{
		OrderNo: 1,
		Detail:  "生姜を千切りにします。",
	})
	sample = append(sample, Instructions{
		OrderNo: 1,
		Detail:  "2、3、(A)を鍋に入れて中火にかけます。混ぜながら味を染み込ませ、水分が無くなったら完成です。お好みでごはんに乗せてお召し上がり下さい。",
	})

	return sample
}

type sample struct {
	Status     int
	RecipeCode int
	Title      string
	Info       string
	Serving    int
	Imgae      string
}

func GetModels() Models {

	// 定義その1＆要素追加
	var sample Models
	// sample = Models{
	// 	Status:      200,
	// 	Recipes:  GetRecipe(),
	// 	Ingredients:    GetMaterial(),
	// 	Instructions: GetInstructions(),
	// }

	// ping := sample{
	// 	http.StatusOK,
	// 	1,
	// 	"ポリポリ食感 きゅうりの佃煮　レシピ・作り方",
	// 	"ごはんを食べる箸が止まらなくなるかも！？しょっぱい味が癖になる、ポリポリ食感きゅうりの佃煮です。簡単に作れる上に、食べ方はお好みで無限大に広がります。大量のきゅうりの消費にもオススメですよ。ぜひ作ってみて下さいね。",
	// 	4,
	// 	"C:/work/Go/rakushiru/image/food_sample.jpg",
	// }
	// return ping
	return sample
}
