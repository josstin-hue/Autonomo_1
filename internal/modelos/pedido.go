package modelos

type Pedido struct {
	ID         int
	UsuarioID  int
	ProductoID int
	Cantidad   int
	Total      float64
}