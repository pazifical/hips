package main

import (
	_ "embed"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

var suffix = ".jpg"

//go:embed forest.jpg
var carrier []byte

var fileToHide string
var fileToUncover string

func init() {
	flag.StringVar(&fileToHide, "hide", "", "file to hide")
	flag.StringVar(&fileToUncover, "uncover", "", "file to uncover")
	flag.Parse()
}

func main() {
	if fileToHide != "" {
		err := hide()
		if err != nil {
			fmt.Println(err)
		}
	} else if fileToUncover != "" {
		err := uncover()
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println("Please provide a command line argument")
	}
}

func hide() error {
	fmt.Printf("Hiding %s\n", fileToHide)
	payload, err := os.ReadFile(fileToHide)
	if err != nil {
		return err
	}

	f, err := os.Create(fmt.Sprintf("%s%s", fileToHide, suffix))
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(carrier)
	if err != nil {
		return err
	}

	_, err = f.Write(payload)
	if err != nil {
		return err
	}

	return nil
}

func uncover() error {
	fmt.Printf("Uncovering %s\n", fileToUncover)
	data, err := os.ReadFile(fileToUncover)
	if err != nil {
		return err
	}

	dir, filename := filepath.Split(fileToUncover)
	newFilepath := filepath.Join(dir, filename[:len(filename)-len(suffix)])

	payload := data[len(carrier):]
	f, err := os.Create(newFilepath)
	if err != nil {
		return err
	}

	_, err = f.Write(payload)
	if err != nil {
		return err
	}

	return nil
}
