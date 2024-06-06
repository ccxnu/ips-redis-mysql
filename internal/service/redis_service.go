package service

import (
	"context"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/ccxnu/ips-redis-mysql/config"
	"github.com/ccxnu/ips-redis-mysql/internal/database"
	"github.com/ccxnu/ips-redis-mysql/internal/util"
	"github.com/ccxnu/ips-redis-mysql/pkg/model"
)

type RedisService struct {
	config       config.Config
	pollInterval time.Duration
	cursor       uint64
}

func NewRedisService(appConfig config.Config) *RedisService {
	return &RedisService{
		config:       appConfig,
		pollInterval: time.Duration(appConfig.PollInterval) * time.Millisecond,
		cursor:       0,
	}
}

func (s *RedisService) Run() error {
	for {
		err := s.handleDataFetch()
		if err != nil {
			log.Printf("Error fetching data: %v", err)
		}

		time.Sleep(s.pollInterval)
	}
}

func (s *RedisService) handleDataFetch() error {
	totalKeys := s.config.TotalKeys
	matchPattern := s.config.MatchPattern

	ctx := context.Background()

	rdb, err := database.NewRedisConnection(s.config)

	if err != nil {
		return err
	}

	defer rdb.Close()

	for {
		var keys []string
		keys, s.cursor, err = rdb.Scan(ctx, s.cursor, matchPattern, totalKeys).Result()
		if err != nil {
			return err
		}

		for _, key := range keys {
			val, err := rdb.Get(ctx, key).Result()
			if err != nil {
				log.Printf("Error getting value of %s: %v", key, err)
				continue
			}

			value, err := util.StringToJson[model.IPValue](val)
			if err != nil {
				log.Fatalf("Failed to parse value info: %v", err)
			}

			// Fetch and process IP info here
			data, err := s.FetchApiInfo(value.IP)
			if err != nil {
				log.Printf("Error getting IP Info: %v", err)
				continue
			}

			if data == nil {
				continue
			}

			newData := &model.IPData{
				IP:          value.IP,
				Country:     data.Country,
				CountryCode: data.CountryCode,
				Region:      data.Region,
				RegionName:  data.RegionName,
				City:        data.City,
				Zip:         data.Zip,
				Lat:         data.Lat,
				Lon:         data.Lon,
				Timezone:    data.Timezone,
				ISP:         data.ISP,
				Org:         data.Org,
				Proveedor:   data.Proveedor,
				UserAgent:   value.UserAgent,
			}

			// Save data to database
			err = SaveToDatabase(newData)

			if err != nil {
				log.Fatalf("Error saving data to database: %s", err)
				continue
			}

			// Delete key from Redis
			/*err = rdb.Del(ctx, key).Err()
			if err != nil {
				log.Fatalf("Error deleting key from Redis: %s", err)
				continue
			}*/
		}
		return nil
	}
}

func (s *RedisService) FetchApiInfo(ip string) (*model.IPData, error) {
	res, err := http.Get(s.config.IpApiUrl + "/" + ip)
	if err != nil {
		log.Printf("Error fetching IP info: %v", err)
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Printf("Unexpected status code: %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body) // response body is []byte
	data, err := util.BodyToJson[model.IPData](body)
	if err != nil {
		log.Fatalf("Failed to parse value info: %v", err)
	}

	return data, nil
}
