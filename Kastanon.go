package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

type CTSToken struct {
	PassageID, TokenID, Form string
}

type Alignment struct {
	AlignmentID, TokenID string
}

func trimStringFromDot(s string) string {
	if strings.Contains(s, ".") {
		return s[:strings.LastIndex(s, ".")]
	}
	return s
}

var re = regexp.MustCompile("(?m)[\r\n]*^//.*$")

var ctsBackend = []CTSToken{}
var alignmentBackend = []Alignment{}

func main() {
	cexfile := "output.cex"
	message := "LoadCEX: loading " + cexfile
	log.Println(message)
	cexdata, err := ioutil.ReadFile(cexfile)
	if err != nil {
		panic(err)
	}
	str := string(cexdata)
	relations := strings.Split(str, "#!relations")[1]
	relations = strings.Split(relations, "#!")[0]
	readRelations(relations)
	log.Println("Read !#relations succesfully!")
	// ctsCatalog := strings.Split(str, "#!ctscatalog")[1]
	// ctsCatalog = strings.Split(ctsCatalog, "#!")[0]
	ctsdata := strings.Split(str, "#!ctsdata")[1]
	ctsdata = strings.Split(ctsdata, "#!")[0]
	readCTSData(ctsdata)
	log.Println("Read !#ctsdata succesfully!")
	log.Println("Starting Kastanon Shell")
	shellreader := bufio.NewReader(os.Stdin)
	fmt.Println("---------------------")
	fmt.Println("Kastanon Shell")
	fmt.Println("---------------------")
	for {
		fmt.Println()
		fmt.Print("-> ")
		userinput, _ := shellreader.ReadString('\n')
		commands := strings.Split(userinput, " ")
		command := strings.TrimSpace(commands[0])
		parameter := ""
		if len(commands) > 2 {
			fmt.Println()
			fmt.Println("invalid input!")
			continue
		}
		if len(commands) == 2 {
			parameter = strings.TrimSpace(commands[1])
		}
		switch command {
		case "alignment":
			if parameter == "" {
				fmt.Println()
				fmt.Println("invalid input!")
			} else {
				findAlignment(parameter)
			}
		case "token":
			if parameter == "" {
				fmt.Println()
				fmt.Println("invalid input!")
			} else {
				findToken(parameter)
			}
		case "passages":
			if parameter == "" {
				fmt.Println()
				fmt.Println("invalid input!")
			} else {
				findPassages(parameter)
			}
		case "passage":
			if parameter == "" {
				fmt.Println()
				fmt.Println("invalid input!")
			} else {
				findPassage(parameter)
			}
		case "q":
			fmt.Println()
			fmt.Println("bye!")
			os.Exit(0)
		default:
			fmt.Println()
			fmt.Println("invalid this!")
		}
	}

	// fmt.Println(ctsBackend[0:5])
}

func readRelations(relations string) {
	relations = re.ReplaceAllString(relations, "")
	reader := csv.NewReader(strings.NewReader(relations))
	reader.Comma = '#'
	reader.LazyQuotes = true
	reader.FieldsPerRecord = 3

	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		if strings.Contains(line[1], "aligns") {
			alignmentBackend = append(alignmentBackend, Alignment{AlignmentID: line[0], TokenID: line[2]})
		}
	}
}

func readCTSData(ctsdata string) {
	ctsdata = re.ReplaceAllString(ctsdata, "")
	reader := csv.NewReader(strings.NewReader(ctsdata))
	reader.Comma = '#'
	reader.LazyQuotes = true
	reader.FieldsPerRecord = -1
	// reader.TrimLeadingSpace = true
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			fmt.Println(line)
			log.Fatal(error)
		}
		ctstoken := CTSToken{}
		switch {
		case len(line) == 2:
			ctstoken.TokenID = line[0]
			ctstoken.PassageID = trimStringFromDot(line[0])
			ctstoken.Form = line[1]
			ctsBackend = append(ctsBackend, ctstoken)
		case len(line) > 2:
			ctstoken.TokenID = line[0]
			ctstoken.PassageID = trimStringFromDot(line[0])
			var textstring string
			for j := 1; j < len(line); j++ {
				textstring = textstring + line[j]
			}
			ctstoken.Form = textstring
			ctsBackend = append(ctsBackend, ctstoken)
		case len(line) < 2:
			log.Println("Wrong line:", line)
		}
	}
}

func findPassages(passagesuffix string) {
	tmpID := ""
	for _, v := range ctsBackend {
		if strings.HasSuffix(v.PassageID, passagesuffix) {
			if tmpID != v.PassageID {
				fmt.Println()
				fmt.Println("-------------------------")
				fmt.Println(v.PassageID)
			}
			tmpID = v.PassageID
			fmt.Print(v.Form)
		}
	}
	fmt.Println()
	fmt.Println("-------------------------")
}

func findPassage(passageid string) {
	fmt.Println()
	fmt.Println("-------------------------")
	for _, v := range ctsBackend {
		if passageid == v.PassageID {
			fmt.Print(v.Form)
		}
	}
	fmt.Println()
	fmt.Println("-------------------------")
}

func findAlignment(alignmentid string) {
	fmt.Println()
	fmt.Println("-------------------------")
	fmt.Println(alignmentid)
	tokenIDs := []string{}
	for _, v := range alignmentBackend {
		if alignmentid == v.AlignmentID {
			tokenIDs = append(tokenIDs, v.TokenID)
		}
	}
	for _, v := range tokenIDs {
		for _, v2 := range ctsBackend {
			if v == v2.TokenID {
				fmt.Println(v2.TokenID, v2.Form)
			}
		}
	}
	fmt.Println()
	fmt.Println("-------------------------")
}

func findToken(tokenid string) {
	fmt.Println()
	fmt.Println("-------------------------")
	for _, v2 := range ctsBackend {
		if tokenid == v2.TokenID {
			fmt.Println("Searching for", v2.TokenID, v2.Form)
		}
	}
	fmt.Println("Aligned with,", tokenid, "are:")
	tokenIDs := []string{}
	alignmentid := ""
	for _, v := range alignmentBackend {
		if tokenid == v.TokenID {
			alignmentid = v.AlignmentID
		}
	}
	for _, v := range alignmentBackend {
		if alignmentid == v.AlignmentID && tokenid != v.TokenID {
			tokenIDs = append(tokenIDs, v.TokenID)
		}
	}
	for _, v := range tokenIDs {
		for _, v2 := range ctsBackend {
			if v == v2.TokenID {
				fmt.Println(v2.TokenID, v2.Form)
			}
		}
	}
	fmt.Println()
	fmt.Println("-------------------------")
}
