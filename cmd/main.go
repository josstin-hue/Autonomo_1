package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"Autonomo_1/internal/database"
	"Autonomo_1/internal/modelos"
	"Autonomo_1/internal/servicios"
)

// ── HELPERS ──────────────────────────────────────────────────────

func cors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func jsonOK(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func jsonErr(w http.ResponseWriter, status int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{"error": msg})
}

// Extrae el ID del path: /api/productos/5  →  5
func idDesdeURL(r *http.Request, base string) (int, error) {
	parte := strings.TrimPrefix(r.URL.Path, base)
	parte  = strings.Trim(parte, "/")
	return strconv.Atoi(parte)
}

// ── PRODUCTOS ────────────────────────────────────────────────────
// GET  /api/productos        → listar        (endpoint 1)
// POST /api/productos        → agregar       (endpoint 2)
// PUT  /api/productos/{id}   → actualizar    (endpoint 9)
// DELETE /api/productos/{id} → eliminar      (endpoint 7)

func handlerProductos(w http.ResponseWriter, r *http.Request) {
	cors(w)
	if r.Method == http.MethodOptions {
		return
	}

	tieneID := strings.TrimPrefix(r.URL.Path, "/api/productos") != ""

	if tieneID {
		id, err := idDesdeURL(r, "/api/productos/")
		if err != nil {
			jsonErr(w, 400, "ID inválido")
			return
		}

		switch r.Method {

		// Endpoint 9: PUT /api/productos/{id}
		case http.MethodPut:
			var body struct {
				Precio float64 `json:"precio"`
				Stock  int     `json:"stock"`
			}
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				jsonErr(w, 400, "JSON inválido")
				return
			}
			if err := servicios.ActualizarProducto(id, body.Precio, body.Stock); err != nil {
				jsonErr(w, 500, err.Error())
				return
			}
			jsonOK(w, map[string]string{"mensaje": "Producto actualizado"})

		// Endpoint 7: DELETE /api/productos/{id}
		case http.MethodDelete:
			if err := servicios.EliminarProducto(id); err != nil {
				jsonErr(w, 500, err.Error())
				return
			}
			jsonOK(w, map[string]string{"mensaje": "Producto eliminado"})

		default:
			jsonErr(w, 405, "Método no permitido")
		}
		return
	}

	switch r.Method {

	// Endpoint 1: GET /api/productos
	case http.MethodGet:
		productos, err := servicios.ListarProductos()
		if err != nil {
			jsonErr(w, 500, err.Error())
			return
		}
		if productos == nil {
			productos = []modelos.Producto{}
		}
		jsonOK(w, productos)

	// Endpoint 2: POST /api/productos
	case http.MethodPost:
		var body struct {
			Nombre string  `json:"nombre"`
			Precio float64 `json:"precio"`
			Stock  int     `json:"stock"`
		}
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			jsonErr(w, 400, "JSON inválido")
			return
		}
		if strings.TrimSpace(body.Nombre) == "" {
			jsonErr(w, 400, "El nombre es obligatorio")
			return
		}
		if err := servicios.AgregarProducto(body.Nombre, body.Precio, body.Stock); err != nil {
			jsonErr(w, 500, err.Error())
			return
		}
		w.WriteHeader(http.StatusCreated)
		jsonOK(w, map[string]string{"mensaje": "Producto agregado"})

	default:
		jsonErr(w, 405, "Método no permitido")
	}
}

// ── USUARIOS ─────────────────────────────────────────────────────
// GET  /api/usuarios → listar    (endpoint 3)
// POST /api/usuarios → registrar (endpoint 4)

func handlerUsuarios(w http.ResponseWriter, r *http.Request) {
	cors(w)
	if r.Method == http.MethodOptions {
		return
	}

	switch r.Method {

	// Endpoint 3: GET /api/usuarios
	case http.MethodGet:
		usuarios, err := servicios.ListarUsuarios()
		if err != nil {
			jsonErr(w, 500, err.Error())
			return
		}
		if usuarios == nil {
			usuarios = []modelos.Usuario{}
		}
		jsonOK(w, usuarios)

	// Endpoint 4: POST /api/usuarios
	case http.MethodPost:
		var body struct {
			Nombre string `json:"nombre"`
			Correo string `json:"correo"`
		}
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			jsonErr(w, 400, "JSON inválido")
			return
		}
		if strings.TrimSpace(body.Nombre) == "" || strings.TrimSpace(body.Correo) == "" {
			jsonErr(w, 400, "Nombre y correo son obligatorios")
			return
		}
		if err := servicios.RegistrarUsuario(body.Nombre, body.Correo); err != nil {
			jsonErr(w, 500, err.Error())
			return
		}
		w.WriteHeader(http.StatusCreated)
		jsonOK(w, map[string]string{"mensaje": "Usuario registrado"})

	default:
		jsonErr(w, 405, "Método no permitido")
	}
}

// ── PEDIDOS ──────────────────────────────────────────────────────
// GET  /api/pedidos       → listar   (endpoint 5)
// POST /api/pedidos       → crear    (endpoint 6)
// GET  /api/pedidos/{id}  → detalle  (endpoint 8)

func handlerPedidos(w http.ResponseWriter, r *http.Request) {
	cors(w)
	if r.Method == http.MethodOptions {
		return
	}

	tieneID := strings.TrimPrefix(r.URL.Path, "/api/pedidos") != ""

	if tieneID {
		id, err := idDesdeURL(r, "/api/pedidos/")
		if err != nil {
			jsonErr(w, 400, "ID inválido")
			return
		}

		// Endpoint 8: GET /api/pedidos/{id}
		if r.Method == http.MethodGet {
			pedido, err := servicios.ObtenerPedido(id)
			if err != nil {
				jsonErr(w, 404, "Pedido no encontrado")
				return
			}
			jsonOK(w, pedido)
			return
		}

		jsonErr(w, 405, "Método no permitido")
		return
	}

	switch r.Method {

	// Endpoint 5: GET /api/pedidos
	case http.MethodGet:
		pedidos, err := servicios.ListarPedidos()
		if err != nil {
			jsonErr(w, 500, err.Error())
			return
		}
		if pedidos == nil {
			pedidos = []modelos.Pedido{}
		}
		jsonOK(w, pedidos)

	// Endpoint 6: POST /api/pedidos
	case http.MethodPost:
		var body struct {
			UsuarioID int                     `json:"usuario_id"`
			Items     []servicios.ItemEntrada `json:"items"`
		}
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			jsonErr(w, 400, "JSON inválido")
			return
		}
		if body.UsuarioID <= 0 {
			jsonErr(w, 400, "ID de usuario inválido")
			return
		}
		if len(body.Items) == 0 {
			jsonErr(w, 400, "Debe agregar al menos un producto")
			return
		}
		if err := servicios.CrearPedido(body.UsuarioID, body.Items); err != nil {
			jsonErr(w, 500, err.Error())
			return
		}
		w.WriteHeader(http.StatusCreated)
		jsonOK(w, map[string]string{"mensaje": "Pedido creado"})

	default:
		jsonErr(w, 405, "Método no permitido")
	}
}

// ── REPORTES ─────────────────────────────────────────────────────
// Endpoint 10: GET /api/reportes/ventas

func handlerReportes(w http.ResponseWriter, r *http.Request) {
	cors(w)
	if r.Method == http.MethodOptions {
		return
	}
	if r.Method != http.MethodGet {
		jsonErr(w, 405, "Método no permitido")
		return
	}

	reporte, err := servicios.ReporteVentas()
	if err != nil {
		jsonErr(w, 500, err.Error())
		return
	}
	if reporte == nil {
		reporte = []modelos.ReporteVenta{}
	}
	jsonOK(w, reporte)
}

// ── MAIN ─────────────────────────────────────────────────────────

func main() {
	database.Inicializar()

	http.HandleFunc("/api/productos/", handlerProductos)
	http.HandleFunc("/api/productos",  handlerProductos)
	http.HandleFunc("/api/usuarios",   handlerUsuarios)
	http.HandleFunc("/api/pedidos/",   handlerPedidos)
	http.HandleFunc("/api/pedidos",    handlerPedidos)
	http.HandleFunc("/api/reportes/ventas", handlerReportes)
	http.Handle("/", http.FileServer(http.Dir("./web")))

	fmt.Println("========================")
	fmt.Println("  SISTEMA E-COMMERCE")
	fmt.Println("========================")
	fmt.Println("Base de datos : ecommerce.db")
	fmt.Println("Servidor en   : http://localhost:8080")
	fmt.Println("")
	fmt.Println("Endpoints disponibles:")
	fmt.Println("  1. GET    /api/productos")
	fmt.Println("  2. POST   /api/productos")
	fmt.Println("  3. GET    /api/usuarios")
	fmt.Println("  4. POST   /api/usuarios")
	fmt.Println("  5. GET    /api/pedidos")
	fmt.Println("  6. POST   /api/pedidos")
	fmt.Println("  7. DELETE /api/productos/{id}")
	fmt.Println("  8. GET    /api/pedidos/{id}")
	fmt.Println("  9. PUT    /api/productos/{id}")
	fmt.Println(" 10. GET    /api/reportes/ventas")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error:", err)
	}
}
