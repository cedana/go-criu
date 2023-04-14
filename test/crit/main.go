package main

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/checkpoint-restore/go-criu/v6/crit"
)

const TEST_IMG_DIR = "test-imgs"

func main() {
	// Get list of image files
	imgs, err := getImgs()
	if err != nil {
		log.Fatal(err)
	}
	// Run recode test
	if err = recodeImgs(imgs); err != nil {
		log.Fatal(err)
	}
	log.Println("=== PASS")
}

func recodeImgs(imgs []string) error {
	for _, img := range imgs {
		log.Println("===", img)
		testImg := fmt.Sprintf("%s.test.img", img)
		c := crit.New(img, testImg, "", false, false)
		// Decode the binary image file
		decodedImg, err := c.Decode()
		if err != nil {
			return errors.New(fmt.Sprint("[DECODE]: ", err))
		}
		// Encode it into test binary image file
		if err = c.Encode(decodedImg); err != nil {
			return errors.New(fmt.Sprint("[ENCODE]: ", err))
		}
		// Open and compare original and test files
		imgBytes, err := os.ReadFile(img)
		if err != nil {
			return err
		}
		testImgBytes, err := os.ReadFile(testImg)
		if err != nil {
			return err
		}
		if !bytes.Equal(imgBytes, testImgBytes) {
			return errors.New("[RECODE]: Files do not match")
		}
	}

	return nil
}

func getImgs() ([]string, error) {
	// Certain image files generated by CRIU do not
	// use the protobuf format and contain raw binary
	// data. Some image files are also generated using
	// external tools (ifaddr, route, tmpfs). As these
	// images cannot be processed by CRIT, they are
	// excluded from the tests.
	skipImgs := []string{
		"pages-",
		"pages-shmem-",
		"iptables-",
		"ip6tables-",
		"nftables-",
		"route-",
		"route6-",
		"ifaddr-",
		"tmpfs-",
		"tmpfs-dev-",
		"autofs-",
		"netns-ct-",
		"netns-exp-",
		"rule-",
	}
	// "*.test.img", "*.json.img" or "tmp.*.img" files
	// must be skipped as they are generated by tests
	criuImg := regexp.MustCompile(`^[^\.]*\.img$`)
	dir, err := filepath.Glob(fmt.Sprintf("%s/*.img", TEST_IMG_DIR))
	if err != nil {
		return nil, err
	}
	var imgs []string

nextFile:
	for _, file := range dir {
		if filepath.Ext(file) == ".img" {
			if !criuImg.MatchString(file) {
				continue
			}
			for _, skip := range skipImgs {
				if strings.HasPrefix(filepath.Base(file), skip) {
					continue nextFile
				}
			}
			imgs = append(imgs, file)
		}
	}

	return imgs, nil
}
