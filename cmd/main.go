package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"Autonomo_1/internal/servicios"
	"Autonomo_1/internal/utilidad"
)

func main() {

	var opcion int

	for {

		utilidad.MostrarTitulo()

		fmt.Println("===== MENÚ PRINCIPAL =====")
		fmt.Println("1 -> Agregar producto")
		fmt.Println("2 -> Listar productos")
		fmt.Println("3 -> Registrar usuario")
		fmt.Println("5 -> Crear pedido")
		fmt.Println("6 -> Listar pedidos")
		fmt.Println("0 -> Salir")

		fmt.Print("\nSeleccione una opción: ")
		fmt.Scan(&opcion)

		switch opcion {

		case 1:

			fmt.Println("\n=== AGREGAR PRODUCTO ===")

			var precio float64
			var stock int

			fmt.Scanln()

			reader := bufio.NewReader(os.Stdin)

			fmt.Print("Nombre del producto: ")
			nombre, _ := reader.ReadString('\n')
			nombre = strings.TrimSpace(nombre)

			fmt.Print("Precio: ")
			fmt.Scan(&precio)

			fmt.Print("Stock: ")
			fmt.Scan(&stock)

			servicios.AgregarProducto(nombre, precio, stock)

			fmt.Println("\nProducto agregado correctamente.")

		case 2:

			fmt.Println("\n=== LISTA DE PRODUCTOS ===")

			productos := servicios.ListarProductos()

			if len(productos) == 0 {

				fmt.Println("No existen productos registrados.")

			} else {

				fmt.Println("--------------------------------------------------------------------------")
				fmt.Printf("%-5s %-40s %-12s %-10s\n", "ID", "NOMBRE", "PRECIO", "STOCK")
				fmt.Println("--------------------------------------------------------------------------")

				for _, producto := range productos {

					fmt.Printf(
						"%-5d %-40s %-12.2f %-10d\n",
						producto.ID,
						producto.Nombre,
						producto.Precio,
						producto.Stock,
					)
				}

				fmt.Println("--------------------------------------------------------------------------")
			}

		case 3:

			fmt.Println("\n=== REGISTRAR USUARIO ===")

			fmt.Scanln()

			reader := bufio.NewReader(os.Stdin)

			fmt.Print("Nombre: ")
			nombre, _ := reader.ReadString('\n')
			nombre = strings.TrimSpace(nombre)

			fmt.Print("Correo: ")
			correo, _ := reader.ReadString('\n')
			correo = strings.TrimSpace(correo)

			servicios.RegistrarUsuario(nombre, correo)

			fmt.Println("\nUsuario registrado correctamente.")

		case 4:

			fmt.Println("\n=== LISTA DE USUARIOS ===")

			usuarios := servicios.ListarUsuarios()

			if len(usuarios) == 0 {

				fmt.Println("No existen usuarios registrados.")

			} else {

				fmt.Println("----------------------------------------------------------------------------")
				fmt.Printf("%-5s %-35s %-35s\n", "ID", "NOMBRE", "CORREO")
				fmt.Println("----------------------------------------------------------------------------")

				for _, usuario := range usuarios {

					fmt.Printf(
						"%-5d %-35s %-35s\n",
						usuario.ID,
						usuario.Nombre,
						usuario.Correo,
					)
				}

				fmt.Println("----------------------------------------------------------------------------")
			}

		case 5:

			fmt.Println("\n=== CREAR PEDIDO ===")

			var usuarioID int
			var productoID int
			var cantidad int

			fmt.Print("ID Usuario: ")
			fmt.Scan(&usuarioID)

			fmt.Print("ID Producto: ")
			fmt.Scan(&productoID)

			fmt.Print("Cantidad: ")
			fmt.Scan(&cantidad)

			servicios.CrearPedido(usuarioID, productoID, cantidad)

			fmt.Println("\nPedido creado correctamente.")

		case 6:

			fmt.Println("\n=== LISTA DE PEDIDOS ===")

			pedidos := servicios.ListarPedidos()

			if len(pedidos) == 0 {

				fmt.Println("No existen pedidos registrados.")

			} else {

				fmt.Println("----------------------------------------------------------------")
				fmt.Printf("%-5s %-10s %-10s %-10s %-10s\n",
					"ID",
					"USUARIO",
					"PRODUCTO",
					"CANTIDAD",
					"TOTAL",
				)

				fmt.Println("----------------------------------------------------------------")

				for _, pedido := range pedidos {

					fmt.Printf(
						"%-5d %-10d %-10d %-10d %-10.2f\n",
						pedido.ID,
						pedido.UsuarioID,
						pedido.ProductoID,
						pedido.Cantidad,
						pedido.Total,
					)
				}

				fmt.Println("----------------------------------------------------------------")
			}

		case 0:

			fmt.Println("\nGracias por utilizar el sistema.")
			return

		default:

			fmt.Println("\nOpción inválida.")
		}

		fmt.Println("\nPresione ENTER para continuar...")
		fmt.Scanln()
		fmt.Scanln()
	}
}
