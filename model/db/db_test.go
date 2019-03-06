package db

import (
	"fmt"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

func TestMain(m *testing.M) {
	fmt.Println("Тесты DB ******************************************************")
	// Читаем конфиги
	if os.Getenv("RUNNING_IN_DOCKER") == "Y" {
		ReadConfig("../../configs/db-docker.yaml")
	} else {
		ReadConfig("../../configs/db-dev.yaml")
	}
	ReadConfig("../../configs/db.yaml")

	// call flag.Parse() here if TestMain uses flags
	os.Exit(m.Run())
}

func Test_dbAvailable(t *testing.T) {
	if !dbAvailable() {
		t.Errorf("dbAvailable() = false")
	}
}
