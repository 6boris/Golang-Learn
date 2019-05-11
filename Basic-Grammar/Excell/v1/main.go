package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
)

func main() {
	excelFileName := "Basic-Grammar/Excell/v1/11.xlsx"
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		panic(err)
	}
	for s, sheet := range xlFile.Sheets {
		if s == 1 {
			break
		}
		for _, row := range sheet.Rows {
			for j, cell := range row.Cells {
				if j == 0 {
					continue
				}
				switch cell.Type() {
				case xlsx.CellTypeString:
					fmt.Println(row.Cells[0], row.Cells[j+1], row.Cells[j+2], row.Cells[j+3])
					fmt.Print(j, cell.String())
					if (j-1)%3 == 1 {
						fmt.Println(111)
						fmt.Print(cell.Float())

					}
					break
				case xlsx.CellTypeNumeric:
					num, err := cell.Float()
					if err != nil {

					}

					fmt.Println(row.Cells[0], ":", num)
					break
				}

			}
			fmt.Println()
		}
	}

}
