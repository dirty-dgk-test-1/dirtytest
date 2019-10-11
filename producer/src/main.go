package main

import (
        "context"
        "fmt"
        kafka "github.com/segmentio/kafka-go"
        "time"
)

func main() {
        config := kafka.WriterConfig{
                Brokers:      []string{"knode1:19092","knode2:29092","knode3:39092"},
                Topic:        "topico",
                BatchTimeout: 100 * time.Millisecond}
        w := kafka.NewWriter(config)
        i := 1
        for {
                message := fmt.Sprintf("msg-%d",i)
                err := w.WriteMessages(context.Background(), kafka.Message{Value: []byte(message)})
                if err == nil {
                        fmt.Println("Sent message: ", message)
                } else {
                        fmt.Println("Error sending message: ", err)
                }
                i++
        }
}
