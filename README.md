# PAAC - Paquete de Saludos en Go

[![Go Version](https://img.shields.io/badge/go-%3E%3D1.24.4-blue.svg)](https://golang.org/)
[![License](https://img.shields.io/badge/license-MIT-green.svg)](LICENSE)

## 📖 Descripción

PAAC (Paquete de Saludos Aleatorios en Go) es una librería simple y útil desarrollada en Go que proporciona funciones para generar saludos personalizados con mensajes aleatorios en español. Es perfecto para aprender los fundamentos de Go o para agregar un toque amigable a tus aplicaciones.

## ✨ Características

- 🎲 **Mensajes aleatorios**: Genera saludos con mensajes aleatorios de una colección predefinida
- 👤 **Saludos individuales**: Saluda a una persona específica
- 👥 **Saludos múltiples**: Saluda a múltiples personas de una vez
- ✅ **Manejo de errores**: Validación de entrada con manejo adecuado de errores
- 🧪 **Probado**: Incluye tests unitarios
- 🚀 **Fácil de usar**: API simple e intuitiva

## 📦 Instalación

### Prerrequisitos

- Go 1.24.4 o superior

### Instalación como dependencia

```bash
go get github.com/mauricioloya/paac-golang
```

### Clonar el repositorio

```bash
git clone https://github.com/mauricioloya/paac-golang.git
cd paac-golang
```

## 🚀 Uso

### Importar el paquete

```go
import "github.com/mauricioloya/paac-golang"
```

### Saludo individual

```go
package main

import (
    "fmt"
    "log"
    "github.com/mauricioloya/paac-golang"
)

func main() {
    // Saludo a una persona
    mensaje, err := utils.Hello("Juan")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(mensaje)
    // Salida ejemplo: "Hola Juan, que seas bienvenido"
}
```

### Saludos múltiples

```go
package main

import (
    "fmt"
    "log"
    "github.com/mauricioloya/paac-golang"
)

func main() {
    // Saludo a múltiples personas
    nombres := []string{"Ana", "Carlos", "María"}
    mensajes, err := utils.Hellos(nombres)
    if err != nil {
        log.Fatal(err)
    }
    
    for nombre, mensaje := range mensajes {
        fmt.Printf("%s: %s\n", nombre, mensaje)
    }
    // Salida ejemplo:
    // Ana: Hola Ana, como estas?
    // Carlos: Hola Carlos, que tal?
    // María: Hola María, hasta que te dejas ver
}
```

## 📚 API Reference

### `Hello(name string) (string, error)`

Genera un saludo personalizado para una persona específica.

**Parámetros:**
- `name` (string): El nombre de la persona a saludar

**Retorna:**
- `string`: El mensaje de saludo generado
- `error`: Error si el nombre está vacío

**Ejemplo:**
```go
mensaje, err := utils.Hello("Pedro")
// Retorna: "Hola Pedro, hace cuanto kilos no nos vemos?"
```

### `Hellos(names []string) (map[string]string, error)`

Genera saludos personalizados para múltiples personas.

**Parámetros:**
- `names` ([]string): Slice con los nombres de las personas a saludar

**Retorna:**
- `map[string]string`: Mapa con nombres como llaves y mensajes como valores
- `error`: Error si alguno de los nombres está vacío

**Ejemplo:**
```go
nombres := []string{"Luis", "Sofia"}
mensajes, err := utils.Hellos(nombres)
// Retorna: map["Luis":"Hola Luis, que tal?" "Sofia":"Hola Sofia, que seas bienvenida"]
```

## 🎯 Mensajes Disponibles

El paquete incluye los siguientes mensajes aleatorios:

- "Hola %v, que seas bienvenido"
- "Hola %v, como estas?"
- "Hola %v, que tal?"
- "Hola %v, hasta que te dejas ver"
- "Hola %v, hace cuanto kilos no nos vemos?"

## 🧪 Ejecutar Tests

Para ejecutar los tests del proyecto:

```bash
# Ejecutar todos los tests
go test

# Ejecutar tests con detalle
go test -v

# Ejecutar tests con cobertura
go test -cover
```

## 🛠️ Desarrollo

### Estructura del proyecto

```
paac-golang/
├── go.mod          # Definición del módulo
├── utils.go        # Funciones principales
├── utils_test.go   # Tests unitarios
└── README.md       # Documentación
```

### Contribuir

1. Fork el proyecto
2. Crea una rama para tu feature (`git checkout -b feature/nueva-caracteristica`)
3. Commit tus cambios (`git commit -am 'Agregar nueva característica'`)
4. Push a la rama (`git push origin feature/nueva-caracteristica`)
5. Abre un Pull Request

## 📄 Licencia

Este proyecto está licenciado bajo la Licencia MIT - ver el archivo [LICENSE](LICENSE) para más detalles.

## 👨‍💻 Autor

**Mauricio Loya**
- GitHub: [@mauricioloya](https://github.com/mauricioloya)

## 🤝 Agradecimientos

Este proyecto fue desarrollado como parte del curso de Go en Udemy, enfocado en aprender las mejores prácticas de desarrollo en Go.

---

⭐ ¡No olvides dar una estrella al proyecto si te fue útil!