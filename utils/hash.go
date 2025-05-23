package utils

import (
	"fmt"
	"os"
)

func DisplayPepper() {
	fmt.Println(os.Getenv("PEPPER"))
}