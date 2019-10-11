package main

import (
        "context"
        "fmt"
        kafka "github.com/segmentio/kafka-go"
        "os"
)

func main() {
        const (
                firstOffset = -2
                lastOffset  = -1
        )
        ctx, _ := context.WithCancel(context.Background())
        config := kafka.ReaderConfig{
                Brokers: []string{"knode1:19092", "knode2:29092", "knode3:39092"},
                Topic:   "topico"}
        r := kafka.NewReader(config)
        if len(os.Args) <= 1 {
                os.Exit(1)
        }
        if len(os.Args) >= 1 && os.Args[1] == "first" {
                r.SetOffset(firstOffset)
        } else {
                r.SetOffset(lastOffset)
        }
        for {
                m, err := r.ReadMessage(ctx)
                if err != nil {
                        fmt.Println("Error reading message: ", err)
                        break
                }
                fmt.Println(m.Offset, string(m.Key), string(m.Value))
	}
}
