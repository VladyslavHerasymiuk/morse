package main

import (
	"./morse"
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func ExampleDecodeITU() {

	decoded, _ := morse.DecodeITU("-- .. ... - . .-. / - . -..- - / ..--- ----- ----- -----")
	fmt.Print(decoded)
}

func ExampleEncodeITU() {
	f,_ := ioutil.ReadFile("ro-en/europarl-v7.ro-en.en")
	data := strings.Split(string(f),"\n")
	out, _ := os.Create("morze.txt")
	for _, str :=  range data {
		encoded := morse.EncodeITU(str)
		//fmt.Println("English string: ", str, "\nMorse string: ", encoded)
		//encoded =+ "\n"
		out.WriteString(encoded)
		out.WriteString("\n")

	}
	//encoded := morse.EncodeITU("mister text 2000")
}

func t() {
	f, _ := os.Open("ro-en/europarl-v7.ro-en.en")
	scanner := bufio.NewScanner(f)

	out, _ := os.Create("morze_fast.txt")
	for scanner.Scan() {
		line := scanner.Text()
		encoded := morse.EncodeITU(line)
		out.WriteString(encoded)
		out.WriteString("\n")
		//fmt.Println(line)
	}
}

func main(){
	//ExampleEncodeITU()
	t()
	// lol1 := morse.EncodeITU("Membership of Parliament: see Minutes")
	// lol, _ := morse.DecodeITU(lol1)
	// fmt.Println(lol, " ", lol1)
}
