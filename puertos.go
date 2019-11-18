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
		conn, err := net.Dial("tcp", ip+":"+strconv.Itoa(min) )
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
	var ip="google.com.mx"
	
	t := time.Now()
	fecha := fmt.Sprintf("%02d:%02d:%02d",t.Hour(), t.Minute(), t.Second())
	fmt.Println(fecha)
	go solicitud(ip,0,13107)
	//go solicitud(ip,13107,26214)
	//go solicitud(ip,26214,39321)
	//go solicitud(ip,39321,52427)
	//go solicitud(ip,52427,65536)
	//go solicitud(ip,80,101)
	<- c
	fecha = fmt.Sprintf("%02d:%02d:%02d",t.Hour(), t.Minute(), t.Second())
	fmt.Println(fecha)
}
