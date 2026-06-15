package servicios

import (
    "Autonomo_1/internal/modelos"
    "Autonomo_1/internal/repositorios"
)
func RegistrarUsuario(nombre string, correo string) {

	usuario := modelos.Usuario{
		ID:     len(repositorios.Usuarios) + 1,
		Nombre: nombre,
		Correo: correo,
	}

	repositorios.Usuarios = append(repositorios.Usuarios, usuario)
}

func ListarUsuarios() []modelos.Usuario {
	return repositorios.Usuarios
}