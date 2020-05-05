package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/yishanhe/animalcrossing-api/models"
	"github.com/yishanhe/animalcrossing-api/pkg/database"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/sheets/v4"
)

var (
	headerMap = make(map[string][]string)
)

func init() {
	headerMap["Fish"] = []string{"#", "Name", "Icon Image", "Critterpedia Image", "Furniture Image", "Sell", "Where/How", "Shadow", "Total Catches to Unlock", "Rain/Snow Catch Up", "NH Jan", "NH Feb", "NH Mar", "NH Apr", "NH May", "NH Jun", "NH Jul", "NH Aug", "NH Sep", "NH Oct", "NH Nov", "NH Dec", "SH Jan", "SH Feb", "SH Mar", "SH Apr", "SH May", "SH Jun", "SH Jul", "SH Aug", "SH Sep", "SH Oct", "SH Nov", "SH Dec", "Color 1", "Color 2", "Size", "Lighting Type", "Icon Filename", "Critterpedia Filename", "Furniture Filename", "Internal ID", "Unique Entry ID"}
	headerMap["Bugs"] = []string{"#", "Name", "Icon Image", "Critterpedia Image", "Furniture Image", "Sell", "Where/How", "Weather", "Total Catches to Unlock", "NH Jan", "NH Feb", "NH Mar", "NH Apr", "NH May", "NH Jun", "NH Jul", "NH Aug", "NH Sep", "NH Oct", "NH Nov", "NH Dec", "SH Jan", "SH Feb", "SH Mar", "SH Apr", "SH May", "SH Jun", "SH Jul", "SH Aug", "SH Sep", "SH Oct", "SH Nov", "SH Dec", "Color 1", "Color 2", "Icon Filename", "Critterpedia Filename", "Furniture Filename", "Internal ID", "Unique Entry ID"}
}

// Retrieve a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config) *http.Client {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	tokFile := "token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

func main() {

	b, err := ioutil.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, "https://www.googleapis.com/auth/spreadsheets.readonly")
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := getClient(config)

	srv, err := sheets.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	readBugSheet(srv)

}

func readBugSheet(srv *sheets.Service) {
	spreadsheetID := "13d_LAJPlxMa_DubPTuirkIV4DERBMXbrWQsmSh8ReK4"
	readRange := "Bugs"
	resp, err := srv.Spreadsheets.Values.Get(spreadsheetID, readRange).ValueRenderOption("FORMULA").Do()
	if err != nil {
		log.Fatalf("Unable to retrieve Sheet %s %v", readRange, err)
	}
	coll := database.NewMongoClient().Database("AnimalCrossingDevDB").Collection("bug")
	headerProcessed := false

	if len(resp.Values) == 0 {
		fmt.Println("No data found.")
	} else {
		for _, row := range resp.Values {
			if !headerProcessed {
				headerProcessed = true
				validateHeader(row, "Bugs")
				continue
			}
			sellPriceInFloat, ok := row[5].(float64)
			if !ok {
				log.Fatal("Type not right")
			}
			catchesToUnlockInFloat, ok := row[8].(float64)
			if !ok {
				log.Fatal("Type not right")
			}
			idInFloat, ok := row[38].(float64)
			if !ok {
				log.Fatal("Type not right")
			}

			months, hours := extractActiveTime(row, 9, 21)
			images := []*models.Image{}
			images = append(images, &models.Image{
				ImageType:     "Icon Image",
				ImageFilename: fmt.Sprintf("%s", row[35]),
				ImageURL:      strings.TrimSuffix(strings.TrimPrefix(fmt.Sprintf("%s", row[2]), "=IMAGE(\""), "\")"),
			})
			images = append(images, &models.Image{
				ImageType:     "Critterpedia Image",
				ImageFilename: fmt.Sprintf("%s", row[36]),
				ImageURL:      fmt.Sprintf("%s", row[3]),
			})
			images = append(images, &models.Image{
				ImageType:     "Furniture Image",
				ImageFilename: fmt.Sprintf("%s", row[37]),
				ImageURL:      fmt.Sprintf("%s", row[4]),
			})

			// color 33 34
			colors := []string{
				fmt.Sprintf("%s", row[33]), fmt.Sprintf("%s", row[34]),
			}

			var ele *models.Bug
			ele = &models.Bug{
				ID:      int64(idInFloat),
				EntryID: fmt.Sprintf("%s", row[39]),
				Name: &models.Name{
					NameEn: fmt.Sprintf("%s", row[1]),
				},
				SellPrice:       int64(sellPriceInFloat),
				Months:          &months,
				Hours:           hours,
				Location:        fmt.Sprintf("%s", row[6]),
				CatchesToUnlock: int64(catchesToUnlockInFloat),
				Weather:         fmt.Sprintf("%s", row[7]),
				Images:          images,
				Colors:          colors,
			}

			data, err := bson.Marshal(ele)
			if err != nil {
				continue
			}
			// utils.PrettyPrint(ele)
			coll.InsertOne(context.Background(), data)

		}
	}
}

func readFishSheet(srv *sheets.Service) {

	spreadsheetID := "13d_LAJPlxMa_DubPTuirkIV4DERBMXbrWQsmSh8ReK4"
	readRange := "Fish"
	resp, err := srv.Spreadsheets.Values.Get(spreadsheetID, readRange).ValueRenderOption("FORMULA").Do()
	if err != nil {
		log.Fatalf("Unable to retrieve Sheet %s %v", readRange, err)
	}
	// a, err := resp.MarshalJSON()
	// fmt.Println(string(a))

	coll := database.NewMongoClient().Database("AnimalCrossingDevDB").Collection("fish")

	// TODO: create inedex
	headerProcessed := false
	if len(resp.Values) == 0 {
		fmt.Println("No data found.")
	} else {
		for _, row := range resp.Values {
			if !headerProcessed {
				headerProcessed = true
				validateHeader(row, "Fish")
				continue
			}
			sellPriceInFloat, ok := row[5].(float64)
			if !ok {
				log.Fatal("Type not right")
			}
			catchesToUnlockInFloat, ok := row[8].(float64)
			if !ok {
				log.Fatal("Type not right")
			}
			idInFloat, ok := row[41].(float64)
			if !ok {
				log.Fatal("Type not right")
			}

			months, hours := extractActiveTime(row, 10, 22)
			images := []*models.Image{}
			images = append(images, &models.Image{
				ImageType:     "Icon Image",
				ImageFilename: fmt.Sprintf("%s", row[38]),
				ImageURL:      strings.TrimSuffix(strings.TrimPrefix(fmt.Sprintf("%s", row[2]), "=IMAGE(\""), "\")"),
			})
			images = append(images, &models.Image{
				ImageType:     "Critterpedia Image",
				ImageFilename: fmt.Sprintf("%s", row[39]),
				ImageURL:      fmt.Sprintf("%s", row[3]),
			})
			images = append(images, &models.Image{
				ImageType:     "Furniture Image",
				ImageFilename: fmt.Sprintf("%s", row[40]),
				ImageURL:      fmt.Sprintf("%s", row[4]),
			})

			// color 34 35
			colors := []string{
				fmt.Sprintf("%s", row[34]), fmt.Sprintf("%s", row[35]),
			}

			var ele *models.Fish
			ele = &models.Fish{
				ID:      int64(idInFloat),
				EntryID: fmt.Sprintf("%s", row[42]),
				Name: &models.Name{
					NameEn: fmt.Sprintf("%s", row[1]),
				},
				SellPrice:       int64(sellPriceInFloat),
				Months:          &months,
				Hours:           hours,
				Location:        fmt.Sprintf("%s", row[6]),
				Shadow:          fmt.Sprintf("%s", row[7]),
				CatchesToUnlock: int64(catchesToUnlockInFloat),
				Weather:         fmt.Sprintf("Rain/Snow needed? %s", row[9]),
				Images:          images,
				Size:            fmt.Sprintf("%s", row[36]),
				Colors:          colors,
				LightingType:    fmt.Sprintf("%s", row[37]),
			}

			data, err := bson.Marshal(ele)
			if err != nil {
				continue
			}
			// utils.PrettyPrint(ele)
			coll.InsertOne(context.Background(), data)

		}

	}
}

func validateHeader(row []interface{}, sheet string) {
	headers, ok := headerMap[sheet]
	if ok {
		for i, val := range row {
			if val != headers[i] {
				log.Fatal(val, " is not matching ", headers[i])
				return
			}
		}
	}
}

func extractActiveTime(row []interface{}, nhStartIdx int, shStartIdx int) (models.Months, string) {
	var hour string
	var northern []int64
	var southern []int64
	northern = make([]int64, 0)
	southern = make([]int64, 0)
	for i := 0; i < 12; i++ {
		nh := fmt.Sprintf("%s", row[i+nhStartIdx])
		sh := fmt.Sprintf("%s", row[i+shStartIdx])
		if nh != "NA" && nh != "No" {
			northern = append(northern, int64(i+1))
			hour = string(nh)
		}
		if sh != "NA" && sh != "No" {
			southern = append(southern, int64(i+1))
		}
	}

	months := models.Months{
		Northern: northern,
		Southern: southern,
	}

	return months, hour
}

func printHeader(row []interface{}) {
	for _, header := range row {
		fmt.Printf("\"%s\",", header)
	}
}
