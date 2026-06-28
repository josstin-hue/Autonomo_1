package servicios

import (
	"Autonomo_1/internal/modelos"
	"Autonomo_1/internal/repositorios"
	"errors"
)

// ── PRODUCTOS ────────────────────────────────────────────────────

func AgregarProducto(nombre string, precio float64, stock int) error {
	return repositorios.AgregarProducto(nombre, precio, stock)
}

func ListarProductos() ([]modelos.Producto, error) {
	return repositorios.ListarProductos()
}

func EliminarProducto(id int) error {
	if id <= 0 {
		return errors.New("ID inválido")
	}
	return repositorios.EliminarProducto(id)
}

func ActualizarProducto(id int, precio float64, stock int) error {
	if id <= 0 {
		return errors.New("ID inválido")
	}
	if precio < 0 {
		return errors.New("El precio no puede ser negativo")
	}
	if stock < 0 {
		return errors.New("El stock no puede ser negativo")
	}
	return repositorios.ActualizarProducto(id, precio, stock)
}

// ── USUARIOS ─────────────────────────────────────────────────────

func RegistrarUsuario(nombre, correo string) error {
	return repositorios.RegistrarUsuario(nombre, correo)
}

func ListarUsuarios() ([]modelos.Usuario, error) {
	return repositorios.ListarUsuarios()
}

// ── PEDIDOS ──────────────────────────────────────────────────────

type ItemEntrada struct {
	ProductoID int `json:"producto_id"`
	Cantidad   int `json:"cantidad"`
}

func CrearPedido(usuarioID int, itemsEntrada []ItemEntrada) error {
	if len(itemsEntrada) == 0 {
		return errors.New("el pedido debe tener al menos un producto")
	}

	var items []modelos.ItemPedido
	total := 0.0

	for _, entrada := range itemsEntrada {
		producto, err := repositorios.ObtenerProductoPorID(entrada.ProductoID)
		if err != nil {
			return errors.New("producto no encontrado")
		}
		subtotal := producto.Precio * float64(entrada.Cantidad)
		total += subtotal
		items = append(items, modelos.ItemPedido{
			ProductoID: entrada.ProductoID,
			Cantidad:   entrada.Cantidad,
			Subtotal:   subtotal,
		})
	}

	_, err := repositorios.CrearPedido(usuarioID, items, total)
	return err
}

func ListarPedidos() ([]modelos.Pedido, error) {
	return repositorios.ListarPedidos()
}

func ObtenerPedido(id int) (*modelos.Pedido, error) {
	if id <= 0 {
		return nil, errors.New("ID inválido")
	}
	return repositorios.ObtenerPedidoPorID(id)
}

// ── REPORTES ─────────────────────────────────────────────────────

func ReporteVentas() ([]modelos.ReporteVenta, error) {
	return repositorios.ReporteVentas()
}
