package main

import (
	"fmt"
	"math"
)

func main() {
	//подынтегральная функция f(x) = ln(0.5x)
	f := func(point *Point) float64 {
		return math.Log(0.5 * point.X)
	}
	//первообразная F(x) = -ln(2) * x + ln(x) * x - x
	F := func(point *Point) float64 {
		return -math.Log(2.0)*point.X + math.Log(point.X)*point.X - point.X
	}

	//Квадратурные формы
	quadratureFormula1 := NewIntegrationSchemeInterval("Gauss3")
	//quadratureFormula2 := NewIntegrationSchemeInterval("Trap")

	//начало и конец отрезка интегрирования
	begin := Point{1, 0, 0}
	end := Point{math.Exp(1), 0, 0}

	//Число сегментов
	numOfSegments := 1

	//точное значение интеграла (ф. Ньютона-Лейбница)
	trueIntegral := F(&end) - F(&begin)
	fmt.Printf("I_true = %g\n", trueIntegral)

	I := make([]float64, 3)
	Ih := make([]float64, 3)

	for i := 0; i < 3; i++ {
		//Гаусс-3
		//I с текущим шагом i
		fmt.Println("Метод Гаусса-3 вычисления определенного интеграла")
		I[i] = quadratureFormula1.CalculateIntegral(&begin, &end, numOfSegments*int(math.Pow(2.0, float64(i))), f)
		Ih[i] = quadratureFormula1.CalculateIntegral(&begin, &end, numOfSegments*int(math.Pow(2.0, float64(i)))*2, f)

		////Трапеции
		//fmt.Println("Метод Трапеций вычисления определенного интеграла")
		//I[i] = quadratureFormula2.CalculateIntegral(&begin, &end, numOfSegments*int(math.Pow(2.0, float64(i))), f)
		//Ih[i] = quadratureFormula2.CalculateIntegral(&begin, &end, numOfSegments*int(math.Pow(2.0, float64(i)))*2, f)

		fmt.Printf("%d -------------------------------------------------\n", i+1)
		fmt.Printf("Для %d отрезков. Шаг h = %g\n", i+1, (end.X-begin.X)/(float64(numOfSegments)*math.Pow(2.0, float64(i))))
		fmt.Printf("Значение по методу квадратур I = %g\n", I[i])
		fmt.Printf("Абсолютная погрешность |I_true - I| = %g\n", math.Abs(trueIntegral-I[i]))
		k := math.Round(math.Log2(math.Abs((1 + (Ih[i]-I[i])/(trueIntegral-Ih[i])))))
		fmt.Printf("Порядок аппроксимации k = %d\n\n", int(k))
		fmt.Printf("Отношение (I_True - I)/(I_True - Ih) = %g\n", (trueIntegral-I[i])/(trueIntegral-Ih[i]))
		R := (Ih[i] - I[i]) / (math.Pow(2.0, k) - 1)
		fmt.Printf("Правило Рунге (Ih - I)/(2^k - 1) = %g\n", R)
		IR := Ih[i] + R
		fmt.Printf("Уточнение по Ричардсону IR = %g\n", IR)
		fmt.Printf("Погрешность |I_True - IR| = %g\n\n", trueIntegral-IR)
	}

}
