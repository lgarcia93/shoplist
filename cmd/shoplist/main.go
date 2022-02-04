package main

import (
	_ "github.com/lgarcia93/shoplist/api/swagger/docs"
	"github.com/lgarcia93/shoplist/internal/pkg/shoplist/starter"
)

func main() {
	starter.InitializeHandlers()
}
