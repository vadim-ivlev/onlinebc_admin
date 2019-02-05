package redis

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("Тесты Redis ******************************************************")

	if os.Getenv("RUNNING_IN_DOCKER") == "Y" {
		ReadConfig("../../configs/redis-docker.yaml")
	} else {
		ReadConfig("../../configs/redis-dev.yaml")
	}
	ReadConfig("../../configs/redis.yaml")

	// call flag.Parse() here if TestMain uses flags
	os.Exit(m.Run())
}

func TestSetGetDelete(t *testing.T) {

	// Записываем k1 -> v1
	fmt.Println("Записываем k1 -> v1")
	if err := Set("k1", "v1"); err != nil {
		t.Errorf("Set() error = %v", err)
	}

	// Считываем k1
	fmt.Println("Считываем k1")
	got, err1 := Get("k1")
	if err1 != nil {
		t.Errorf("Get() error = %v", err1)
		return
	}

	if got != "v1" {
		t.Errorf("Get() = %v, want %v", got, "v1")
	}

	// Удаляем k1
	fmt.Println("Удаляем k1")
	if err2 := Del("k1"); err2 != nil {
		t.Errorf("Del() error = %v", err2)
	}

	// Проверяем что удалили k1
	fmt.Println("Проверяем что удалили k1")
	got, err3 := Get("k1")
	if err3 == nil {
		t.Errorf("Get() error = %v", err3)
	}
	if got != "" {
		t.Errorf("Get() = %v, want %v", got, "")
	}

}
