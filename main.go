package main

import(
	"fmt"
	"os"
	"bufio"
	"strings"


)

func main(){

	file, err := os.Open("plantilla/mensage.txt")

	if err != nil{
		fmt.Println("error abriendo archivo")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan(){
		line := scanner.Text()
		words:=  strings.Split(line, ":")		
/*
		if strings.Contains(words[0], ":") == true {
			words[0] = words[0][:len(words[0]) - 1]
		}
*/
		for i,word := range words{
			temp := strings.ReplaceAll(word, " ", "")
			temp = strings.ToLower(temp)
			words[i] = temp
		}

		fmt.Println(strings.Join(words,","))
		fmt.Println(len(words))
		
	}
	
}
