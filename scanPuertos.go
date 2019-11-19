package main

import(
    "fmt"
    "net"
	"strconv"
	"time"
)
var c = make(chan int) 

func solicitud(ip string, min int, max int) {
	for min < max{
		conn, err := net.Dial("tcp", ip+":"+strconv.Itoa(min))
		min++
		if err != nil {
			continue
		}
		defer conn.Close()
		fmt.Println("puerto: ", min-1," Online ")
	}
	c <- 1
}

func main(){
	var (
		ip="google.com.mx"
		inicio=0
		fin=0
		i=1
	)
		fmt.Println("Analizando "+ip+", tardara un rato mas.... (15 min aprox)")
	t := time.Now()
	fecha := fmt.Sprintf("%02d:%02d:%02d",t.Hour(), t.Minute(), t.Second())
	fmt.Println("Inicio de analisis: "+fecha)
	inicio=0
	for i<=20{
		fin=inicio+3279
		go solicitud(ip,inicio,fin)
		inicio=fin
		i++
	}
	<- c
	t = time.Now()
	fechaFin := fmt.Sprintf("%02d:%02d:%02d",t.Hour(), t.Minute(), t.Second())
	fmt.Println("Fin de analisis :"+fechaFin)
}
