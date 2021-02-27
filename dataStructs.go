package main

type dbConfig struct {
	Host     string `json:"host"`
	Port     int32  `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Dbname   string `json:"dbname"`
}

type SQLSelectStatements struct {
	Statements map[string]string
}

type TemperatureResponse struct {
	Temp float32 `json:"temp"`
	Unit string  `json:"unit"`
}
