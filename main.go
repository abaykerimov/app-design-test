// Ниже реализован сервис бронирования номеров в отеле. В предметной области
// выделены два понятия: Order — заказ, который включает в себя даты бронирования
// и контакты пользователя, и RoomAvailability — количество свободных номеров на
// конкретный день.
//
// Задание:
// - провести рефакторинг кода с выделением слоев и абстракций
// - применить best-practices там где это имеет смысл
// - исправить имеющиеся в реализации логические и технические ошибки и неточности
package main

import (
	"applicationDesignTest/internal"
	"errors"
	"net/http"
	"os"
)

func main() {
	a, err := internal.InitializeApplication()
	if err != nil {
		a.Logger.LogErrorf("Failed init app: %s", err)
	}

	err = a.Start()
	if errors.Is(err, http.ErrServerClosed) {
		a.Logger.LogErrorf("Server closed")
	} else if err != nil {
		a.Logger.LogErrorf("Server failed: %s", err)
		os.Exit(1)
	}

}
