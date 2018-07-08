package common

import "fmt"

func ConstructPostgresDataSourceString(databaseName string, user string, sslMode string) string {
	return fmt.Sprintf("user=%s dbname=%s sslmode=%s", user, databaseName, sslMode)
}
