# Taller de programación imperativa en Go

**Curso:** Paradigmas de Programación - TEC  
**Estudiante:** [Luis Fernando Mejías Mejías]  
**Lenguaje:** Go

## Cómo correr el programa

Se necesita tener Go instalado. Desde la terminal, dentro de la carpeta del taller, se ejecuta:

```bash
go run main.go
```

Los datos de prueba están puestos directamente en `main`. El programa muestra los resultados de los ocho ejercicios.

## Pruebas realizadas

**Ejercicios 1 al 4**

| Ejercicio            | Resultado                                        |
| -------------------- | ------------------------------------------------ |
| 1. Intercambio       | `8, -3` pasa a `-3, 8`                           |
| 2. Estadísticas      | Suma `4`, mínimo `-5`, máximo `7`, negativos `2` |
| 2. Recorrido inverso | `-5 7 0 -2 4`                                    |
| 3. Copia             | Original `[4 -2 0 7 -5]`, copia `[99 -2 0 7 -5]` |
| 4. Ordenamiento      | `[-5 -2 0 4 7]`                                  |

**Ejercicio 5: búsqueda de pares**

| Arreglo         | Objetivo | Resultado |
| --------------- | -------- | --------- |
| `[-4 -1 0 2 5]` | `1`      | `(-4, 5)` |
| `[-4 -1 0 2 5]` | `8`      | No existe |
| `[0 0]`         | `0`      | `(0, 0)`  |
| `[3]`           | `6`      | No existe |

**Ejercicio 6: 3Sum**

| Arreglo            | Resultado                              |
| ------------------ | -------------------------------------- |
| `[-1 0 1 2 -1 -4]` | `(-1, -1, 2)` y `(-1, 0, 1)`           |
| `[0 1 1]`          | Ninguna terna                          |
| `[0 0 0]`          | Una terna: `(0, 0, 0)`                 |
| `[0 0 0 0]`        | Una terna: `(0, 0, 0)`                 |
| `[-2 0 0 2 2]`     | Una terna: `(-2, 0, 2)`                |
| `[0 0]`            | Ninguna terna, porque faltan elementos |

**Ejercicio 7: máximos acumulados**

Con las alturas `[3 0 2 0 4]` se obtuvo:

- Izquierda: `[3 3 3 3 4]`
- Derecha: `[4 4 4 4 4]`

**Ejercicio 8: agua atrapada**

| Alturas                     | Agua total        |
| --------------------------- | ----------------- |
| `[5]`                       | `0`               |
| `[0 1 2 3]`                 | `0`               |
| `[3 2 1 0]`                 | `0`               |
| `[2 2 2]`                   | `0`               |
| `[3 0 3]`                   | `3`               |
| `[3 0 2 0 4]`               | `7`               |
| `[4 2 0 3 2 5]`             | `9`               |
| `[0 1 0 2 1 0 1 3 2 1 2 1]` | `6`               |
| `[4 1 3 0 4]`               | `8` (caso propio) |

En el caso propio se eligió `[4 1 3 0 4]`. Las columnas de los extremos miden 4, entonces en el centro se acumulan 3, 1 y 4 unidades de agua. En total son 8.

También se muestra el agua en cada posición para `[3 0 2 0 4]`:

```text
Alturas:       [3 0 2 0 4]
Max izquierda: [3 3 3 3 4]
Max derecha:   [4 4 4 4 4]
Agua:          [0 3 1 3 0]
```

## Explicaciones

### Punteros y variables locales

En `swap` se usan punteros para cambiar los valores originales de `a` y `b`. Al escribir `*a` se modifica el número al que apunta. Si se cambiara solamente `a`, se cambiaría el puntero que recibe la función, no el número de `main`.

La variable `aux` se crea dentro de `swap`, así que solamente se puede usar ahí. Cuando la función termina, ya no se puede acceder a ella desde `main`.

### Copia de un slice

Para hacer la copia se usa `make` y un ciclo que pasa los números uno por uno. Si se hiciera `otro := original`, los dos slices compartirían los mismos elementos y cambiar la copia también podría cambiar el original.

### Búsqueda de pares

Como los números ya están ordenados, se empieza por los dos extremos. Si la suma es muy pequeña, se mueve el índice izquierdo para probar un número mayor. Si es muy grande, se mueve el derecho para probar uno menor. Así se pueden descartar valores que no sirven.

### Ternas repetidas en 3Sum

Primero se ordena una copia del arreglo. Después se toma un número y se buscan los otros dos desde los extremos. Se saltan los números repetidos para que una misma terna no salga varias veces. Cuando se encuentra una terna, la búsqueda sigue para revisar las demás.

### Cálculo del agua

El agua depende del menor de los máximos de izquierda y derecha, porque si un lado es más bajo, el agua se escaparía por ahí. Por eso se toma el menor de los dos y se le resta la altura de la posición actual.

### Memoria en Go y C

En C se usa `malloc` para reservar memoria y `free` para liberarla. También hay que comprobar que la reserva haya funcionado. En Go se usa `make` para crear los slices y no hace falta liberar esa memoria manualmente, porque Go cuenta con un recolector de basura. Aun así, hay que tener cuidado con los slices que comparten datos.
