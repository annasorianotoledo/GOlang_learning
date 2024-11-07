### VARIABLES ###

Declaramos una variable de la siguiente manera:

var firstName string

GO es un lenguaje fuertemente tipado así que siempre al declarar una variable tenemos que especificar el tipo. Podemos declarar dos variables a la vez si son del mismo tipo, por ejemplo:
var firstName, lastName string


Podemos declarar las variables como hemos hecho antes o también podemos declararlas e inicializarlas (darles valor). Si queremos declarar múltiples variables con diferentes tipos podemos hacer lo siguiente:

var (
    firstName string = "John"
    lastName  string = "Doe"
    age       int    = 32
)


Si las inicializamos podemos ignorar el paso de escribir el tipo porque GO lo infiere al inicializarla con un valor. Por ejemplo: var firstName = "Anna"
Go sabrá que firstName es tipo string porque le hemos dado valor.

Este sería el método más habitual a la hora de declarar variables: 

package main

import "fmt"

func main() {
    firstName, lastName := "John", "Doe"
    age := 32
    fmt.Println(firstName, lastName, age)
}

El paquete fmt se utiliza para utilizar el método PrintIn.

!! Cuando se usan los dos puntos con el signo igual, la variable que se está declarando tiene que ser una nueva. Si usa dos puntos con un signo igual para una variable que ya se ha declarado, el programa no se compilará.
Sólo se puede utilizar := DENTRO DE UNA FUNCIÓN, para declarar variables fuera de una función, tenemos que utilizar la palabra var !!

## DECLARACIÓN DE CONSTANTES

const HTTPStatusOK = 200

Al igual que con var, si no le especificamos un tipo pero le damos un valor, GO infiere el tipo. Normalmente en GO las constantes se escriben con una combinación de mayúsculas y minúsculas o con todas las letras en mayúscula.

Para declarar un bloque de constantes que son secuenciales se utiliza iota, más info aquí https://go.dev/wiki/Iota

Se pueden declarar constantes sin usarlas y no se recibirá mensaje de error. No se pueden declarar con :=


## ERRORES DECLARACIÓN Y NO UTILIZACIÓN DE VARIABLES

Cuando generamos una variable y no la usamos, GO genera un error y no una advertencia. 

## TIPOS DE DATOS BÁSICOS

En Go hay cuatro categorías de tipos de datos:

Tipos básicos: números, cadenas y booleanos
Tipos agregados: matrices y estructuras
Tipos de referencia: punteros, segmentos, mapas, funciones y canales
Tipos de interfaz: interfaz

# NÚMEROS ENTEROS

La palabra clave para definir tipo entero es int. Pero están las opciones de int8, int16, int32 y int64; que son los enteros con un tamaño de 8 a 64 bits respectivamente.

# NÚMEROS FLOTANTES

Go proporciona tipos de datos para dos tamaños de números de punto flotante: float32 y float64. 

# VALORES BOOLEAN

Sólo tienen dos valores posibles: true y false. Se declaran mediante la palabra bool. No se puede convertir un tipo booleano a 0 o 1, se debe hacer explícitamente.

# STRINGS

Siempre se tendrán que definir con comillas dobles ("), las simples (') se usan para los caracteres individuales y para los runes.

A veces será necesario utilizar caracteres de escape, como los siguientes:

\n para líneas nuevas
\r para retornos de carro
\t para pestañas
\' para comillas simples
\" para comillas dobles
\\ para barras diagonales inversas

# VALORES PREDETERMINADOS PARA VARIABLES DECLARADAS PERO SIN VALOR

0 para los tipos int (y todos sus subtipos, como int64)
+0.000000e+000 para los tipos float32 y float64
false para los tipos bool
Un valor vacío para los tipos string

# CONVERSIÓN DE TIPOS

La conversión debe llevarse a cabo explícitamente. También podemos hacer uso del paquete 'strconv' para convertir un tipo string en un tipo int y viceversa.

Por ejemplo:

package main

import (
    "fmt"
    "strconv"
)

func main() {
    i, _ := strconv.Atoi("-42")
    s := strconv.Itoa(-42)
    fmt.Println(i, s)
}

El (_) que se usa como el nombre de una variable en el código anterior. En GO, el objeto _ significa que no vamos a usar el valor de esa variable y que queremos omitirlo. De lo contrario, el programa no se compilará porque necesitamos usar todas las variables declaradas. 

### FUNCIONES ###

Agrupa un conjunto de instrucciones a las que se puede llamar desde otras partes de la app. 

# FUNCIÓN MAIN

Todos los programas ejecutables que se usan en GO tienen esta función porque es el punto inicial del programa. Sólo puede haber una función main () en el programa. 

La función main no tiene ningún parámetro y no devuelve nada. Se puede acceder a los argumentos de la línea de comandos en GO. con el paquete os y la variable os.Args que contiene los argumentos que se pasan al programa: 
El siguiente código lee los dos números desde la línea de comandos y los suma:

package main

import (
    "fmt"
    "os"
    "strconv"
)

func main() {
    number1, _ := strconv.Atoi(os.Args[1])
    number2, _ := strconv.Atoi(os.Args[2])
    fmt.Println("Sum:", number1+number2)
}

Las variables están vacías pero utilizamos el _ para que no nos de error a la hora de compilar el código y en la línea de comandos le pasamos los argumentos que serán los valores de las dos variables.

!!IMPORTANTE!! La variable os.Args contiene todos los argumentos de la línea de comandos que se pasan al programa. Estos valores son de tipo string, por lo que debe convertirlos en int para sumarlos.

# FUNCIONES PERSONALIZADAS

Sintaxis para crear una función:

func name(parameters) (results) {
    body-content
}

 Puede haber cero o más parámetros. Además, se pueden definir los tipos de valor devueltos de la función, que también pueden ser cero o más.

package main

import (
    "fmt"
    "os"
    "strconv"
)

func main() {
    sum := sum(os.Args[1], os.Args[2])
    fmt.Println("Sum:", sum)
}

func sum(number1 string, number2 string) int {
    int1, _ := strconv.Atoi(number1)
    int2, _ := strconv.Atoi(number2)
    return int1 + int2
}

La función sum, toma dos args de string y los convierte en int, después devuelve el resultado de la suma. Al definir un tipo de valor devuelto, la función tiene que devolver dicho tipo.

Es posible establecer un nombre para el valor devuelto de una función, como si fuera una variable. Por ejemplo, la función sum se podría refactorizar de la siguiente manera:

func sum(number1 string, number2 string) (result int) {
    int1, _ := strconv.Atoi(number1)
    int2, _ := strconv.Atoi(number2)
    result = int1 + int2
    return
}

# DEVOLUCIÓN DE VARIOS VALORES EN UNA FUNCIÓN

package main

import (
    "fmt"
    "os"
    "strconv"
)

func main() {
    sum, mul := calc(os.Args[1], os.Args[2])
    fmt.Println("Sum:", sum)
    fmt.Println("Mul:", mul)
}

func calc(number1 string, number2 string) (sum int, mul int) {
    int1, _ := strconv.Atoi(number1)
    int2, _ := strconv.Atoi(number2)
    sum = int1 + int2
    mul = int1 * int2
    return
}

El anterior código necesita dos variables para almacenar los resultados de la función, sino no se compilará.

Si no necesita uno de los valores que devuelve una función, puede descartarlo asignando el valor devuelto a la variable _.

# CAMBIO DE LOS VALORES DE LOS PARÁMETROS DE LA FUNCIÓN (PUNTEROS)

GO es un lenguaje de programación de "paso por valor", cada vez que se pasa un valor a una función, GO toma ese valor y crea una copia local (nueva variable en memoria), los cambios que realice en esa variable en la función no afectarán a la variable que haya enviado a la función.

package main

import "fmt"

func main() {
    firstName := "John"
    updateName(firstName)
    fmt.Println(firstName)
}

func updateName(name string) {
    name = "David"
}

El resultado es John, aunque en la función updateName se modifique a David. La función updateName modifica sólo la copia local. GO pasa el valor de la variable, no la propia variable.

Si queremos que la función updateName pueda cambiar la variable, tenemos que utilizar un puntero. Un puntero es una variable que contiene la dirección de memoria de otra variable. Cuando se envía un puntero a una función, no se pasa un valor, sino una dirección de memoria. Todos los cambios que se realicen en esa variable, afectarán al autor de la llamada.

Hay dos tipos de punteros
- El operador & toma la dirección del objeto que le sigue.
- El operador * desreferencia un puntero. Le proporciona acceso al objeto en la dirección que contiene el puntero.

Por ejemplo, en el siguiente código vemos la utilización de esos punteros para cambiar el valor a una variable a través de una función.

package main

import "fmt"

// Función que recibe un puntero a un int y modifica su valor
func changeValue(num *int) {
    *num = 42  // Modifica el valor apuntado por num
}

func main() {
    x := 10
    fmt.Println("Antes de la función:", x)
    changeValue(&x)  // Pasa la dirección de x a la función
    fmt.Println("Después de la función:", x)
}

Estas operaciones permiten trabajar directamente con la memoria y permiten que las funciones modifiquen los valores de las variables originales, lo cual es esencial para optimizar el uso de recursos y manipular estructuras de datos complejas.

### PAQUETES ###

