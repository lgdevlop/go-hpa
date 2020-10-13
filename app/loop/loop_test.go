package loop

import (
	"testing"
)

func TestLoop(testing *testing.T) {
	mensagemEsperada := "Code.education Rocks!"
	mensagem := Loop()

	if mensagem != mensagemEsperada {
		testing.Errorf("Mensagem esperada é %v, mas o resultado foi %v", mensagemEsperada, mensagem)
	}
}
