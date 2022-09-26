package testunitario

import "testing"

/*func TestSuma(t *testing.T) {

	total := Suma(5, 5)

	if total != 11 {
		t.Errorf("suma incorrecta tiene %d y se esperaba %d", total, 10)
	}

}*/

func TestSuma(t *testing.T) {
	tabla := []struct {
		a int
		b int
		c int
	}{
		{1, 2, 3},
		{2, 2, 4},
		{25, 25, 50},
	}

	for _, item := range tabla {
		total := Suma(item.a, item.b)
		if total != item.c {
			t.Errorf("suma incorrecta tiene %d y se esperaba %d", total, item.c)
		}
	}

}

func TestGetMax(t *testing.T) {
	tabla := []struct {
		a int
		b int
		c int
	}{
		{1, 2, 2},
		{4, 2, 4},
		{25, 28, 28},
	}

	for _, item := range tabla {
		max := GetMax(item.a, item.b)
		if max != item.c {
			t.Errorf("suma incorrecta tiene %d y se esperaba %d", max, item.c)
		}
	}

}

func TestFibo(t *testing.T) {
	tabla := []struct {
		a int
		b int
	}{
		{1, 1},
		{8, 21},
		{50, 12586269025},
	}

	for _, item := range tabla {
		fibo := Fibonacci(item.a)
		if fibo != item.b {
			t.Errorf("fibo incorrecta tiene %d y se esperaba %d", fibo, item.b)
		}
	}

}
