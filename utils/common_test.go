package utils

import (
	"testing"
	"log"
	
)

func TestRevert(t *testing.T){
	array := []interface{}{
		"134", "345", "567",
	}

	array1 := []interface{}{
		"134", "345", "567","9999",
	}

	array = RevertArray(array)
	log.Printf("array:%v", array)

	array1 = RevertArray(array1)
	log.Printf("array1:%v", array1)
}

func TestTimeLocation(t *testing.T) {
	time := FormatTime(1511680608304)
	log.Printf(time)
}