package main

import (
    "encoding/json"
    "fmt"
    "math/rand"
    "os"
    "time"
)

func main() {
    ticker := time.NewTicker(15 * time.Second)

    for range ticker.C {
        data := map[string]interface{}{
            "status": map[string]interface{}{
                "water": rand.Intn(100) + 1,
                "wind":  rand.Intn(100) + 1,
            },
        }

        json, err := json.Marshal(data)
        if err != nil {
            fmt.Println(err)
            os.Exit(1)
        }

        err = os.WriteFile("data.json", json, 0644)
        if err != nil {
            fmt.Println(err)
            os.Exit(1)
        }
    }
}
