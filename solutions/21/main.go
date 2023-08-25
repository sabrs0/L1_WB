/*
Пусть имеется интерфейс для работы с бд - IDataBase.
Пусть также имеется база данных, которая реализует этот интерфейс и

	по id, который представлен целым числом,которое хранится как строка, и  возращает какие-то данные

Есть также класс Queryer, который делает запрос в бд, которая реализует IDataBase.

# Затем пусть логика хранения данных изменилась и теперь id - это также число, но в в виде int64

Нужен адаптер между Queryer и Новым типом БД.
*/
package main

import (
	"fmt"
	"strconv"
)

// Интерфейса базы
type IDataBase interface {
	GetDataByID(id string) (string, error)
}

// Старая база. Хранит id в виде строки
type LegacyDataBase struct {
	login    string
	password string
	data     map[string]int
}

// Метод старой базы
func (od LegacyDataBase) GetDataByID(id string) (string, error) {

	if _, isOK := od.data[id]; isOK {
		return fmt.Sprintf("Data by id %s is %d\n", id, od.data[id]), nil
	} else {
		return "", fmt.Errorf("Cant Get Data By ID")
	}
}

// Новая база. Хранит id в виде int64
type CurrentDataBase struct {
	data map[int64]int
}

func (nd CurrentDataBase) GetDataByID(id int64) (string, error) {
	if _, isOK := nd.data[id]; isOK {
		return fmt.Sprintf("Data by id %d is %d\n", id, nd.data[id]), nil
	} else {
		return "", fmt.Errorf("Cant Get Data By ID")
	}
}

// Адаптер, он реализует интерфейс IDataBase
type CurrentDataBaseAdapter struct {
	CurrentDataBase CurrentDataBase
}

func NewCurrentDataBaseAdapter(curDB CurrentDataBase) CurrentDataBaseAdapter {

	return CurrentDataBaseAdapter{
		CurrentDataBase: curDB,
	}
}

// Выполняем некую логику с id чтобы была возможность получать данные из currentDataBase
func (adptr CurrentDataBaseAdapter) GetDataByID(id string) (string, error) {

	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return "", err
	}
	return adptr.CurrentDataBase.GetDataByID(idInt)
}

// Просто тот, кто делает запрос
type Queryer struct {
}

// Сам запрос
func (q Queryer) RecData(db IDataBase, id string) {
	data, err := db.GetDataByID(id)
	if err != nil {
		fmt.Println("Cant get data : ", err.Error())
	} else {
		fmt.Println("Data recieved: ", data)
	}

}

func main() {
	newDB := CurrentDataBase{
		data: make(map[int64]int),
	}
	newDB.data[123] = 456
	newDB.data[456] = 789

	oldDB := LegacyDataBase{
		data: make(map[string]int),
	}
	oldDB.data["12"] = 34
	oldDB.data["56"] = 78
	oldDB.data["0"] = 1
	q := Queryer{}

	q.RecData(oldDB, "56")
	dbAdapter := NewCurrentDataBaseAdapter(newDB)
	q.RecData(dbAdapter, "456")

}
