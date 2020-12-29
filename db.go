package main

import (
	"database/sql"
	"fmt"
	"math/rand"
	"strconv"

	_ "github.com/lib/pq"
)

var db *sql.DB

type table struct {
	ID    int
	Group int
	Name  string
	Score int
}

// Открываем базу данных
func connect() error {
	var err error
	db, err = sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.PgHost, cfg.PgPort, cfg.PgUser, cfg.PgPass, cfg.PgBase))
	if err != nil {
		panic(err.Error())
	}

	return nil

}

// Функция на добавление пользователя
func queryAdd(Name string, Group int, Score int) error {
	var err error
	_, err = db.Exec(`INSERT INTO users ("group", "name", "score") VALUES ($1, $2, $3)`, Group, Name, Score)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil

}

// На удаление
func queryRmv(ID int) error {
	_, err := db.Exec(`DELETE FROM users WHERE ("id")=($1)`, ID)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

// На изменение очков
func queryCng(Score int, ID int) error {
	_, err := db.Exec(`UPDATE users SET "score"=($1) WHERE("id")=($2)`, Score, ID)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

// Подготовка к сортировке очков
func queryRow(Group int) ([]int, error) {
	rows, err := db.Query(`select "score" from users WHERE "group" = ($1)`, Group)
	if err != nil {
		return nil, err
	}
	values := []string{}
	for rows.Next() {
		var score int

		err := rows.Scan(&score)
		if err != nil {
			return nil, err
		}
		a := strconv.Itoa(score)
		values = append(values, a)
	}
	a := []int{}
	for i := 0; i < len(values); i++ {
		b, err := strconv.Atoi(values[i])
		if err != nil {
			return nil, err
		}
		a[i] = b
	}
	return a, err
}

// На вывод топ x пользователей
func queryTop(a []int, x int) []int {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	pivot := rand.Int() % len(a)

	a[pivot], a[right] = a[right], a[pivot]

	for i := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	quicksort(a[:left])
	quicksort(a[left+1:])
	a = a[:x]
	return a
}
