/* YouTube情報テーブル */
CREATE TABLE youtube_info (
    recipe_id VARCHAR(7) NOT NULL -- レシピID
  , youtube_id VARCHAR(40) NOT NULL -- タイトル
  , title VARCHAR(200) -- 紹介
  , del_flg VARCHAR(1) NOT NULL DEFAULT  '0' -- 削除フラグ
  , regist_date DATETIME NOT NULL -- 登録日時
  , update_date DATETIME NOT NULL -- 更新日時
  , PRIMARY KEY (recipe_id)
);

insert into youtube_info(
    recipe_id, 
    youtube_id,
    title,
    del_flg,
    regist_date,
    update_date)
values(
    '0000001', 
    'ZWvXweioDx4', 
    '下茹で不要！チーズ大根ステーキの作り方 #レシピ #shorts #料理 #cooking #recipe', 
    '0', 
    NOW(), 
    NOW());

insert into youtube_info(
    recipe_id, 
    youtube_id,
    title,
    del_flg,
    regist_date,
    update_date)
values(
    '0000002', 
    'Vvm7eHsTMp8', 
    '【糖質制限レシピ】コスパ最強！低糖質・高タンパク！「鶏むね肉と白菜のトマト煮」の作り方', 
    '0', 
    NOW(), 
    NOW());

insert into youtube_info(
    recipe_id, 
    youtube_id,
    title,
    del_flg,
    regist_date,
    update_date)
values(
    '0000001', 
    '5WLd2A4z_3o', 
    '【炙ったら全部美味しい】炙り〇〇レシピ BEST10', 
    '0', 
    NOW(), 
    NOW());
