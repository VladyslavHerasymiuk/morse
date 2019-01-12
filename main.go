package main

import (
	"bufio"
	"encoding/csv"
	"log"
	"morse/morse"
	"os"
	"sync"
)

//func ExampleDecodeITU() {
//
//	decoded, _ := morse.DecodeITU("-- .. ... - . .-. / - . -..- - / ..--- ----- ----- -----")
//	fmt.Print(decoded)
//}
//
//func ExampleEncodeITU() {
//	f,_ := ioutil.ReadFile("ro-en/europarl-v7.ro-en.en")
//	data := strings.Split(string(f),"\n")
//	out, _ := os.Create("morze.txt")
//	for _, str :=  range data {
//		encoded := morse.EncodeITU(str)
//		//fmt.Println("English string: ", str, "\nMorse string: ", encoded)
//		//encoded =+ "\n"
//		out.WriteString(encoded)
//		out.WriteString("\n")
//
//	}
//	//encoded := morse.EncodeITU("mister text 2000")
//}
//
//func t() {
//	f, _ := os.Open("ro-en/europarl-v7.ro-en.en")
//	scanner := bufio.NewScanner(f)
//
//	out, _ := os.Create("morze_fast.txt")
//	for scanner.Scan() {
//		line := scanner.Text()
//		encoded := morse.EncodeITU(line)
//		out.WriteString(encoded)
//		out.WriteString("\n")
//		//fmt.Println(line)
//	}
//}
//
//func main(){
//	//ExampleEncodeITU()
//	t()
//	// lol1 := morse.EncodeITU("Membership of Parliament: see Minutes")
//	// lol, _ := morse.DecodeITU(lol1)
//	// fmt.Println(lol, " ", lol1)
//}

type CsvWriter struct {
	mutex *sync.Mutex
	csvWriter *csv.Writer
}

func NewCsvWriter(fileName string) (*CsvWriter, error) {
	csvFile, err := os.Create(fileName)
	if err != nil {
		return nil, err
	}
	w := csv.NewWriter(csvFile)
	return &CsvWriter{csvWriter:w, mutex: &sync.Mutex{}}, nil
}

func (w *CsvWriter) Write(row []string) {
	w.mutex.Lock()
	w.csvWriter.Write(row)
	w.mutex.Unlock()
}

func (w *CsvWriter) Flush() {
	w.mutex.Lock()
	w.csvWriter.Flush()
	w.mutex.Unlock()
}

func main() {
	w, err := NewCsvWriter("foo-safe.txt")
	if err != nil {
		log.Panic(err)
	}
	f, _ := os.Open("ro-en/europarl-v7.ro-en.en")
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		encoded := morse.EncodeITU(line)
		w.Write([]string{encoded})
	}
	w.Flush()
}