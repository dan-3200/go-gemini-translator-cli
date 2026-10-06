// App CLI de tradução de palavras
package main

import (
	front "app/internal"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// executa antes do 'main()'
func init() {
	// carregar variveis de ambiente
	if err := godotenv.Load(); err != nil && !os.IsNotExist(err) {
		fmt.Printf("Erro ao carregar as variáveis de ambiente: %v\n", err)
	}
}

func main() {
	front.Run()
}
