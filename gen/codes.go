package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	path      = "errcodes.csv"
	startCode = 10000
	comma     = ';'
)

type Gen struct {
	StartCode int
	p         *Parser
}

func NewGen(startCode int, p *Parser) *Gen {
	return &Gen{
		StartCode: startCode,
		p:         p,
	}
}

func (g *Gen) GenerateJS(outPath string) error {
	template := fmt.Sprintf(`export default {
%s
}
`, g.JSObject())

	if err := g.WriteFile(outPath, []byte(template)); err != nil {
		return err
	}

	return nil
}

func (g *Gen) GenerateGO(outPath string) error {
	template := fmt.Sprintf(`package errcodes
var (
	codeMessages = [...]string{
%s
	}
	codesHTTP = [...]int{
%s
	}
)

const (
	startCode = 10000

	ErrUnknown CodeError = iota - 2
	ErrNil
)

const (
%s
)

type CodeError int

func (e CodeError) Error() string {
	code := e - startCode
	if code < 0 || int(code) >= len(codeMessages) {
		return "Неизвестная ошибка"
	}

	return codeMessages[code]
}

func GetHTTPCode(e CodeError) int {
	code := e - startCode
	if code < 0 || int(code) >= len(codesHTTP) {
		return 500
	}

	return codesHTTP[code]
}
	`, g.GoErrText(), g.GoHTTPCodes(), g.GoVariables())

	if err := g.WriteFile(outPath, []byte(template)); err != nil {
		return err
	}

	return nil
}

func (g *Gen) GoVariables() string {
	result := fmt.Sprintf("\t%s CodeError = iota + %d", g.p.VarNames[0], startCode)
	for i := 1; i < len(g.p.VarNames); i++ {
		result = fmt.Sprintf("%s\n\t%s", result, g.p.VarNames[i])
	}

	return result
}

func (g *Gen) GoErrText() string {
	result := ""

	for i, text := range g.p.ErrTextBack {
		result += fmt.Sprintf("\t\t"+`%d: "%s",`+"\n", i, text)
	}

	return result
}

func (g *Gen) GoHTTPCodes() string {
	result := ""

	for i, text := range g.p.HTTPCodes {
		result += fmt.Sprintf("\t\t"+`%d: %s,`+"\n", i, text)
	}

	return result
}

func (g *Gen) JSObject() string {
	result := ""

	for i, text := range g.p.ErrTextFront {
		result += fmt.Sprintf("\t"+`%d: "%s",`+"\n", g.StartCode+i, text)
	}

	return result
}

func (g *Gen) WriteFile(path string, content []byte) error {
	fOut, err := os.Create(path)
	if err != nil {
		log.Printf("failed to create file: %v", err)
		return err
	}

	if _, err = fOut.Write(content); err != nil {
		log.Printf("failed to write file: %v", err)
		return err
	}

	return nil
}

type Parser struct {
	r                                              *csv.Reader
	VarNames, ErrTextBack, ErrTextFront, HTTPCodes []string
}

func NewParser(r *csv.Reader) *Parser {
	r.Comma = comma

	return &Parser{
		r:            r,
		VarNames:     make([]string, 0, 0),
		ErrTextBack:  make([]string, 0, 0),
		ErrTextFront: make([]string, 0, 0),
		HTTPCodes:    make([]string, 0, 0),
	}
}

func (p *Parser) ReadColumns() {
	for {
		rec, err := p.r.Read()
		if err != nil {
			break
		}

		p.VarNames = append(p.VarNames, strings.Trim(rec[0], " "))
		p.ErrTextBack = append(p.ErrTextBack, strings.Trim(rec[1], " "))
		p.ErrTextFront = append(p.ErrTextFront, strings.Trim(rec[2], " "))
		p.HTTPCodes = append(p.HTTPCodes, strings.Trim(rec[3], " "))
	}
}

func main() {
	f, err := os.Open(path)
	if err != nil {
		log.Printf("Failed to open file: %v", err)
		return
	}
	defer func(f *os.File) {
		if err := f.Close(); err != nil {
			log.Printf("Failed to close file: %v", err)
		}
	}(f)

	r := csv.NewReader(f)

	p := NewParser(r)
	p.ReadColumns()

	g := NewGen(startCode, p)

	if err := g.GenerateGO("codes.go"); err != nil {
		log.Printf("failed generate %v", err)
		return
	}
	if err := g.GenerateJS("codes.js"); err != nil {
		log.Printf("failed generate %v", err)
		return
	}

	log.Println("Generate succesful")
}
