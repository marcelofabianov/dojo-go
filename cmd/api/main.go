package main

import (
	_ "github.com/marcelofabianov/fault"

	_ "github.com/marcelofabianov/dojo-go/docs"
	"github.com/marcelofabianov/dojo-go/internal/di"
)

// @title           Dojo Go API
// @version         1.0
// @description     API de exemplo para um Dojo de Go, demonstrando CRUD com arquitetura limpa.
// @termsOfService  http://swagger.io/terms/

// @contact.name   Marcelo Fabiano
// @contact.url    http://www.exemplo.com
// @contact.email  suporte@exemplo.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1
func main() {
	di.New().Run()
}
