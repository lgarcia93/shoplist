package db

var dbDDL = `CREATE TABLE if not exists ShopItem (
  id INTEGER auto_increment PRIMARY KEY ,
  title TEXT NOT NULL,
  description TEXT NOT NULL,
  price double NOT NULL
);
`
