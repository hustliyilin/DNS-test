package  main

import (
        "fmt"
        "net"
        "time"
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
        address = "peer:7052"
        ctx, cancel := context.WithTimeout(context.Background(), time.Second)
        defer cancel()
        fmt.Print("ctx = ", ctx)
        fmt.Println(" ")
        fmt.Print("network = ", network)
        fmt.Println(" ")
        fmt.Print("address = ", address)
        fmt.Println(" ")

        _,err := dialContext(ctx, network, address)
        if(err != nil)  {
                fmt.Println("dialContext error!\n")
        }
  
        return;
}
