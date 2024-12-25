# Golang - Proyecto Gestión de Usuarios

Este es un proyecto desarrollado como parte de un curso de Golang. Permite gestionar usuarios mediante un CRUD básico (Crear, Listar, Actualizar y Eliminar), simulando un sistema de administración de usuarios.

## Funcionalidades

1. **Crear Usuario**: Registra un nuevo usuario solicitando información como nombre, edad, email y dirección.
2. **Listar Usuarios**: Muestra una lista completa de los usuarios registrados.
3. **Actualizar Usuario**: Permite modificar los datos de un usuario existente seleccionándolo por su ID.
4. **Eliminar Usuario**: Borra un usuario del sistema utilizando su ID.
5. **Limpiar Consola**: Limpia la pantalla para facilitar la lectura.

## Requisitos Previos

- Go 1.20 o superior instalado en tu máquina.
- Un terminal o IDE que permita ejecutar programas de Golang.
- Conocimientos básicos de programación en Go.

## Estructura del Código

El proyecto se estructura de la siguiente manera:
- **`Usuarios`**: Define la estructura del usuario con atributos como `id`, `nombre`, `edad`, `email` y `direccion`.
- **Mapas**: Utiliza un mapa (`map[int]Usuarios`) para almacenar los usuarios registrados.
- **Funciones CRUD**:
  - `crearUsuario()`: Solicita datos al usuario y lo registra en el mapa.
  - `listarUsuarios()`: Itera sobre los registros y los imprime en consola.
  - `actualizarUsuarios()`: Modifica los datos de un usuario específico por ID.
  - `eliminarUsuarios()`: Elimina un usuario del mapa por su ID.

## Instalación y Ejecución

1. Clona este repositorio en tu máquina local:
   ```
   bash
   git clone https://github.com/anidroid1184/Golang-Proyecto-Gestion-Usuarios.git
   cd Golang-Proyecto-Gestion-Usuarios
   ```
2. Asegúrate de que tienes Golang instalado. Puedes verificarlo con:
  ```
  bash
  go version
  ```

3. Ejecuta el archivo principal:
  ```
  bash
  go run main.go
  ```

4. Sigue las instrucciones en pantalla para interactuar con el sistema.
Ejemplo de Uso
Seleccione por favor una de las siguientes opciones:
    A) Crear
    B) Listar
    C) Actualizar
    D) Eliminar
    E) Limpiar consola
Presione quit, q, esc o wq para salir
Crear un Usuario:

Opción: A
Introduce los datos solicitados (nombre, edad, email, dirección).
Listar Usuarios:

Opción: B
Visualiza los usuarios registrados con su ID.
Actualizar Usuario:

Opción: C
Selecciona un ID y elige el parámetro a modificar.
Eliminar Usuario:

Opción: D
Introduce el ID del usuario que deseas eliminar.

Autor
Creado por Juan Sebastián Valencia Londoño como parte del aprendizaje de Golang.

Licencia
Este proyecto está bajo la Licencia MIT. Puedes consultar el archivo LICENSE para más detalles.


Este README proporciona toda la información necesaria para comprender, instalar y usar el proyecto, además de seguir las mejores prácticas de documentación.
