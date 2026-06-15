package servicios

import (
	"Autonomo_1/internal/modelos"
	"Autonomo_1/internal/repositorios"
)

func CrearPedido(usuarioID int, productoID int, cantidad int) {

	for _, producto := range repositorios.Productos {

		if producto.ID == productoID {

			total := producto.Precio * float64(cantidad)

			pedido := modelos.Pedido{
				ID:         len(repositorios.Pedidos) + 1,
				UsuarioID:  usuarioID,
				ProductoID: productoID,
				Cantidad:   cantidad,
				Total:      total,
			}

			repositorios.Pedidos = append(repositorios.Pedidos, pedido)

			break
		}
	}
}

func ListarPedidos() []modelos.Pedido {
	return repositorios.Pedidos
}
