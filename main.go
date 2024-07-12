package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

// argumentos en ejecucion de programa
// func main() {
// 	nombre := flag.String("nombre", "", "El nombre de la persona")
// 	edad := flag.Int("edad", 18, "La edad de la persona")
// 	flag.Parse()
// 	fmt.Println("Tu nombre es: ", *nombre)
// 	fmt.Println("Tu edad es: ", *edad)

// }

func main() {

	// err := errors.New("Este es un error fatal de prueba")

	// fatal detiene la ejecucion del script
	// log.Fatal(err)
	// fmt.Println("hola soy otro texto")
	// fmt.Println("Starging apllication...")
	// log.Panic("error panic")

	// viene un error puedo pintarlo asi
	// err := "todo esta mal"
	// log.Panicf(err)

	f, err := os.OpenFile("logs.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)

	if err != nil { //si err es distitno de nil hay error por que err == err hay error si err es nulo o nil no hay error
		log.Fatal(err)
	}

	defer f.Close()
	log.SetOutput(f)
	log.Printf("Error linea %v", 1)
}

// math rand
// func main() {

// 	aleatorio := rand.Intn(101)
// 	fmt.Println(aleatorio)
// 	min := 1000
// 	max := 10000
// 	rand.Seed(time.Now().UnixNano())
// 	aleatorio2 := rand.Intn(max-min) + min
// 	fmt.Println(aleatorio2)

// 	password := generatePassword(20, 1,1,1)
// 	fmt.Println(password)

// }

var (
	lowerCharSet   = "abcdefghijklmnopqrstuvwxyz"
	upperCharSet   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialCharSet = "!@#$%&/()"
	numbeSet       = "0123456789"
	allCharSet     = lowerCharSet + upperCharSet + specialCharSet + numbeSet
)

func generatePassword(passwordLength, minSpecialChar, minNum, minUppercase int) string {
	var password strings.Builder

	//set especial character

	for i := 0; i < minSpecialChar; i++ {
		random := rand.Intn(len(specialCharSet))
		password.WriteString(string(specialCharSet[random]))
	}

	//set numeric

	for i := 0; i < minNum; i++ {
		random := rand.Intn(len(numbeSet))
		password.WriteString(string(numbeSet[random]))
	}
	//set uppercase

	for i := 0; i < minUppercase; i++ {
		random := rand.Intn(len(upperCharSet))
		password.WriteString(string(upperCharSet[random]))
	}

	remainingLength := passwordLength - minSpecialChar - minNum - minUppercase

	for i := 0; i < remainingLength; i++ {
		random := rand.Intn(len(allCharSet))
		password.WriteString(string(allCharSet[random]))
	}

	inRune := []rune(password.String())

	rand.Shuffle(len(inRune), func(i, j int) {
		inRune[i], inRune[j] = inRune[j], inRune[i]
	})
	return string(inRune)

}

// modulo strings

// func main() {
// 	cadena := "Mi muñeca me hablo"
// 	fmt.Println("to upper ", strings.ToUpper(cadena))
// 	fmt.Println("to upper ", strings.ToLower(cadena))
// 	letras := strings.Split(cadena, "")
// 	fmt.Println(letras)
// 	pos := strings.Index(cadena, "Mi")
// 	if pos == -1 {
// 		fmt.Println("No esta la palabra mi dentro de ", cadena)
// 	} else {
// 		fmt.Println("La palabra mi si esta dentro de ", cadena, "a partir de la posicion: ", pos)

// 	}
// 	imprimirRepetida := strings.Repeat(cadena, 2)
// 	fmt.Println(imprimirRepetida)
// 	cadenaNueva := strings.Replace(cadena, "Mi", "La", -1)
// 	fmt.Println(cadenaNueva)

// 	fmt.Println(string(cadenaNueva[5:9]))

// }

// modulo time
// func main() {

// 	fmt.Println("La hora actual es: ", time.Now())
// 	fecha := time.Now()
// 	fmt.Println("El mes es: ", fecha.Year())
// 	fmt.Println("El numero del mes es: ", int(fecha.Month()))
// 	fmt.Println("El dia es: ", fecha.Day())
// 	fmt.Println("La hora es: ", fecha.Hour())
// 	fmt.Println("los minutos son es: ", fecha.Minute())
// 	fmt.Println("Los segundos son es: ", fecha.Second())
// 	fmt.Printf("El tipo es: %T \n", fecha)

// 	fmt.Printf("%v/%v/%v \n", fecha.Day(), int(fecha.Month()), fecha.Year())
// 	fmt.Printf("%v %v de %v de %v a las %v:%v:%v \n", fecha.Weekday(), int(fecha.Day()), fecha.Month(), fecha.Year(), fecha.Hour(), fecha.Minute(), fecha.Second())
// 	ahora := time.Now()
// 	fmt.Println("Fecha en este momento:")
// 	fmt.Printf("%v/%v/%v \n", ahora.Day(), int(ahora.Month()), ahora.Year())
// 	fmt.Println("Ahora mas 20 dias: ")
// 	fecha1 := ahora.Add(time.Hour * 24 * 20)
// 	fmt.Printf("%v/%v/%v \n", fecha1.Day(), int(fecha1.Month()), fecha1.Year())
// 	fmt.Println("Fecha en este momento:")
// 	fmt.Printf("%v/%v/%v \n", ahora.Day(), int(ahora.Month()), ahora.Year())
// 	fmt.Println("Ahora menos 20 dias: ")
// 	fecha2 := ahora.Add(time.Hour * 24 * -20)
// 	fmt.Printf("%v/%v/%v \n", fecha2.Day(), int(fecha2.Month()), fecha2.Year())
// 	fmt.Println("Fecha en este momento:")
// 	fmt.Printf("%v/%v/%v \n", ahora.Day(), int(ahora.Month()), ahora.Year())
// 	fmt.Println("Fecha dentro de un año : ")
// 	fecha3 := ahora.Add(time.Hour * 24 * 365)
// 	fmt.Printf("%v/%v/%v \n", fecha3.Day(), int(fecha3.Month()), fecha3.Year())
// 	fmt.Println(FormatearFecha(ahora))
// }

func FormatearFecha(fecha time.Time) string {

	v := fmt.Sprintf("%v/%v/%v", fecha.Day(), int(fecha.Month()), fecha.Year())

	return v

}
