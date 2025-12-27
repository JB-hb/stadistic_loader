package fillX

import(

	"fmt"

	"github.com/xuri/excelize/v2"
)

type DayNode struct {
	Offices [] map[string]map[string]int
	next DayNode
}

type DaysList struct {
	Head DayNode
	Length int
}

daysCell := []string{"C","D","E","F","G"}

func WriteWeekExcel(days DaysList, elements []string, offices []string) error {
	
	excel, err := excelize.OpenFile("../plantilla/estadisticas.xlsx")

	if err != nil{
		return err
	}

	defer excel.Close()

	for i := 0; i < days.Length; i++ {
		data := days.GetByIndex(i)
		for j := 0; j < len(offices); j++{
			temp := data.GetOffice(offices[j])
			for g := 0; g < len(elements); g++{
				totalOff := baseOff + (j * 18) + g
				cell := fmt.Sprintf("%s%d",daysCell[i], totalOff)	
				value := temp[elements[g]]
				excel.SetCellValue("Sheet1", cell, value)
			}	
		}
	} 

} 
