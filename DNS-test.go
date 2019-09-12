package  main

import (
        "fmt"
        "net"
        //"time"
        "context"
        //"google.golang.org/grpc"
)

func dialContext(ctx context.Context, network, address string) (net.Conn, error) {
        return (&net.Dialer{}).DialContext(ctx, network, address)
}

func main(){
        var ctx context.Context
        var network,address string
        network = "tcp"
        //address = "www.github.com:443"
        address = "127.0.0.1:22"
        ctx = context.Background()
        //ctx, cancel := context.WithTimeout(context.Background(), time.Second)
        //defer cancel()
        fmt.Print("ctx = ", ctx)
        fmt.Println(" ")
        fmt.Print("network = ", network)
        fmt.Println(" ")
        fmt.Print("address = ", address)
        fmt.Println(" ")

        _,err := dialContext(ctx, network, address)
        if(err != nil)  {
                fmt.Print("err = ", err)
                fmt.Println(" ")
                fmt.Println("dialContext error!\n")
        }
  
        return;
}
