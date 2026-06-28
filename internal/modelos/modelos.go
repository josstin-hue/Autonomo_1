package modelos

type Producto struct {
	ID     int     `json:"ID"`
	Nombre string  `json:"Nombre"`
	Precio float64 `json:"Precio"`
	Stock  int     `json:"Stock"`
}

type Usuario struct {
	ID     int    `json:"ID"`
	Nombre string `json:"Nombre"`
	Correo string `json:"Correo"`
}

type ItemPedido struct {
	ProductoID     int     `json:"producto_id"`
	NombreProducto string  `json:"nombre_producto"`
	Cantidad       int     `json:"cantidad"`
	Subtotal       float64 `json:"subtotal"`
}

type Pedido struct {
	ID            int          `json:"ID"`
	UsuarioID     int          `json:"UsuarioID"`
	NombreUsuario string       `json:"NombreUsuario"`
	Items         []ItemPedido `json:"Items"`
	Total         float64      `json:"Total"`
}

type ReporteVenta struct {
	ProductoID       int     `json:"producto_id"`
	NombreProducto   string  `json:"nombre_producto"`
	UnidadesVendidas int     `json:"unidades_vendidas"`
	TotalRecaudado   float64 `json:"total_recaudado"`
}
