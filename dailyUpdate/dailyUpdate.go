package dailyUpdate

import (
	"fmt"
	xls "github.com/xuri/excelize/v2"
	"log"
)

func Update() {
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
	cols.Next()
	cols.Next()
	cols.Next()
	col, err := cols.Rows()
	if err != nil {
		log.Panic(err)
	}
	for i := 1; i <= len(col); i++ {
		cell := fmt.Sprintf("D%d", i)
		f.SetCellValue("Лист1", cell, 0)
		err = f.Save()
		if err != nil {
			log.Panic(err)
		}
	}
}
