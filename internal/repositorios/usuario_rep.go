package repositorios

import (
	"Autonomo_1/internal/database"
	"Autonomo_1/internal/modelos"
)

func RegistrarUsuario(nombre, correo string) error {
	_, err := database.DB.Exec(
		`INSERT INTO usuarios (nombre, correo) VALUES (?, ?)`,
		nombre, correo,
	)
	return err
}

func ListarUsuarios() ([]modelos.Usuario, error) {
	rows, err := database.DB.Query(`SELECT id, nombre, correo FROM usuarios`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var usuarios []modelos.Usuario
	for rows.Next() {
		var u modelos.Usuario
		if err := rows.Scan(&u.ID, &u.Nombre, &u.Correo); err != nil {
			return nil, err
		}
		usuarios = append(usuarios, u)
	}
	return usuarios, nil
}
