package main

import(
	"fmt"
	"os"
	"bufio"
	"strings"
	//"strconv"



)

func main(){

	baseOffSet := 1

	//days := []string{"C","D","E","F","G"}

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

	if err != nil{
		fmt.Println("error abriendo archivo")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var multy int
	var totalOff int

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
			//var value int

			if words[0] == "oficina"{
				name = strings.ToLower(words[1])	
				_, ok := lookUpOffice[name]
				if ok{
					multy = lookUpOffice[name]
					fmt.Println(multy)
				}
			}else{
				//value,_ = strconv.Atoi(words[1])
				totalOff = baseOffSet + (multy * lookUp[words[0]])
			}
			var cell string = fmt.Sprintf("%s%d", "A", totalOff)
			fmt.Println(cell)
		}
	}

		
}
