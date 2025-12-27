package messageManger

import(
	"fmt"
	"strings"
	"strconv"
)

type Stats struct {
	Key string
	Value int
}

func GetItems(message string) slice []Stats {

	list := []Stats{}

	lookUp := map[string]int{
                "oficina":1,
                "ua":1,
                "ucc":1,
                "uco":1,
                "rcam":1,
                "ucp":1,
                "ci":1,
                "tc":1,
                "upccib":1,
                "upcc":1,
                "icpc":1,
                "icpp":1,
                "cédulasentregadas":1,
                "pasaportesentregados":1,
                "prorrogasentregadas":1,
                "cédulasingresadas":1,
                "pasaportesingresados":1,
                "prorrogasingresadas":1,
        }

	lookUpOffice := map[string]int{
                "cagua":1,
                "maracay":0,
                "lavictoria":2,
                "josefelix":3,
                "sansebastian":4,
        }

	lines := strings.Split(message, "\n")	

	for _, line := range lines{

		words:=  strings.Split(line, ":")

                for i, word := range words{
                        temp := strings.ReplaceAll(word, " ", "")
                        temp = strings.ToLower(temp)
                        words[i] = temp
                }

		_, ok := lookUp[words[0]]

		var item Stats

		if ok {
			if words[0] == "oficina"{
				item = Stats{
					Key: "oficina"
					Value: lookUpOffice[words[0]]
				}
			}else{
				val, _ = strconv.Atoi(words[1])
				item = Stats{
					Key: words[0]	
					Value: val
				}
			}

			list = append(list, item)
		}

	}

	return list

}
