package engine

import "testing"

func TestStartEngines(t *testing.T) {

	teste := StartEngines(50)
	resultado := true
	if teste != resultado {
		t.Error("Expected:", resultado, "Got:", teste)
	}

}
