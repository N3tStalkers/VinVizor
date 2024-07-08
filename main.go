package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	color "github.com/fatih/color"

)

type VehicleInfo struct {
	VIN                 string `json:"VIN"`
	Make                string `json:"Make"`
	Model               string `json:"Model"`
	ModelYear           string `json:"ModelYear"`
	BodyClass           string `json:"BodyClass"`
	DriveType           string `json:"DriveType"`
	EngineConfiguration string `json:"EngineConfiguration"`
	EngineCylinders     string `json:"EngineCylinders"`
	FuelTypePrimary     string `json:"FuelTypePrimary"`
}

func vinLookup(vin string) (*VehicleInfo, error) {
	url := "https://vpic.nhtsa.dot.gov/api/vehicles/DecodeVinValuesBatch/"
	payload := strings.NewReader("format=json&data=" + vin)

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data struct {
		Results []VehicleInfo `json:"Results"`
	}
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	if len(data.Results) == 0 {
		return nil, fmt.Errorf("VIN not found or no results returned")
	}

	return &data.Results[0], nil
}

func printVehicleInfo(info *VehicleInfo) {
	yellow := color.New(color.FgYellow).SprintFunc()
	cyan := color.New(color.FgCyan).SprintFunc()
	fmt.Printf("VIN: %s\n", yellow(info.VIN))
	fmt.Printf("Make: %s\n", cyan(info.Make))
	fmt.Printf("Model: %s\n", cyan(info.Model))
	fmt.Printf("Model Year: %s\n", cyan(info.ModelYear))
	fmt.Printf("Body Class: %s\n", cyan(info.BodyClass))
	fmt.Printf("Drive Type: %s\n", cyan(info.DriveType))
	fmt.Printf("Engine Configuration: %s\n", cyan(info.EngineConfiguration))
	fmt.Printf("Engine Cylinders: %s\n", cyan(info.EngineCylinders))
	fmt.Printf("Fuel Type: %s\n", cyan(info.FuelTypePrimary))
}

func main() {
	fmt.Println(`
  #  #     #           #  #     #                        
  #  #                 #  #                              
  #  #    ##    ###    #  #    ##    ####    ##    # #   
  ####     #    #  #   ####     #      ##   #  #   ## #  
   ##      #    #  #    ##      #    ##     #  #   #     
   ##     ###   #  #    ##     ###   ####    ##    #     
                                                        
                                                        
	`)
	var vin string
	fmt.Print("Enter VIN : ")
	fmt.Scanln(&vin)

	info, err := vinLookup(vin)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	printVehicleInfo(info)
}
