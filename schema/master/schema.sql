-- キャラクターマスタ
CREATE TABLE characters (
   id INTEGER NOT NULL PRIMARY KEY, -- キャラクターID
   name TEXT NOT NULL               -- 名前
);

-- 武器マスタ
CREATE TABLE weapons (
   id INTEGER NOT NULL PRIMARY KEY, -- 武器ID
   name TEXT NOT NULL               -- 名前
);

-- クエストマスタ
CREATE TABLE quests (
   id INTEGER NOT NULL PRIMARY KEY, -- クエストID
   name TEXT NOT NULL               -- 名前
);

-- 装備マスタ
CREATE TABLE equipments (
   id INTEGER NOT NULL PRIMARY KEY, -- 装備ID
   name TEXT NOT NULL               -- 名前
);
