package main

import "math"

type IntegrationSchemeInterval interface {
	CalculateIntegral(Begin *Point, End *Point, NumberOfSegments int, Func func(point *Point) float64) float64
}

//конструктор: на вход подаётся тип квадратурной формулы

func NewIntegrationScheme(Type string) *IntegrationScheme {
	//заполнение массивов точек и весов интегрирования
	switch Type {
	//Методы Гаусса
	case "Gauss1":
		return &IntegrationScheme{
			Weight: []float64{2.0},
			Points: []Point{{0, 0, 0}},
		}
	case "Gauss2":
		return &IntegrationScheme{
			Weight: []float64{1.0, 1.0},
			Points: []Point{
				{-1.0 / math.Sqrt(3.0), 0, 0},
				{1.0 / math.Sqrt(3.0), 0, 0},
			},
		}
	case "Gauss3":
		return &IntegrationScheme{
			Weight: []float64{5.0 / 9.0, 8.0 / 9.0, 5.0 / 9.0},
			Points: []Point{
				{-(math.Sqrt(3.0 / 5.0)), 0, 0},
				{0, 0, 0},
				{math.Sqrt(3.0 / 5.0), 0, 0},
			},
			IntegrationSchemeType: Type,
		}
	case "Gauss4":
		return &IntegrationScheme{
			Weight: []float64{
				(18.0 - math.Sqrt(30.0)) / 36.0,
				(18.0 + math.Sqrt(30.0)) / 36.0,
				(18.0 + math.Sqrt(30.0)) / 36.0,
				(18.0 - math.Sqrt(30.0)) / 36.0,
			},
			Points: []Point{
				{-math.Sqrt((3.0 + 2.0*math.Sqrt(6.0/5.0)) / 7.0), 0, 0},
				{-math.Sqrt((3.0 - 2.0*math.Sqrt(6.0/5.0)) / 7.0), 0, 0},
				{math.Sqrt((3.0 - 2.0*math.Sqrt(6.0/5.0)) / 7.0), 0, 0},
				{math.Sqrt((3.0 + 2.0*math.Sqrt(6.0/5.0)) / 7.0), 0, 0},
			},
		}
	case "Gauss5":
		return &IntegrationScheme{
			Weight: []float64{
				(322.0 - 13.0*math.Sqrt(70.0)) / 900.0,
				(322.0 + 13.0*math.Sqrt(70.0)) / 900.0,
				128.0 / 225.0,
				(322.0 + 13.0*math.Sqrt(70.0)) / 900.0,
				(322.0 - 13.0*math.Sqrt(70.0)) / 900.0,
			},
			Points: []Point{
				{-math.Sqrt(5.0+2.0*math.Sqrt(10.0/7.0)) / 3.0, 0, 0},
				{-math.Sqrt(5.0-2.0*math.Sqrt(10.0/7.0)) / 3.0, 0, 0},
				{0, 0, 0},
				{math.Sqrt(5.0-2.0*math.Sqrt(10.0/7.0)) / 3.0, 0, 0},
				{math.Sqrt(5.0+2.0*math.Sqrt(10.0/7.0)) / 3.0, 0, 0},
			},
		}
		//схема метода трапеций
	case "Trapezoid":
		return &IntegrationScheme{
			Weight: []float64{1.0, 1.0},
			Points: []Point{
				{-1.0, 0, 0},
				{1.0, 0, 0},
			},
			IntegrationSchemeType: Type,
		}
	case "Parabola":
		return &IntegrationScheme{
			Weight: []float64{1.0 / 3.0, 4.0 / 3.0, 1.0 / 3.0},
			Points: []Point{
				{1, 0, 0},
				{0, 0, 0},
				{-1, 0, 0},
			},
		}
	default:
		return &IntegrationScheme{}
	}

}

//метод для вычисления определённого интеграла:
//Begin и End - начало и конец отрезка
//NumOfSegments - число сегментов
//Func - подынтегральная функция

func (it *IntegrationScheme) CalculateIntegral(Begin *Point, End *Point, NumberOfSegments int, Func func(point *Point) float64) float64 {
	//начальная точка сегмента
	var x0 float64
	//результат (квадратурная сумма)
	res := 0.0
	//шаг на отрезке
	h := (End.X - Begin.X) / float64(NumberOfSegments)
	//сумма по всем сегментам разбиения
	for i := 0; i < NumberOfSegments; i++ {
		x0 = Begin.X + float64(i)*h
		//сумма по узлам интегрирования
		for integPoint := 0; integPoint < len(it.Points); integPoint++ {
			//переход с мастер-элемента [-1, 1]
			p := Point{x0 + (1.0+it.Points[integPoint].X)*h/2.0, 0, 0}
			res += it.Weight[integPoint] * Func(&p)
		}
	}
	//формируем результат с учётом якобиана на отрезке [-1, 1]
	return res * (h / 2.0)
}
