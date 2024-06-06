package service

import (
	"database/sql"
	"log"

	"github.com/ccxnu/ips-redis-mysql/config"
	"github.com/ccxnu/ips-redis-mysql/internal/database"
	"github.com/ccxnu/ips-redis-mysql/pkg/model"
)

func SaveToDatabase(data *model.IPData) (bool, error) {

	db, err := database.NewMysqlConnection(&config.AppConfig)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	defer db.Close()

	// Check if the IP already exists in the SaveToDatabase
	ipExists, err := GetInfoById(db, data.IP)
	if err != nil {
		log.Printf("Error checking if IP exists in the database: %v", err)
		return false, err
	}

	// Poco pirata esto
	if ipExists {
		return true, err
	}

	query := `INSERT INTO GeoData (ip, country, countryCode, region, regionName,
                                city, zip, latitud, longitud, timezone, isp,
                                org, proveedor, userAgent) VALUES
                                (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	_, err = db.Exec(query, data.IP, data.Country, data.CountryCode, data.Region, data.RegionName, data.City, data.Zip, data.Lat, data.Lon, data.Timezone, data.ISP, data.Org, data.Proveedor, data.UserAgent)

	if err != nil {
		log.Printf("Error saving data to database: %v", err)
		return false, err
	}

	return false, nil
}

func GetInfoById(db *database.DB, ip string) (bool, error) {
	ipExists := false

	querySearh := `SELECT ip FROM GeoData WHERE ip = ?`
	if err := db.QueryRow(querySearh, ip).Scan(); err != sql.ErrNoRows {
		log.Printf("IP %s already exists in the database", ip)
		ipExists = true
		return ipExists, nil
	}

	return ipExists, nil
}
