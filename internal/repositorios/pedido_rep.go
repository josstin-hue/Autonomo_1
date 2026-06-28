package repositorios

import (
	"Autonomo_1/internal/database"
	"Autonomo_1/internal/modelos"
)

func CrearPedido(usuarioID int, items []modelos.ItemPedido, total float64) (int64, error) {
	tx, err := database.DB.Begin()
	if err != nil {
		return 0, err
	}

	res, err := tx.Exec(
		`INSERT INTO pedidos (usuario_id, total) VALUES (?, ?)`,
		usuarioID, total,
	)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	pedidoID, _ := res.LastInsertId()

	for _, item := range items {
		_, err := tx.Exec(
			`INSERT INTO pedido_items (pedido_id, producto_id, cantidad, subtotal)
			 VALUES (?, ?, ?, ?)`,
			pedidoID, item.ProductoID, item.Cantidad, item.Subtotal,
		)
		if err != nil {
			tx.Rollback()
			return 0, err
		}
	}

	return pedidoID, tx.Commit()
}

func ListarPedidos() ([]modelos.Pedido, error) {
	rows, err := database.DB.Query(`
		SELECT p.id, p.usuario_id, COALESCE(u.nombre, 'Desconocido'), p.total
		FROM pedidos p
		LEFT JOIN usuarios u ON u.id = p.usuario_id
		ORDER BY p.id DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pedidos []modelos.Pedido
	for rows.Next() {
		var p modelos.Pedido
		if err := rows.Scan(&p.ID, &p.UsuarioID, &p.NombreUsuario, &p.Total); err != nil {
			return nil, err
		}
		items, err := listarItemsPedido(p.ID)
		if err != nil {
			return nil, err
		}
		p.Items = items
		pedidos = append(pedidos, p)
	}
	return pedidos, nil
}

func ObtenerPedidoPorID(id int) (*modelos.Pedido, error) {
	row := database.DB.QueryRow(`
		SELECT p.id, p.usuario_id, COALESCE(u.nombre, 'Desconocido'), p.total
		FROM pedidos p
		LEFT JOIN usuarios u ON u.id = p.usuario_id
		WHERE p.id = ?
	`, id)

	var p modelos.Pedido
	if err := row.Scan(&p.ID, &p.UsuarioID, &p.NombreUsuario, &p.Total); err != nil {
		return nil, err
	}

	items, err := listarItemsPedido(p.ID)
	if err != nil {
		return nil, err
	}
	p.Items = items
	return &p, nil
}

func listarItemsPedido(pedidoID int) ([]modelos.ItemPedido, error) {
	rows, err := database.DB.Query(`
		SELECT pi.producto_id, COALESCE(pr.nombre, 'Producto eliminado'), pi.cantidad, pi.subtotal
		FROM pedido_items pi
		LEFT JOIN productos pr ON pr.id = pi.producto_id
		WHERE pi.pedido_id = ?
	`, pedidoID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []modelos.ItemPedido
	for rows.Next() {
		var item modelos.ItemPedido
		if err := rows.Scan(&item.ProductoID, &item.NombreProducto, &item.Cantidad, &item.Subtotal); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}
