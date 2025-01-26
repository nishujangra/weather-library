package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Database struct {
	Connection *sql.DB
}

func NewDatabase(databaseURL string) *Database {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	log.Println("Connected to the database successfully!")
	return &Database{Connection: db}
}

func (db *Database) CreateTable() {
	query := `
	CREATE TABLE IF NOT EXISTS weather_forecast (
		id SERIAL PRIMARY KEY,
		city_name VARCHAR(255) NOT NULL,
		lat FLOAT NOT NULL,
		lon FLOAT NOT NULL,
		temperature FLOAT NOT NULL,
		humidity FLOAT NOT NULL,
		description VARCHAR(255) NOT NULL,
		date_of_forecast VARCHAR(255) NOT NULL,
		recorded_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`

	_, err := db.Connection.Exec(query)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}

	log.Println("Table created or already exists!")
}

func (db *Database) InsertWeatherData(cityName string, lat, lon, temperature, humidity float64, description string, dateOfForecast string) {
	query := `
	INSERT INTO weather_forecast (city_name, lat, lon, temperature, humidity, description, date_of_forecast)
	VALUES ($1, $2, $3, $4, $5, $6, $7);`

	_, err := db.Connection.Exec(query, cityName, lat, lon, temperature, humidity, description, dateOfForecast)
	if err != nil {
		log.Printf("Failed to insert data into the database: %v", err)
	}
	// } else {
	// 	log.Println("Weather data inserted into the database successfully!")
	// }
}
