package servicios

import (
    "Autonomo_1/internal/modelos"
    "Autonomo_1/internal/repositorios"
)
func AgregarProducto(nombre string, precio float64, stock int) {

	producto := modelos.Producto{
		ID:     len(repositorios.Productos) + 1,
		Nombre: nombre,
		Precio: precio,
		Stock:  stock,
	}

	repositorios.Productos = append(repositorios.Productos, producto)
}

func ListarProductos() []modelos.Producto {
	return repositorios.Productos
}
