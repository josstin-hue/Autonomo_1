package repositorios

import (
	"Autonomo_1/internal/database"
	"Autonomo_1/internal/modelos"
)

func ReporteVentas() ([]modelos.ReporteVenta, error) {
	rows, err := database.DB.Query(`
		SELECT
			pr.id,
			pr.nombre,
			COALESCE(SUM(pi.cantidad), 0)  AS unidades_vendidas,
			COALESCE(SUM(pi.subtotal), 0)  AS total_recaudado
		FROM productos pr
		LEFT JOIN pedido_items pi ON pi.producto_id = pr.id
		GROUP BY pr.id, pr.nombre
		ORDER BY total_recaudado DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reporte []modelos.ReporteVenta
	for rows.Next() {
		var r modelos.ReporteVenta
		if err := rows.Scan(&r.ProductoID, &r.NombreProducto, &r.UnidadesVendidas, &r.TotalRecaudado); err != nil {
			return nil, err
		}
		reporte = append(reporte, r)
	}
	return reporte, nil
}
