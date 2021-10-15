package Module

type material struct {
	Index  int
	Name   string
	Amount string
}
type recipe struct {
	RecipeCode int
	Title      string
	Info       string
	Serving    int
	Imgae      string
}
type How struct {
	Index int
	How   string
}
type models struct {
	Status     int
	RecipeInfo []material
	Material   recipe
	How        []How
}

func GetMaterial() []material {

	// 定義その1＆要素追加
	var sample []material
	sample = append(sample, material{
		Index:  1,
		Name:   "きゅうり",
		Amount: "6本",
	})
	sample = append(sample, material{
		Index:  2,
		Name:   "塩",
		Amount: "小さじ2",
	})
	sample = append(sample, material{
		Index:  3,
		Name:   "生姜",
		Amount: "2片",
	})
	sample = append(sample, material{
		Index:  4,
		Name:   "(B)しょうゆ",
		Amount: "大さじ6",
	})
	sample = append(sample, material{
		Index:  5,
		Name:   "(B)みりん",
		Amount: "大さじ4",
	})
	sample = append(sample, material{
		Index:  6,
		Name:   "(B)酢",
		Amount: "大さじ2",
	})
	sample = append(sample, material{
		Index:  7,
		Name:   "(B)砂糖",
		Amount: "大さじ1",
	})
	sample = append(sample, material{
		Index:  8,
		Name:   "(B)鷹の爪輪切り",
		Amount: "小さじ1",
	})
	sample = append(sample, material{
		Index:  9,
		Name:   "にんにく",
		Amount: "2片",
	})

	return sample

}

func GetRecipe() recipe {

	// 定義その1＆要素追加
	var sample recipe
	sample = recipe{
		RecipeCode: 1,
		Title:      "ポリポリ食感 きゅうりの佃煮　レシピ・作り方",
		Info:       "ごはんを食べる箸が止まらなくなるかも！？しょっぱい味が癖になる、ポリポリ食感きゅうりの佃煮です。簡単に作れる上に、食べ方はお好みで無限大に広がります。大量のきゅうりの消費にもオススメですよ。ぜひ作ってみて下さいね。",
		Serving:    4,
		Imgae:      "C:/work/Go/rakushiru/image/food_sample.jpg",
	}
	return sample
}

func GetHow() []How {

	var sample []How
	sample = append(sample, How{
		Index: 1,
		How:   "きゅうりのヘタを取り、薄い輪切りにします。",
	})
	sample = append(sample, How{
		Index: 1,
		How:   "1をボウルに入れ塩を揉み込み、10分置き、出てきた水分を絞ります。",
	})
	sample = append(sample, How{
		Index: 1,
		How:   "生姜を千切りにします。",
	})
	sample = append(sample, How{
		Index: 1,
		How:   "2、3、(A)を鍋に入れて中火にかけます。混ぜながら味を染み込ませ、水分が無くなったら完成です。お好みでごはんに乗せてお召し上がり下さい。",
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

func GetModels() models {

	// 定義その1＆要素追加
	var sample models
	sample = models{
		Status:     200,
		RecipeInfo: GetMaterial(),
		Material:   GetRecipe(),
		How:        GetHow(),
	}

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
