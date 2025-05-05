package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	Parse(data string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	for record := range dataset {
		if dp.Parse(dataset[record]) != nil {
			log.Printf("Error parsing record %d: %v", record, dp.Parse(dataset[record]))
			continue
		}

		info, err := dp.ActionInfo()
		if err != nil {
			log.Printf("Error parsing record %s: %v", info, err)
			continue
		}

		fmt.Println(info)
	}
}
