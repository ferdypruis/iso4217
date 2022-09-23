// Generate data structures based on the XML's available at https://www.currency-iso.org/
//go:build ignore
// +build ignore

package main

import (
	"encoding/xml"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"text/template"
	"time"
)

const urlListOne = "https://www.six-group.com/dam/download/financial-information/data-center/iso-currrency/lists/list-one.xml"
const urlListThree = "https://www.six-group.com/dam/download/financial-information/data-center/iso-currrency/lists/list-three.xml"

const tpl = `generator/data.go.tpl`
const outfile = `data.go`

const formatPblshd = "2006-01-02" // yyyy-mm-dd

type currency struct {
	Alpha    string
	Numeric  string
	Exponent int
	Name     string
	Historic bool
}

func timeLatest(times ...time.Time) time.Time {
	var tl time.Time
	for _, t := range times {
		if t.After(tl) {
			tl = t
		}
	}
	return tl
}

func main() {
	// Parse template file
	tmpl, err := template.ParseFiles(tpl)
	if err != nil {
		log.Fatalln(`iso4217: error parsing template %s; %s`, tpl, err)
	}

	// Open output file
	w, err := os.Create(outfile)
	if err != nil {
		log.Fatalln(`iso4217: error opening output file %s; %s`, outfile, err)
	}
	defer w.Close()

	currencies, publishDateListOne := listOne()
	historic, publishDateListThree := listThree()

	// Add historic data
	for code, currency := range historic {
		if _, ok := currencies[code]; ok {
			// Skip historic currency if still on List One
			continue
		}
		currencies[code] = currency
	}

	data := map[string]interface{}{
		"generator":  "from XML published on " + timeLatest(publishDateListOne, publishDateListThree).Format(formatPblshd),
		"currencies": currencies,
	}

	// Render template into outfile
	if err := tmpl.Execute(w, data); err != nil {
		log.Fatalln(`iso4217: error rendering template:`, err)
	}

	log.Printf("iso4217: generated %d currencies", len(currencies))
}

func listOne() (map[string]currency, time.Time) {
	resp, err := http.Get(urlListOne)
	if err != nil {
		log.Fatalln(`iso4217: error downloading list_one:`, err)
	}
	defer resp.Body.Close()

	var published time.Time
	var currencies = make(map[string]currency)

	// Parse list_one
	decoder := xml.NewDecoder(resp.Body)

	for {
		t, err := decoder.Token()
		if t != nil {
			elem, ok := t.(xml.StartElement)
			if !ok {
				continue
			}

			switch elem.Name.Local {
			case "ISO_4217":
				// Look for publish date
				for _, attr := range elem.Attr {
					if attr.Name.Local != "Pblshd" {
						continue
					}
					published, _ = time.Parse(formatPblshd, attr.Value)
				}
			case "CcyNtry":
				var entry XmlCcyNtry
				if err := decoder.DecodeElement(&entry, &elem); err != nil {
					log.Fatalln(`iso4217: error parsing list_one:`, err)
				}

				// Process entry
				if entry.Ccy == "" {
					// Skip currencies without code
					continue
				}

				var c = currency{
					Alpha:   entry.Ccy,
					Numeric: entry.CcyNbr,
					Name:    entry.CcyNm,
				}

				if entry.CcyMnrUnts != "N.A." {
					c.Exponent, _ = strconv.Atoi(entry.CcyMnrUnts)
				} else {
					c.Exponent = 0
				}

				currencies[entry.Ccy] = c
			}
		}

		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalln(`iso4217: error parsing list_one:`, err)
		}
	}

	return currencies, published
}

func listThree() (map[string]currency, time.Time) {
	resp, err := http.Get(urlListThree)
	if err != nil {
		log.Fatalln(`iso4217: error downloading list_three:`, err)
	}
	defer resp.Body.Close()

	var published time.Time
	var currencies = make(map[string]currency)

	// Parse list_one
	decoder := xml.NewDecoder(resp.Body)

	for {
		t, err := decoder.Token()
		if t != nil {
			elem, ok := t.(xml.StartElement)
			if !ok {
				continue
			}

			switch elem.Name.Local {
			case "ISO_4217":
				// Look for publish date
				for _, attr := range elem.Attr {
					if attr.Name.Local != "Pblshd" {
						continue
					}
					published, _ = time.Parse(formatPblshd, attr.Value)
				}
			case "HstrcCcyNtry":
				var entry XmlHstrcCcyNtry
				if err := decoder.DecodeElement(&entry, &elem); err != nil {
					log.Fatalln(`iso4217: error parsing list_three:`, err)
				}

				// Process entry
				if entry.Ccy == "" {
					// Skip currencies without code
					continue
				}

				currencies[entry.Ccy] = currency{
					Alpha:    entry.Ccy,
					Numeric:  entry.CcyNbr,
					Name:     entry.CcyNm,
					Historic: true,
				}
			}
		}

		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalln(`iso4217: error parsing list_three:`, err)
		}
	}

	return currencies, published
}

type XmlCcyNtry struct {
	CtryNm     string `xml:"CtryNm"`
	CcyNm      string `xml:"CcyNm"`
	Ccy        string `xml:"Ccy"`
	CcyNbr     string `xml:"CcyNbr"`
	CcyMnrUnts string `xml:"CcyMnrUnts"`
}

type XmlHstrcCcyNtry struct {
	CtryNm    string `xml:"CtryNm"`
	CcyNm     string `xml:"CcyNm"`
	Ccy       string `xml:"Ccy"`
	CcyNbr    string `xml:"CcyNbr"`
	WthdrwlDt string `xml:"WthdrwlDt"`
}
