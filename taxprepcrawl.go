package main

import (
	"log"
	"fmt"
	"regexp"
  	"github.com/PuerkitoBio/goquery"
  	"encoding/csv"
  	"os"
  	"time"
)

const taxPrepUrl = "http://search.irs.gov/search?q=%d&site=efile&client=efile_frontend&output=xml_no_dtd&proxystylesheet=efile_frontend&filter=0&getfields=*&partialfields=buszip\\%3A%d\\%7Cip_w\\%3A%d&num=1000"
func main(){
	  // Create a csv file
    f, err := os.Create("./taxpayers.csv")
    if err != nil {
        fmt.Println(err)
    }
    defer f.Close()
    // Write Unmarshaled json data to CSV file
    w := csv.NewWriter(f)

    fmt.Println("Downloading Tax Payer Info")

	for zipIter:=100;zipIter<1000;zipIter++{
		fmt.Printf("\nDownloading data for zipcode:%d\n",zipIter)
		taxUrl := fmt.Sprintf(taxPrepUrl, zipIter,zipIter,zipIter)
		doc, err := goquery.NewDocument(taxUrl) 
	  	if err != nil {
	    	log.Fatal(err)
	  	}

	  	doc.Find("table").Each(func(i int, s *goquery.Selection) {
	  		row := []string{}
	    	s.Find("td[width='100\\%']").Each(func (j int, t *goquery.Selection){
	    		if j==2{
	    			re := regexp.MustCompile(fmt.Sprintf("\u00a0%d", zipIter))
					if re.FindString(t.Text())!=""{
						row = append(row, t.Text())
					}
				} else{
					row = append(row, t.Text())
				}
			})	
			if len(row)==5{
				fmt.Printf(".")
				w.Write(row)	
			}
	  	})

	  	//try to be a little polite to the IRS server
	  	time.Sleep(100 * time.Millisecond)
	}

    w.Flush()
	
}