package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type bingMetaData struct {
	Images []oneImage `json:"images"`
}

type oneImage struct {
	StartDate string `json:"startdate"`
	URL       string `json:"url"`
	Title     string `json:"title"`
}
const (
	baseURL = "http://bing.com"
	bingAPI = "https://www.bing.com/HPImageArchive.aspx?format=js&idx=%s&n=1&mkt=%s"
)

var (
	resolutionMap map[string]string
	markets       map[string]bool
	myClient	  *http.Client
)

const jpegExt = ".jpeg"

func init() {
	rootCmd.AddCommand(getWallpaperCmd)

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

// TODO: Make this extensible with flags supporting various options, such as mkt for market, date for getting specific
// TODO: day, etc
var getWallpaperCmd = &cobra.Command{
	Use:   "get-wallpaper",
	Short: "Downloads Bing daily wallpaper",
	Long:  `This downloads the latest Bing Wallpaper of the day to user specified directory.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Getting Bing daily wallpaper data")
		metaData := getBingWallpaperMeta()
		success, filename := downloadWallpaper(metaData, state.ImageDir)
		if !success {
			fmt.Println("Uh oh!! Could not save Image file.")
		}
		state.LatestImage = filename
		writeState(state)
	},
}

func getBingWallpaperMeta() bingMetaData {
	//TODO: hardcoded config, add format also
	index := 1
	mkt := "en-US"
	// Get Image URL from Bing API.
	bingMeta := new(bingMetaData)
	err := getJson(fmt.Sprintf(bingAPI, index, mkt), bingMeta)
	if err != nil {
		log.Printf("Decoding Json Error : %v \n", err)
		return *bingMeta
	}
	return *bingMeta
}

func downloadWallpaper(metaData bingMetaData, dirPath string) (bool, string) {
	// Select first image TODO: check for array length before slicing
	imageURL := baseURL + metaData.Images[0].URL
	response, err := http.Get(imageURL)
	if err != nil {
		log.Printf("Error in GET request : %v \n", err)
		return false, ""
	}
	defer response.Body.Close()

	imageFile := strings.Replace(strings.ToLower(metaData.Images[0].Title), " ", "_", -1)
	imageFile = strings.Replace(imageFile, "’", "", -1)
	imageFile = strings.Replace(imageFile, "!", "", -1)

	fmt.Println(dirPath + "/" + imageFile + jpegExt)
	fmt.Println("Created above file")
	file, err := os.Create(dirPath + "/" + imageFile + jpegExt) // Perm : 0666
	//fmt.Println(dirPath + imageFile + ".jpeg")
	if err != nil {
		log.Printf("Error creating file : %v \n", err)
		return false, ""
	}
	defer file.Close()

	// Use io.Copy to just dump the response body to the file. TODO: check if it supports large files
	b, err := io.Copy(file, response.Body)
	if err != nil {
		log.Printf("Error copying image response to file : %v \n", err)
		return false, ""
	}
	fmt.Println("File name: ", imageFile + ".jpeg")
	fmt.Println("File size: ", b)
	fmt.Println("Directory Path: ", dirPath)
	return true, (imageFile+jpegExt)
}


func getJson(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		log.Printf("Error in GET request : %v \n", err)
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}