package main

import(
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
	"time"

	"github.com/xuri/excelize/v2"

)

func main(){

	baseOffSet := 1

	days := []string{"C","D","E","F","G"}

	loc, _ := time.LoadLocation("America/Caracas")
	t := time.Now().In(loc)

	lookUp := map[string]int{
		"oficina":2,
		"ua":1,
		"ucc":2,
		"uco":3,
		"rcam":4,
		"ucp":5,
		"ci":6,
		"tc":7,
		"upccib":8,
		"upcc":9,
		"icpc":10,
		"icpp":11,
		"cédulasentregadas":12,
		"pasaportesentregados":13,
		"prorrogasentregadas":14,
		"cédulasingresadas":15,
		"pasaportesingresados":16,
		"prorrogasingresadas":17,
	}

	lookUpOffice := map[string]int{
		"cagua":2,
		"maracay":1,
		"lavictoria":3,
		"josefelix":4,
		"sansebastian":5,
	}

	file, err := os.Open("plantilla/mensage.txt")

	excel, erre := excelize.OpenFile("plantilla/estadisticas.xlsx")

	if err != nil{
		fmt.Println("error abriendo archivo")
	}
	if erre != nil{
		fmt.Println("error abriendo archivo excel")
	}

	defer file.Close()
	defer excel.Close()

	scanner := bufio.NewScanner(file)

	var multy int
	var totalOff int

	day := t.Weekday() - 1

	for scanner.Scan(){
		line := scanner.Text()
		words:=  strings.Split(line, ":")		

		for i, word := range words{
			temp := strings.ReplaceAll(word, " ", "")
			temp = strings.ToLower(temp)
			words[i] = temp
		}

		_, ok := lookUp[words[0]]	

		if ok{
			var name string
			var value int

			if words[0] == "oficina"{
				name = strings.ToLower(words[1])	
				_, ok := lookUpOffice[name]
				if ok{
					multy = lookUpOffice[name]
				}
			}else{
				value,_ = strconv.Atoi(words[1])
				totalOff = baseOffSet + (multy * lookUp[words[0]])
				var cell string 

				if (day > 5  || day < 0){
					cell = fmt.Sprintf("%s%d", days[0], totalOff)
				}else{
					cell = fmt.Sprintf("%s%d", days[day], totalOff)
				}

				fmt.Println(cell)

				excel.SetCellValue("Sheet1", cell, value)
			}
		}
	}

	o := excel.Save()

	if o  != nil{
		fmt.Println(o)
	}
}
