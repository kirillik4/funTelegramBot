package pisaBrain

import (
	"fmt"
	xls "github.com/xuri/excelize/v2"
	"log"
	"math/rand/v2"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func randRange(min, max int) int {
	return rand.IntN(max-min) + min
}

func textGen(row []string, len ...int) string {
	var text string
	if row[3] == "0" {
		text = fmt.Sprintf("@%s, твой писюн %dсм.\nСледующая попытка завтра.\n", row[1], len[0])
	} else {
		text = fmt.Sprintf("@%s, попробуй завтра.\nТвой писюн %sсм.", row[1], row[2])
	}
	return text
}

type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Less(i, j int) bool { return p[i].Value > p[j].Value }

func mapSorter(pisasNames map[string]int) PairList {
	p := make(PairList, len(pisasNames))
	i := 0
	for k, v := range pisasNames {

		p[i] = Pair{k, v}
		i++
	}

	sort.Sort(p)
	//p is sorted

	return p
}

func PisaMove(UserId int, UserName string) string {
	ID := strconv.Itoa(UserId)
	f, err := xls.OpenFile("LenPisas.xlsx", xls.Options{Password: "1q2w3e4r5t"})
	if err != nil {
		log.Panic(err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Panic(err)
		}
	}()
	rows, err := f.GetRows("Лист1")
	if err != nil {
		log.Panic(err)
	}
	for i, row := range rows {
		if row[0] == ID {
			result, err := f.SearchSheet("Лист1", ID)
			if err != nil {
				log.Panic(err)
			}
			dailyAttempt := strings.Replace(result[0], "A", "D", 1)
			cell := strings.Replace(result[0], "A", "C", 1)
			if row[3] == "0" {
				move := randRange(-10, 10)
				lenBefor, err := strconv.Atoi(row[2])
				if err != nil {
					log.Panic(err)
				}
				lenAfter := lenBefor + move
				f.SetCellValue("Лист1", cell, lenAfter)
				f.SetCellValue("Лист1", dailyAttempt, 1)
				err = f.Save()
				if err != nil {
					log.Panic(err)
				}
				return textGen(row, lenAfter)
			} else {
				return textGen(row)
			}

		} else if row[0] != ID && i == 29 {
			firstLen := addUser(UserId, UserName)
			return firstLen

		}
	}

	return ""
}

func addUser(userID int, userName string) string {
	move := randRange(-10, 10)
	f, err := xls.OpenFile("LenPisas.xlsx", xls.Options{Password: "1q2w3e4r5t"})
	if err != nil {
		log.Panic(err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Panic(err)
		}
	}()
	cols, err := f.Cols("Лист1")
	if err != nil {
		log.Panic(err)
	}
	cols.Next()
	col, err := cols.Rows()
	if err != nil {
		log.Panic(err)
	}
	ints := make([]int, len(col))
	for i, s := range col {
		ints[i], _ = strconv.Atoi(s)
	}
	newRow := slices.Min(ints)
	result, err := f.SearchSheet("Лист1", strconv.Itoa(newRow))
	cellID := result[0]
	cellName := strings.Replace(result[0], "A", "B", 1)
	cellLen := strings.Replace(result[0], "A", "C", 1)
	dailyAttempt := strings.Replace(result[0], "A", "D", 1)
	f.SetCellValue("Лист1", cellID, userID)
	f.SetCellValue("Лист1", cellName, userName)
	f.SetCellValue("Лист1", cellLen, move)
	f.SetCellValue("Лист1", dailyAttempt, 1)
	err = f.Save()
	if err != nil {
		log.Panic(err)
	}
	return fmt.Sprintf("@%s, твой писюн %sсм.\nСледующая попытка завтра.\n", userName, strconv.Itoa(move))
}
func TopPisas() string {
	f, err := xls.OpenFile("LenPisas.xlsx", xls.Options{Password: "1q2w3e4r5t"})
	if err != nil {
		log.Panic(err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Panic(err)
		}
	}()
	pisasNames := make(map[string]int)
	cols, err := f.Cols("Лист1")
	if err != nil {
		log.Panic(err)
	}
	cols.Next()
	cols.Next()
	col1, err := cols.Rows()
	if err != nil {
		log.Panic(err)
	}
	cols.Next()
	col2, err := cols.Rows()
	if err != nil {
		log.Panic(err)
	}
	for i, userName := range col1 {

		if col2[i] == "" {
			continue
		} else {
			pisaLen, err := strconv.Atoi(col2[i])
			if err != nil {
				log.Panic(err)
			}
			pisasNames[userName] = pisaLen
		}
	}
	sortedPisas := mapSorter(pisasNames)
	text := fmt.Sprintf("Топ дилдоков.\n\n1.@%s -  %dсм.\n2.@%s -  %dсм.\n3.@%s -  %dсм.\n4.@%s -  %dсм.\n5.@%s -  %dсм.\n", sortedPisas[0].Key, sortedPisas[0].Value, sortedPisas[1].Key, sortedPisas[1].Value, sortedPisas[2].Key, sortedPisas[2].Value, sortedPisas[3].Key, sortedPisas[3].Value, sortedPisas[4].Key, sortedPisas[4].Value)
	return text
}
