package repositorios

import (
	"Autonomo_1/internal/database"
	"Autonomo_1/internal/modelos"
)

func AgregarProducto(nombre string, precio float64, stock int) error {
	_, err := database.DB.Exec(
		`INSERT INTO productos (nombre, precio, stock) VALUES (?, ?, ?)`,
		nombre, precio, stock,
	)
	return err
}

func ListarProductos() ([]modelos.Producto, error) {
	rows, err := database.DB.Query(`SELECT id, nombre, precio, stock FROM productos`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var productos []modelos.Producto
	for rows.Next() {
		var p modelos.Producto
		if err := rows.Scan(&p.ID, &p.Nombre, &p.Precio, &p.Stock); err != nil {
			return nil, err
		}
		productos = append(productos, p)
	}
	return productos, nil
}

func ObtenerProductoPorID(id int) (*modelos.Producto, error) {
	row := database.DB.QueryRow(
		`SELECT id, nombre, precio, stock FROM productos WHERE id = ?`, id,
	)
	var p modelos.Producto
	if err := row.Scan(&p.ID, &p.Nombre, &p.Precio, &p.Stock); err != nil {
		return nil, err
	}
	return &p, nil
}

func EliminarProducto(id int) error {
	_, err := database.DB.Exec(`DELETE FROM productos WHERE id = ?`, id)
	return err
}

func ActualizarProducto(id int, precio float64, stock int) error {
	_, err := database.DB.Exec(
		`UPDATE productos SET precio = ?, stock = ? WHERE id = ?`,
		precio, stock, id,
	)
	return err
}
