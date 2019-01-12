package test

import (
	"testing"
	"../morse"
	"strings"
	"io/ioutil"
	"fmt"
	"regexp"
)

func TestLooksLikeMorse(t *testing.T) {

	if !morse.LooksLikeMorse("- .... . .-..") {
		t.Error("fail 1")
	}
	if morse.LooksLikeMorse("one one one...") {
		t.Error("fail 2")
	}
	if morse.LooksLikeMorse("") {
		t.Error("fail 3")
	}
}

func TestDecodeBrokenITU(t *testing.T) {

	_, err := morse.DecodeITU("-- ?.. ...")
	if err == nil {
		t.Error("expected error")
	}
}

func TestDecodeFile(t *testing.T) {
	engl,_ := ioutil.ReadFile("../ro-en/europarl-v7.ro-en.en")
	data_eng := strings.Split(string(engl),"\n")
	morze,_ := ioutil.ReadFile("../morze.txt")
	data_mrz := strings.Split(string(morze),"\n")
	reg, _ := regexp.Compile("[^a-zA-Z0-9]+")
	for num, _ :=  range data_eng {
		eng_decode, _ :=  morse.DecodeITU(data_mrz[num])
		//fmt.Println("English string: ", data_eng[num], "\nMorse string: ", eng_decode)
		if reg.ReplaceAllString(strings.ToLower(data_eng[num]), "") != eng_decode {
			t.Error("fail")
			fmt.Println("English string: ", data_eng[num], "\nMorse string: ", eng_decode)
		}
	}
}
