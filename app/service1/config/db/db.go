package db

// DBConfig database configuration interface
type DBConfig interface {
	String() string
	DSN() string
}
