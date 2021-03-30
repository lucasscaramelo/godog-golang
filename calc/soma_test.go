package soma

import (
	"fmt"

	"github.com/cucumber/godog"
)

var Resultado int

func faoASomaDe(v1, v2 int) error {
	Resultado = v1 + v2
	return nil
}

func oResultadoDeveSer(soma int) error {
	fmt.Println(Resultado)

	if Resultado != soma {
		return fmt.Errorf("Erro ao realizar soma. Experado: %v Obtido: %v", soma, Resultado)
	}

	return nil
}

func featureContext(s *godog.ScenarioContext) {
	s.Step(`^faço a soma de (\d+) \+ (\d+)$`, faoASomaDe)
	s.Step(`^o resultado deve ser (\d+)$`, oResultadoDeveSer)
}
