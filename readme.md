# 🛒 Sistema de Gestión E-Commerce en Go

Sistema de gestión de comercio electrónico desarrollado en **Go**, con una API REST de 10 endpoints, base de datos **SQLite** para persistencia de datos e interfaz web con **POO en JavaScript**.

Desarrollado como proyecto final de la asignatura **Programación Orientada a Objetos** — UIDE.

---

## 🚀 Tecnologías utilizadas

| Capa | Tecnología |
|------|-----------|
| Backend | Go (net/http, encoding/json) |
| Base de datos | SQLite (go-sqlite3 v1.14.22) |
| Frontend | HTML + CSS + JavaScript (POO) |
| Cliente DB | DBeaver |

---

## 📁 Estructura del proyecto

```
Autonomo_1/
├── cmd/
│   └── main.go               # Punto de entrada, servidor HTTP y rutas
├── internal/
│   ├── database/
│   │   └── db.go             # Conexión SQLite y creación de tablas
│   ├── modelos/
│   │   └── modelos.go        # Structs: Producto, Usuario, Pedido, ReporteVenta
│   ├── repositorios/
│   │   ├── producto_rep.go   # CRUD de productos en SQLite
│   │   ├── usuario_rep.go    # CRUD de usuarios en SQLite
│   │   ├── pedido_rep.go     # CRUD de pedidos con JOIN
│   │   └── reporte_rep.go    # Consultas agregadas de ventas
│   └── servicios/
│       └── servicios.go      # Lógica de negocio y validaciones
├── web/
│   └── index.html            # Interfaz web con POO en JavaScript
├── ecommerce.db              # Base de datos SQLite (generada automáticamente)
└── go.mod                    # Módulo Go y dependencias
```

---

## 📡 Endpoints de la API REST

Todos los endpoints reciben y devuelven datos en formato **JSON**.

| # | Método | Ruta | Descripción |
|---|--------|------|-------------|
| 1 | GET | `/api/productos` | Listar todos los productos |
| 2 | POST | `/api/productos` | Agregar un nuevo producto |
| 3 | GET | `/api/usuarios` | Listar todos los usuarios |
| 4 | POST | `/api/usuarios` | Registrar un nuevo usuario |
| 5 | GET | `/api/pedidos` | Listar todos los pedidos |
| 6 | POST | `/api/pedidos` | Crear un nuevo pedido |
| 7 | DELETE | `/api/productos/{id}` | Eliminar un producto por ID |
| 8 | GET | `/api/pedidos/{id}` | Ver detalle de un pedido |
| 9 | PUT | `/api/productos/{id}` | Actualizar precio y stock |
| 10 | GET | `/api/reportes/ventas` | Reporte de ventas por producto |

---

## 🗄️ Base de datos

El sistema genera automáticamente el archivo `ecommerce.db` con las siguientes tablas:

```sql
productos     → id, nombre, precio, stock
usuarios      → id, nombre, correo
pedidos       → id, usuario_id, total
pedido_items  → id, pedido_id, producto_id, cantidad, subtotal
```

---

## 🖥️ Clases POO en JavaScript (Frontend)

| Clase | Responsabilidad |
|-------|----------------|
| `Producto` | Formateo de precio, generación de filas HTML y opciones de selector |
| `Usuario` | Generación de filas HTML y opciones de selector |
| `ItemPedido` | Cálculo de subtotal, renderizado en el formulario de pedido |
| `Pedido` | Formateo de total, resumen de items, generación de filas HTML |
| `ReporteVenta` | Cálculo de participación porcentual y barra visual |

---

## ⚙️ Requisitos previos

Antes de ejecutar el proyecto asegúrate de tener instalado:

- [Go 1.21+](https://golang.org/dl/)
- [TDM-GCC](https://jmeubank.github.io/tdm-gcc/) (requerido por el driver go-sqlite3 en Windows)

Verifica la instalación de GCC:
```bash
gcc --version
```

---

## ▶️ Instalación y ejecución

**1. Clona el repositorio:**
```bash
git clone https://github.com/josstin-hue/Autonomo_1.git
cd Autonomo_1
```

**2. Descarga las dependencias:**
```bash
go mod tidy
```

**3. Ejecuta el servidor:**
```bash
go run ./cmd/main.go
```

**4. Abre el navegador en:**
```
http://localhost:8080
```

El archivo `ecommerce.db` se crea automáticamente en la raíz del proyecto la primera vez que se inicia el servidor.

---

## 🗃️ Visualizar la base de datos con DBeaver

1. Abre **DBeaver** → Nueva conexión → **SQLite**
2. En *Path*, selecciona el archivo:
   ```
   ruta_del_proyecto\Autonomo_1\ecommerce.db
   ```
3. Conéctate y expande **Tables** para ver los datos

Para ver todos los pedidos con nombres reales ejecuta en el Editor SQL:
```sql
SELECT
    p.id         AS pedido_id,
    u.nombre     AS cliente,
    pr.nombre    AS producto,
    pi.cantidad,
    pi.subtotal,
    p.total
FROM pedidos p
JOIN usuarios u      ON u.id  = p.usuario_id
JOIN pedido_items pi ON pi.pedido_id = p.id
JOIN productos pr    ON pr.id = pi.producto_id
ORDER BY p.id DESC;
```

---

## 👤 Autor

**Josstin Armando Proaño Pazmiño**  
Ingeniería en Sistemas de la Información — UIDE  
Asignatura: Programación Orientada a Objetos

---

## 📎 Enlaces

- 🎥 Video YouTube: https://youtu.be/TjAMw9V4cv4
- 💻 Repositorio GitHub: https://github.com/josstin-hue/Autonomo_1
