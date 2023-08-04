package main

import "math"

type IntegrationSchemeInterval interface {
	CalculateIntegral(Begin *Point, End *Point, NumberOfSegments int, Func func(point *Point) float64) float64
}

//конструктор: на вход подаётся тип квадратурной формулы

func NewIntegrationScheme(Type string) *IntegrationScheme {
	//заполнение массивов точек и весов интегрирования
	switch Type {
	//схема метода Гаусс-3
	case "Gauss3":
		return &IntegrationScheme{
			Weight: []float64{
				5.0 / 9.0, 8.0 / 9.0, 5.0 / 9.0,
			},
			Points: []Point{{
				X: -(math.Sqrt(3.0 / 5.0)), Y: 0, Z: 0,
			}, {
				X: 0, Y: 0, Z: 0,
			}, {
				X: math.Sqrt(3.0 / 5.0), Y: 0, Z: 0,
			}},
			IntegrationSchemeType: Type,
		}
		//схема метода трапеций
	case "Trap":
		return &IntegrationScheme{
			Weight: []float64{
				1.0, 1.0,
			},
			Points: []Point{{
				X: -1.0, Y: 0, Z: 0,
			}, {
				X: 1.0, Y: 0, Z: 0,
			},
			},
			IntegrationSchemeType: Type,
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
