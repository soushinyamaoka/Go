package Module

type material struct {
	index  int
	name   string
	amount string
}
type recipe struct {
	recipeCode int
	title      string
	info       string
	serving    int
	imgae      string
}
type how struct {
	index int
	how   string
}
type models struct {
	recipeInfo []material
	material   recipe
	how        []how
}

func GetMaterial() []material {

	// 定義その1＆要素追加
	var sample []material
	sample = append(sample, material{
		index:  1,
		name:   "きゅうり",
		amount: "6本",
	})
	sample = append(sample, material{
		index:  2,
		name:   "塩",
		amount: "小さじ2",
	})
	sample = append(sample, material{
		index:  3,
		name:   "生姜",
		amount: "2片",
	})
	sample = append(sample, material{
		index:  4,
		name:   "(B)しょうゆ",
		amount: "大さじ6",
	})
	sample = append(sample, material{
		index:  5,
		name:   "(B)みりん",
		amount: "大さじ4",
	})
	sample = append(sample, material{
		index:  6,
		name:   "(B)酢",
		amount: "大さじ2",
	})
	sample = append(sample, material{
		index:  7,
		name:   "(B)砂糖",
		amount: "大さじ1",
	})
	sample = append(sample, material{
		index:  8,
		name:   "(B)鷹の爪輪切り",
		amount: "小さじ1",
	})
	sample = append(sample, material{
		index:  9,
		name:   "にんにく",
		amount: "2片",
	})

	return sample

}

func GetRecipe() recipe {

	// 定義その1＆要素追加
	var sample recipe
	sample = recipe{
		recipeCode: 1,
		title:      "ポリポリ食感 きゅうりの佃煮　レシピ・作り方",
		info:       "ごはんを食べる箸が止まらなくなるかも！？しょっぱい味が癖になる、ポリポリ食感きゅうりの佃煮です。簡単に作れる上に、食べ方はお好みで無限大に広がります。大量のきゅうりの消費にもオススメですよ。ぜひ作ってみて下さいね。",
		serving:    4,
		imgae:      "C:/work/Go/rakushiru/image/food_sample.jpg",
	}
	return sample
}

func GetHow() []how {

	var sample []how
	sample = append(sample, how{
		index: 1,
		how:   "きゅうりのヘタを取り、薄い輪切りにします。",
	})
	sample = append(sample, how{
		index: 1,
		how:   "1をボウルに入れ塩を揉み込み、10分置き、出てきた水分を絞ります。",
	})
	sample = append(sample, how{
		index: 1,
		how:   "生姜を千切りにします。",
	})
	sample = append(sample, how{
		index: 1,
		how:   "2、3、(A)を鍋に入れて中火にかけます。混ぜながら味を染み込ませ、水分が無くなったら完成です。お好みでごはんに乗せてお召し上がり下さい。",
	})

	return sample
}

func GetModels() models {

	// 定義その1＆要素追加
	var sample models
	sample = models{
		recipeInfo: GetMaterial(),
		material:   GetRecipe(),
		how:        GetHow(),
	}
	return sample
}
