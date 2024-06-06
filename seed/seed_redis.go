package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
)

type IPInfo struct {
	IP        string `json:"ip"`
	UserAgent string `json:"userAgent"`
}

func main() {
	ctx := context.Background()

	// Conexión a Redis
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	// Ejemplos de datos
	data := []IPInfo{
		{"10.0.0.1", "Mozilla/5.0 (Windows NT 10.0; Win64; x64)"},
		{"172.16.0.1", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7)"},
		{"192.168.1.1", "Mozilla/5.0 (Linux; Android 10)"},
		{"224.0.0.1", "Mozilla/5.0 (iPhone; CPU iPhone OS 14_0 like Mac OS X)"},
		{"240.0.0.1", "Mozilla/5.0 (iPad; CPU OS 14_0 like Mac OS X)"},
		{"10.0.0.2", "Mozilla/5.0 (X11; Linux x86_64)"},
		{"172.16.0.2", "Mozilla/5.0 (Windows NT 6.1; WOW64)"},
		{"192.168.1.2", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6)"},
		{"224.0.0.2", "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:89.0)"},
		{"240.0.0.2", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:89.0)"},
		{"10.0.0.3", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:90.0)"},
		{"172.16.0.3", "Mozilla/5.0 (Linux; Android 11)"},
		{"192.168.1.3", "Mozilla/5.0 (iPhone; CPU iPhone OS 13_6 like Mac OS X)"},
		{"224.0.0.3", "Mozilla/5.0 (iPad; CPU OS 13_6 like Mac OS X)"},
		{"240.0.0.3", "Mozilla/5.0 (X11; Linux x86_64)"},
		{"10.0.0.4", "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:68.0)"},
		{"172.16.0.4", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6)"},
		{"192.168.1.4", "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:88.0)"},
		{"224.0.0.4", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:88.0)"},
		{"240.0.0.4", "Mozilla/5.0 (Linux; Android 9; Pixel 3)"},
		{"2001:0db8:85a3:0000:0000:8a2e:0370:7334", "Mozilla/5.0 (Windows NT 10.0; Win64; x64)"},
		{"2001:0db8:85a3:0000:0000:8a2e:0370:7335", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7)"},
		{"2001:0db8:85a3:0000:0000:8a2e:0370:7336", "Mozilla/5.0 (Linux; Android 10)"},
		{"2001:0db8:85a3:0000:0000:8a2e:0370:7337", "Mozilla/5.0 (iPhone; CPU iPhone OS 14_0 like Mac OS X)"},
		{"2001:0db8:85a3:0000:0000:8a2e:0370:7338", "Mozilla/5.0 (iPad; CPU OS 14_0 like Mac OS X)"},
	}

	// Insertar datos en Redis
	for _, entry := range data {
		key := fmt.Sprintf("ipPublic_%s", entry.IP)
		value, err := json.Marshal(entry)
		if err != nil {
			log.Fatalf("Error al serializar el valor: %v", err)
		}

		err = rdb.Set(ctx, key, value, 0).Err()
		if err != nil {
			log.Fatalf("Error al insertar el valor en Redis: %v", err)
		}
	}

	fmt.Println("Datos insertados en Redis.")
}
