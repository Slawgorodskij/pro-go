DROP TABLE IF EXISTS OrderLines;
DROP TABLE IF EXISTS Orders;
DROP TABLE IF EXISTS Products;
DROP TABLE IF EXISTS Categories;
CREATE TABLE IF NOT EXISTS Categories
(
    Id   INTEGER NOT NULL PRIMARY KEY,
    Name TEXT
);

CREATE TABLE IF NOT EXISTS Products
(
    Id          INTEGER NOT NULL PRIMARY KEY,
    Name        TEXT,
    Description TEXT,
    Category    INTEGER,
    Price       decimal(8, 2),
    CONSTRAINT CatRef FOREIGN KEY (Category) REFERENCES Categories (Id)
);

CREATE TABLE IF NOT EXISTS OrderLines
(
    Id        INTEGER NOT NULL PRIMARY KEY,
    OrderId   INT,
    ProductId INT,
    Quantity  INT,
    CONSTRAINT OrderRef FOREIGN KEY (ProductId) REFERENCES Products (Id),
    CONSTRAINT OrderRef FOREIGN KEY (OrderId) REFERENCES Orders (Id)
);
CREATE TABLE IF NOT EXISTS Orders
(
    Id         INTEGER NOT NULL PRIMARY KEY,
    Name       TEXT    NOT NULL,
    StreetAddr TEXT    NOT NULL,
    City       TEXT    NOT NULL,
    Zip        TEXT    NOT NULL,
    Country    TEXT    NOT NULL,
    Shipped    BOOLEAN
);