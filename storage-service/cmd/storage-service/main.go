package main

import (
	"fmt"
	"github.com/LavaJover/storage-master/storage-service/internal/config"
)

func main(){
	cfg := config.MustLoad()
	fmt.Println(cfg)
}