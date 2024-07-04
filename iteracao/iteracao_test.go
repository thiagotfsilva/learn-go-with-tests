package iteracao

import (
	"fmt"
	"testing"
)

func TestRepetir(t *testing.T) {
	repeticoes := Repertir("a", 6)
	esperado := "aaaaaa"

	if repeticoes != esperado {
		t.Errorf("esperado '%s' mas obteve '%s'", esperado, repeticoes)
	}
}

func BenchmarkRepetir(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repertir("a", 6)
	}
}

func ExampleRepetir() {
	repeticoes := Repertir("b", 3)
	fmt.Println(repeticoes)
	// Output: bbb
}