package service

import (
	"database/sql"
	"log"

	"github.com/ccxnu/ips-redis-mysql/config"
	"github.com/ccxnu/ips-redis-mysql/internal/database"
	"github.com/ccxnu/ips-redis-mysql/pkg/model"
)

func SaveToDatabase(data *model.IPData) error {

	db, err := database.NewMysqlConnection(&config.AppConfig)
	if err != nil {
		return err
	}

	defer db.Close()

	querySearh := `SELECT ip FROM GeoData WHERE ip = ?`
	if err := db.QueryRow(querySearh, data.IP).Scan(); err != sql.ErrNoRows {
		log.Printf("IP %s already exists in the database", data.IP)
		return nil
	}

	query := `INSERT INTO GeoData (ip, country, countryCode, region, regionName,
                                city, zip, latitud, longitud, timezone, isp,
                                org, proveedor, userAgent) VALUES
                                (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	_, err = db.Exec(query, data.IP, data.Country, data.CountryCode, data.Region, data.RegionName, data.City, data.Zip, data.Lat, data.Lon, data.Timezone, data.ISP, data.Org, data.Proveedor, data.UserAgent)

	if err != nil {
		return err
	}

	return nil
}
