package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func openFile(filePath string) *os.File {
	// Abre o Arquivo de config
	f, err := os.Open(filePath)

	// Loga se der erro
	if err != nil {
		processError(err)
	}
	return f

}

func closeFile(f *os.File) {

	// Fecha o arquivo
	err := f.Close()
	if err != nil {
		processError(err)
	}
}

func readFile (fileName string) []byte {

	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		processError(err)
	}
	return b
}

func processError(err error) {
	fmt.Println(err)
	os.Exit(2)
}