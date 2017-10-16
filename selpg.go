package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

var (
	flagSet    = flag.NewFlagSet(os.Args[0], flag.PanicOnError)
	start_page = flagSet.Int("s", -1, "read from <s> page")
	end_page   = flagSet.Int("e", -1, "read until <e> page")
	page_len   = flagSet.Int("l", 72, "lines per page(default:72 lines/page)")
	fin        = flagSet.Bool("f", false, "read one page until '\f' ")
)

func process_args() bool {
	if len(os.Args) <= 2 {
		err := errors.New("The command need both start_page:-s=number and  end_page:-e=number")
		fmt.Fprintln(os.Stderr, "warning(command format): ", err)
		return false
	}
	if os.Args[1][0:2] != "-s" {
		err := errors.New("The command should be like as -s=number -e=number [options] [filename]")
		fmt.Fprintln(os.Stderr, "warning(command format): ", err)
		return false
	}
	if *start_page <= 0 {
		err := errors.New("The start_page can not be less than 1")
		fmt.Fprintln(os.Stderr, "warning(arguments): ", err)
		return false
	}
	if os.Args[2][0:2] != "-e" {
		err := errors.New("The command should be like as -s=number -e=number [options] [filename]")
		fmt.Fprintln(os.Stderr, "warning(command format): ", err)
		return false
	}
	if *end_page <= 0 {
		err := errors.New("The end_page can not be less than 1")
		fmt.Fprintln(os.Stderr, "warning(arguments): ", err)
		return false
	}
	if *start_page > *end_page {
		err := errors.New("The end_page can not be less than start_page")
		fmt.Fprintln(os.Stderr, "warning(arguments): ", err)
		return false
	}
	if *page_len <= 0 {
		err := errors.New("The page_line can not be less than 0")
		fmt.Fprintln(os.Stderr, "warning(arguments): ", err)
		return false
	}
	return true
}

func process_putin(Ibuf *bufio.Reader, Obuf *os.File) {
	var count int
	count = *end_page - *start_page + 1
	if !*fin { /*read all the char from the file from the startpage*/
		for i := 1; i < *start_page; i++ {
			for j := 0; j < *page_len; j++ {
				Ibuf.ReadString('\n')
			}
		}
		for i := 0; i < count; i++ {
			for j := 0; j < *page_len; j++ {
				line, err := Ibuf.ReadString('\n')
				if err != nil {
					if err == io.EOF &&
						i != count &&
						j != *page_len {
						err2 := errors.New("The pages in the file is too less to read")
						fmt.Fprintln(os.Stderr, "warning(file reading) ", err2)
						return
					} else {
						fmt.Fprint(os.Stderr, "warning(file reading) ", err.Error())
					}
				}
				if Obuf != nil {
					Obuf.WriteString(line)
				} else {
					fmt.Print(line)
				}
			}
		}
	} else { /*the cut of the page*/
		for i := 1; i < *start_page; i++ {
			Ibuf.ReadString('\f')
		}
		for i := 0; i < count; i++ {
			line, err := Ibuf.ReadString('\f')
			if err != nil {
				if err == io.EOF && i != count {
					err3 := errors.New("The pages in the file is too less to read")
					fmt.Fprintln(os.Stderr, "warning(file reading) ", err3)
					return
				} else {
					fmt.Fprint(os.Stderr, "warning(file reading) ", err.Error())
				}
			}
			if Obuf != nil {
				Obuf.WriteString(line)
			} else {
				fmt.Print(line)
			}
		}
	}
}


func write(In string, Ou string) {
	var Ibuf *bufio.Reader
	if In != "" {
		inFile, err := os.OpenFile(In, os.O_RDWR, 0644)
		if err != nil {
			fmt.Fprintln(os.Stderr, "warning(open file) ", err.Error())
		}
		Ibuf = bufio.NewReader(inFile)
	} else {
		Ibuf = bufio.NewReader(os.Stdin)
	}
	var Obuf *os.File
	var err error
	if Ou != "" {
		Obuf, err = os.OpenFile(Ou, os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0644)
		if err != nil {
			fmt.Fprintln(os.Stderr, "warning(open file) ", err.Error())
		}
	} else {
		Obuf = nil
	}
	process_putin(Ibuf, Obuf)
}



func main() {
	flagSet.Parse(os.Args[1:])
	if process_args() {
		var inputFile string
		var outputFile string
		if flagSet.NArg() > 0 {
			inputFile = flagSet.Arg(0)
		}
		if flagSet.NArg() > 1 {
			outputFile = flagSet.Arg(1)
		}
		write(inputFile, outputFile)
	}
}
