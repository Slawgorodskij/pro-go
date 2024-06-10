package main

import "database/sql"

type Category struct {
	Id   int
	Name string
}

type Product struct {
	Id   int
	Name string
	Category
	Price float64
}

//func queryDatabase(db *sql.DB) []Product {
//	products := []Product{}
//	rows, err := db.Query(`SELECT
//    Products.Id,
//    Products.Name,
//    Products.Price,
//    Categories.Id as Cat_Id,
//    Categories.Name as Cat_Name
//    FROM Products, Categories
//	WHERE Products.Category = Categories.Id
//    `)
//	if err == nil {
//		for rows.Next() {
//			p := Product{}
//			scanErr := rows.Scan(&p.Id, &p.Name, &p.Price, &p.Category.Id, &p.Category.Name)
//			if scanErr == nil {
//				products = append(products, p)
//			} else {
//				Printfln("Scan error: %v", scanErr)
//				break
//			}
//		}
//	} else {
//		Printfln("Error: %v", err)
//	}
//	return products
//}

//func main() {
//
//	db, err := openDatabase()
//	if err == nil {
//		products := queryDatabase(db)
//		for i, p := range products {
//			Printfln("#%v: %v", i, p)
//		}
//		db.Close()
//	} else {
//		panic(err)
//	}
//}

// Выполнение операторов с заполнителями

//func queryDatabase(db *sql.DB, categoryName string) []Product {
//	products := []Product{}
//	rows, err := db.Query(`
//	SELECT
//    	Products.Id,
//    	Products.Name,
//    	Products.Price,
//    	Categories.Id as Cat_Id,
//    	Categories.Name as Cat_Name
//    FROM Products, Categories
//	WHERE Products.Category = Categories.Id
//		AND Categories.Name = ?`, categoryName)
//	if err == nil {
//		for rows.Next() {
//			p := Product{}
//			scanErr := rows.Scan(&p.Id, &p.Name, &p.Price, &p.Category.Id, &p.Category.Name)
//			if scanErr == nil {
//				products = append(products, p)
//			} else {
//				Printfln("Scan error: %v", scanErr)
//				break
//			}
//		}
//	} else {
//		Printfln("Error: %v", err)
//	}
//	return products
//}
//
//func main() {
//
//	db, err := openDatabase()
//	if err == nil {
//		for _, cat := range []string{"Soccer", "Watersports"} {
//			Printfln("--- %v Result ---", cat)
//			products := queryDatabase(db, cat)
//			for i, p := range products {
//				Printfln("#%v: %v %v %v", i, p.Name, p.Category.Name, p.Price)
//			}
//		}
//		db.Close()
//	} else {
//		panic(err)
//	}
//}

// Выполнение запросов для отдельных строк

//func queryDatabase(db *sql.DB, id int) (p Product) {
//	row := db.QueryRow(`
//	SELECT
//    	Products.Id,
//    	Products.Name,
//    	Products.Price,
//    	Categories.Id as Cat_Id,
//    	Categories.Name as Cat_Name
//    FROM Products, Categories
//	WHERE Products.Category = Categories.Id
//		AND Products.Id = ?`, id)
//
//	if row.Err() == nil {
//		scanErr := row.Scan(&p.Id, &p.Name, &p.Price, &p.Category.Id, &p.Category.Name)
//		if scanErr != nil {
//			Printfln("Scan error: %v", scanErr)
//		}
//	} else {
//		Printfln("Row error: %v", row.Err().Error())
//	}
//	return
//}
//
//func main() {
//
//	db, err := openDatabase()
//	if err == nil {
//		for _, id := range []int{1, 3, 10} {
//			p := queryDatabase(db, id)
//			Printfln("Product: %v", p)
//		}
//		db.Close()
//	} else {
//		panic(err)
//	}
//}

// Выполнение других запросов

//func queryDatabase(db *sql.DB, id int) (p Product) {
//	row := db.QueryRow(`
//	SELECT
//    	Products.Id,
//    	Products.Name,
//    	Products.Price,
//    	Categories.Id as Cat_Id,
//    	Categories.Name as Cat_Name
//    FROM Products, Categories
//	WHERE Products.Category = Categories.Id
//		AND Products.Id = ?`, id)
//
//	if row.Err() == nil {
//		scanErr := row.Scan(&p.Id, &p.Name, &p.Price, &p.Category.Id, &p.Category.Name)
//		if scanErr != nil {
//			Printfln("Scan error: %v", scanErr)
//		}
//	} else {
//		Printfln("Row error: %v", row.Err().Error())
//	}
//	return
//}
//
//func insertRow(db *sql.DB, p *Product) (id int64) {
//	res, err := db.Exec(`
//	INSERT INTO Products (Name, Category, Price)
//	VALUES (?, ?, ?)`, p.Name, p.Category.Id, p.Price)
//	if err == nil {
//		id, err = res.LastInsertId()
//		if err != nil {
//			Printfln("Result error: %v", err.Error())
//		}
//	} else {
//		Printfln("Exec error: %v", err.Error())
//	}
//	return
//}
//
//func main() {
//
//	db, err := openDatabase()
//	if err == nil {
//		newProduct := Product{Name: "Stadium", Category: Category{Id: 2}, Price: 79500}
//		newID := insertRow(db, &newProduct)
//		p := queryDatabase(db, int(newID))
//		Printfln("New Product: %v", p)
//		db.Close()
//	} else {
//		panic(err)
//	}
//}

// Использование подготовленных операторов

//func queryDatabase(db *sql.DB, id int) (p Product) {
//	row := db.QueryRow(`
//	SELECT
//    	Products.Id,
//    	Products.Name,
//    	Products.Price,
//    	Categories.Id as Cat_Id,
//    	Categories.Name as Cat_Name
//    FROM Products, Categories
//	WHERE Products.Category = Categories.Id
//		AND Products.Id = ?`, id)
//
//	if row.Err() == nil {
//		scanErr := row.Scan(&p.Id, &p.Name, &p.Price, &p.Category.Id, &p.Category.Name)
//		if scanErr != nil {
//			Printfln("Scan error: %v", scanErr)
//		}
//	} else {
//		Printfln("Row error: %v", row.Err().Error())
//	}
//	return
//}
//
//func insertAndUseCategory(name string, productIDs ...int) {
//	res, err := insertNewCategory.Exec(name)
//	if err == nil {
//		newID, _ := res.LastInsertId()
//		for _, id := range productIDs {
//			changeProductCategory.Exec(int(newID), id)
//		}
//	} else {
//		Printfln("Prepared statement error: %v", err)
//	}
//}
//
//func main() {
//
//	db, err := openDatabase()
//	if err == nil {
//		insertAndUseCategory("Misc Products", 2)
//		p := queryDatabase(db, 2)
//		Printfln("Product: %v", p)
//		db.Close()
//	} else {
//		panic(err)
//	}
//}

// Использование транзакций

//func queryDatabase(db *sql.DB, id int) (p Product) {
//	row := db.QueryRow(`
//	SELECT
//    	Products.Id,
//    	Products.Name,
//    	Products.Price,
//    	Categories.Id as Cat_Id,
//    	Categories.Name as Cat_Name
//    FROM Products, Categories
//	WHERE Products.Category = Categories.Id
//		AND Products.Id = ?`, id)
//
//	if row.Err() == nil {
//		scanErr := row.Scan(&p.Id, &p.Name, &p.Price, &p.Category.Id, &p.Category.Name)
//		if scanErr != nil {
//			Printfln("Scan error: %v", scanErr)
//		}
//	} else {
//		Printfln("Row error: %v", row.Err().Error())
//	}
//	return
//}
//
//func insertAndUseCategory(db *sql.DB, name string, productIDs ...int) (err error) {
//	tx, err := db.Begin()
//	updatedFailed := false
//	if err == nil {
//		catResult, err := tx.Stmt(insertNewCategory).Exec(name)
//		if err == nil {
//			newID, _ := catResult.LastInsertId()
//			preparedStatement := tx.Stmt(changeProductCategory)
//			for _, id := range productIDs {
//				changeResult, err := preparedStatement.Exec(newID, id)
//				if err == nil {
//					changes, _ := changeResult.RowsAffected()
//					if changes == 0 {
//						updatedFailed = true
//						break
//					}
//				}
//			}
//		}
//	}
//	if err != nil || updatedFailed {
//		Printfln("Aborting transaction %v", err)
//		tx.Rollback()
//	} else {
//		tx.Commit()
//	}
//	return
//}
//
//func main() {
//
//	db, err := openDatabase()
//	if err == nil {
//		insertAndUseCategory(db, "Category_1", 2)
//		p := queryDatabase(db, 2)
//		Printfln("Product: %v", p)
//		insertAndUseCategory(db, "Category_2", 100)
//		db.Close()
//	} else {
//		panic(err)
//	}
//}

// Использование рефлексии для сканирования данных в структуру

func queryDatabase(db *sql.DB) (products []Product, err error) {
	rows, err := db.Query(`
	SELECT 
	    Products.Id, 
	    Products.Name,
		Products.Price,
		Categories.Id as "Category.Id", 
		Categories.Name as "Category.Name"
	FROM Products, Categories
	WHERE Products.Category = Categories.Id`)

	if err != nil {
		return
	} else {
		results, err := scanIntoStruct(rows, &Product{})
		if err == nil {
			products = (results).([]Product)
		} else {
			Printfln("Scanning error: %v", err)
		}
	}
	return
}

func main() {

	db, err := openDatabase()
	if err == nil {
		products, _ := queryDatabase(db)
		for _, p := range products {
			Printfln("Product: %v", p)
		}
		db.Close()
	} else {
		panic(err)
	}
}
