package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func Inicializar() {
	var err error
	DB, err = sql.Open("sqlite3", "./ecommerce.db")
	if err != nil {
		log.Fatal("Error al abrir la base de datos:", err)
	}

	crearTablas()
}

func crearTablas() {
	sentencias := []string{
		`CREATE TABLE IF NOT EXISTS productos (
			id     INTEGER PRIMARY KEY AUTOINCREMENT,
			nombre TEXT    NOT NULL,
			precio REAL    NOT NULL,
			stock  INTEGER NOT NULL
		);`,

		`CREATE TABLE IF NOT EXISTS usuarios (
			id     INTEGER PRIMARY KEY AUTOINCREMENT,
			nombre TEXT NOT NULL,
			correo TEXT NOT NULL
		);`,

		`CREATE TABLE IF NOT EXISTS pedidos (
			id          INTEGER PRIMARY KEY AUTOINCREMENT,
			usuario_id  INTEGER NOT NULL,
			total       REAL    NOT NULL
		);`,

		`CREATE TABLE IF NOT EXISTS pedido_items (
			id          INTEGER PRIMARY KEY AUTOINCREMENT,
			pedido_id   INTEGER NOT NULL,
			producto_id INTEGER NOT NULL,
			cantidad    INTEGER NOT NULL,
			subtotal    REAL    NOT NULL
		);`,
	}

	for _, s := range sentencias {
		if _, err := DB.Exec(s); err != nil {
			log.Fatal("Error al crear tabla:", err)
		}
	}
}
