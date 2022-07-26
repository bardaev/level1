package tasks

import (
	"math"
	"strconv"
)

func Task8() {
	var val1 int64 = 256
	var pos1 int64 = 5
	var bite1 bool = true

	PrintInputOutput(
		8.1,
		"Установка бита в "+strconv.FormatBool(bite1),
		strconv.FormatInt(setBit(val1, pos1, bite1), 2),
		"Значение: "+strconv.FormatInt(val1, 2),
		"Позиция: "+strconv.FormatInt(pos1, 10),
		"Бит: "+strconv.FormatBool(bite1),
	)

	var val2 int64 = 1023
	var pos2 int64 = 8
	var bite2 bool = false

	PrintInputOutput(
		8.1,
		"Установка бита в "+strconv.FormatBool(bite2),
		strconv.FormatInt(setBit(val2, pos2, bite2), 2),
		"Значение: "+strconv.FormatInt(val2, 2),
		"Позиция: "+strconv.FormatInt(pos2, 10),
		"Бит: "+strconv.FormatBool(bite2),
	)
}

// Бит в 1. Возводим 2 в степень позиции. Получим еденицу в нужном месте и делаем побитовое или
// Бит в 0. Делаем битовый сдвиг вправо и влево чтобы были нули
// и делаем побитовое или с 2 в степени позиция минус 1
// получаем маску вида 1111101111
// и делаем побитовое и. В итоге получаем установленный ноль в нужной позиции
func setBit(val int64, pos int64, bite bool) int64 {
	var result int64 = val

	if bite {
		result = result | int64(math.Pow(2, float64(pos)-1))
	} else {
		tmp := val >> pos
		tmp = tmp << pos
		tmp = tmp | int64(math.Pow(2, float64(pos-1))-1)
		result = result & tmp
	}

	return result
}
