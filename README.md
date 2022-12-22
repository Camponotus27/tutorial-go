# Tutorial de go

## Como usar git

Para descargar el proyecto se debe ejecutar (crea una carpeta con el proyecto en la ubicacion actual del terminal)

```bash
    git clone https://github.com/Camponotus27/tutorial-go
```

Para descargar cambios nuevos del repositorio sin tener que clonarlo otra vez se usa (pull significa jalar)

```bash
    git pull
```

## Iniciar proyecto

Inicia el archivo de modulos (administra las dependencias o packetes)

Este comando crea el archivo `go.mod`

```bash
    go mod init
```

## Crear entrada de la aplicacion

Se debe crer un archivo en la raiz del proyecto (nombre cualquiera) con el package main, dentro de este archivo debe existir una funcion llamada main

Este comando ejecuta la aplicacion

```bash
    go run ./main.go
```
