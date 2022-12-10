package product

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Product struct {
	ID   string
	Name string
}

func Run() {

	db, err := sql.Open("sqlite3", "app/sqlitedata/product.db")
	if err != nil {
		panic(err)
	}

	prod := Product{
		ID:   "123456",
		Name: "Escova",
	}
	fmt.Println("Produto: ", prod.Name)

	//InsertProduct(db, prod)

	count, err := getCount(db)
	if err != nil {
		panic(err)
	}
	fmt.Println("Quantidade de Produtos: ", count)

	givenProduct, err := getProductByID(db, "123456")
	if err != nil {
		panic(err)
	}
	fmt.Println("Get Produto: ", givenProduct.Name)
	//

}

func InsertProduct(db *sql.DB, p Product) {
	db.Exec("INSERT INTO products (id, name) VALUES (?, ?)", p.ID, p.Name)
}

func getProductByID(db *sql.DB, productID string) (Product, error) {
	var id string
	var name string

	const query = `
		SELECT 
			id, 
			name 
		FROM 
			products 
		WHERE 
			id=$1
	`
	row := db.QueryRow(query, productID)

	nErr := row.Scan(&id, &name)
	if nErr != nil {
		return Product{}, fmt.Errorf("erro to to scan in get product by id %w", nErr)
	}

	p := Product{
		ID:   id,
		Name: name,
	}
	return p, nil
}

func getCount(db *sql.DB) (int, error) {
	var count int
	const query = `
		SELECT 
			count(*)
		FROM 
			products
	`
	row := db.QueryRow(query)

	nErr := row.Scan(&count)
	if nErr != nil {
		return 0, fmt.Errorf("erro to to scan in get count %w", nErr)
	}

	return count, nil
}
