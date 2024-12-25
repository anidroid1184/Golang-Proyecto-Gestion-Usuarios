package main

import (
	"fmt"
	// pasar a lowercase
	"strings"
	// pasar de string a int
	"strconv"
	// para leer entradas
	"bufio"
	"os"
	// limpiar consola
	"os/exec"
)

type Usuarios struct {
	id        int
	nombre    string
	edad      int
	email     string
	direccion string
}

// creamos un mapa para crear los usuarios
var users map[int]Usuarios

/*func crearUsuario(nombre, email, direccion string, edad int) Usuarios {

}*/

var id int = 0000

func crearUsuario() {
	fmt.Println("Creando usuario")
	fmt.Println("Ingrese su nombre")
	nombre := readLine()
	fmt.Println("Ingrese su edad")
	edad, err := strconv.Atoi(readLine())
	for {
		// pasar de string a int
		if err != nil || 0 > edad || edad >= 100 {
			fmt.Println("Ingrese una edad valida")
			edad, err = strconv.Atoi(readLine())
		} else {
			break
		}
	}
	fmt.Println("Ingrese su email")
	email := readLine()
	fmt.Println("Ingrese su direccion")
	direccion := readLine()

	id++
	usuario := Usuarios{id, nombre, edad, email, direccion}

	// añadir usuario al map
	users[id] = usuario

	fmt.Printf("Usuarios creado:\nid:\t%d\nusuario:\t%v retornando al menu principal", id, users[id])
	fmt.Println("----------------------------")
}
func listarUsuarios() {
	fmt.Println("Listando usuarios")
	for i := 1; i <= id; i++ {
		fmt.Printf("\nId: #%d\tNombre: %s\tEdad: %d\tEmail: %s\tDireccion: %s\n", i, users[i].nombre, users[i].edad, users[i].email, users[i].direccion)
	}
	fmt.Println("----------------------------")
}
func actualizarUsuarios() {
	listarUsuarios()
	fmt.Println("Para realizar cambios en su usario, ingrese porfavor el id:")
	idUser, idErr := strconv.Atoi(readLine())
	for {
		// pasar de string a int
		if idErr != nil || idUser > id {
			fmt.Println("Ingrese un id valido")
			idUser, idErr = strconv.Atoi(readLine())
		} else {
			break
		}

	}

	fmt.Printf("\nIngrese el parametro que desea cambiar: \n\ta)Nombre\n\tb)edad\n\tc)email\n\td)direccion\n")
	parametro := readLine()

	switch parametro {
	case "a", "Nombre":
		fmt.Println("Ingrese el nuevo nombre")
		nuevoNombre := readLine()
		// hago una copia para modificar la entrada

		usuario := users[idUser]
		usuario.nombre = nuevoNombre
		users[idUser] = usuario

	case "b", "edad":
		fmt.Println("Ingrese la nueva edad")
		nuevaEdad, edadErr := strconv.Atoi(readLine())
		for {
			// pasar de string a int
			if edadErr != nil || nuevaEdad >= 100 || nuevaEdad < 0 {
				fmt.Println("Ingrese una edad valida")
				nuevaEdad, edadErr = strconv.Atoi(readLine())
			} else {
				break
			}
		}
		usuario := users[idUser]
		usuario.edad = nuevaEdad
		users[idUser] = usuario

	case "c", "email":
		fmt.Println("Ingrese el nuevo Email")
		nuevoEmail := readLine()

		usuario := users[idUser]
		usuario.email = nuevoEmail
		users[idUser] = usuario

	case "d", "direccion":
		fmt.Println("Ingrese una nueva direccion")
		nuevaDireccion := readLine()

		usuario := users[idUser]
		usuario.direccion = nuevaDireccion
		users[idUser] = usuario
	default:
		fmt.Printf("Ingrese una opción valida\n")
	}

	fmt.Println("Usuario Actualizado")
	fmt.Println("Nueva lista de usuarios: ")
	listarUsuarios()
}

func eliminarUsuarios() {
	listarUsuarios()
	fmt.Println("Para realizar eliminare su usario, ingrese porfavor el id:")
	idUser, idErr := strconv.Atoi(readLine())
	for {
		// pasar de string a int
		if idErr != nil || idUser > id {
			fmt.Println("Ingrese un id valido")
			idUser, idErr = strconv.Atoi(readLine())
		} else {
			break
		}

	}
	delete(users, idUser) // eliminar en el map
	fmt.Println("Eliminando usuario")
	fmt.Println("Nueva lista de usuarios: ")
	listarUsuarios()

}

var reader *bufio.Reader // lector de entradas

func readLine() string {
	// verificamos que no retorne error
	if option, err := reader.ReadString('\n'); err != nil {
		panic("No es posible obtener su respuesta")
	} else {
		// como nos retornara el string con un salto de linea se lo quitamos
		option = strings.TrimSuffix(option, "\n")
		// tambien volvemos mayuscula la entrada
		option = strings.ToLower(option)
		return option
	}
}

func clearConsole() {
	cmd := exec.Command("clear") // pasamos el comando que limpiara la consola
	cmd.Stdout = os.Stdout       // almacenamos el comando
	cmd.Run()                    // corremos el comando

}

func main() {
	var option string

	users = make(map[int]Usuarios)

	reader = bufio.NewReader(os.Stdin) // con esta variable podemos obtener las entradas
	// leer entrada
	for {
		fmt.Printf("Seleccione porfavor una de las siguientes opciones:\n\tA) Crear\n\tB) Listar\n\tC) Actualizar\n\tD) Eliminar\n\tE) Limpiar consola\nPresione quite, q, esc o wq para salir\n")
		// condicionamos para que las opciones se muestren continuamente
		option = readLine()

		if option == "quit" || option == "q" || option == "esc" || option == "wq" {
			break
		}

		switch option {
		case "a", "crear":
			crearUsuario()

		case "b", "listar":
			listarUsuarios()

		case "c", "actualizar":
			actualizarUsuarios()

		case "d", "eliminar":
			eliminarUsuarios()

		case "e", "limpiar consola", "limpiar":
			clearConsole()

		default:
			fmt.Printf("Ingrese una opción valida\n")
		}
	}

	fmt.Printf("\nGracias por usar nuestros servicios, adios\n")
}
