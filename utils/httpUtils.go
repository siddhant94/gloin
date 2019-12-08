package utils

import (
	"bing-wod/logging"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	baseURL = "http://bing.com"
	bingAPI = "https://www.bing.com/HPImageArchive.aspx?format=js&idx=%s&n=1&mkt=%s"
)

var (
	resolutionMap map[string]string
	markets       map[string]bool
	myClient	  *http.Client
)

//"https://www.bing.com/HPImageArchive.aspx?format=js&idx=0&n=1&mkt=en-US"

func init() {
	myClient = &http.Client{Timeout: 10 * time.Second}

	resolutionMap = map[string]string{}
	resolutionMap["1366"] = "1366x768.jpg"
	resolutionMap["1920"] = "1920x1080.jpg"

	markets = map[string]bool{
		"en-US": true,
		"zh-CN": true,
		"ja-JP": true,
		"en-AU": true,
		"en-UK": true,
		"de-DE": true,
		"en-NZ": true,
		"en-CA": true,
	}
}
type bingMetaData struct {
	Images []oneImage `json:"images"`
}

type oneImage struct {
	StartDate string `json:"startdate"`
	URL       string `json:"url"`
	Title     string `json:"title"`
}

func GetBingWallpaperMeta() bingMetaData {
	//TODO: hardcoded config, add format also
	index := 1
	mkt := "en-US"
	// Get Image URL from Bing API.
	bingMeta := new(bingMetaData)
	err := getJson(fmt.Sprintf(bingAPI, index, mkt), bingMeta)
	if err != nil {
		logging.Logger.Printf("Decoding Json Error : %v \n", err)
		return *bingMeta
	}
	return *bingMeta
}

func DownloadWallpaper(metaData bingMetaData, dirPath string) (bool, string) {
	// Select first image TODO: check for array length before slicing
	imageURL := baseURL + metaData.Images[0].URL
	response, err := http.Get(imageURL)
	if err != nil {
		logging.Logger.Printf("Error in GET request : %v \n", err)
		return false, ""
	}
	defer response.Body.Close()

	//open a file for writing
	imageFile := strings.Replace(strings.ToLower(metaData.Images[0].Title), " ", "_", -1)
	file, err := os.Create(dirPath + "/" + imageFile + ".jpeg")
	fmt.Println(dirPath + imageFile + ".jpeg")
	if err != nil {
		logging.Logger.Printf("Error creating file : %v \n", err)
		return false, ""
	}
	defer file.Close()

	// Use io.Copy to just dump the response body to the file. TODO: check if it supports large files
	b, err := io.Copy(file, response.Body)
	if err != nil {
		logging.Logger.Printf("Error copying image response to file : %v \n", err)
		return false, ""
	}
	fmt.Println("File size: ", b)
	return true, imageFile
}

func getJson(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		logging.Logger.Printf("Error in GET request : %v \n", err)
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}
