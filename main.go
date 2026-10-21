package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

// checking for error
func check(err error) {
	if err != nil {
		fmt.Println("Error Happened")
		os.Exit(1)
	}
}
func main() {
	//get the user to input the folder needed to organize
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("which folder do you want to organize?- ")
	scanner.Scan()

	targetFolder := scanner.Text()

	//checking if the folder exists or not
	_, err := os.Stat(targetFolder)
	if err != nil {
		fmt.Print("folder doesn't exist")
	} else {
		createDefaultFolder(targetFolder)

		organizeFolder(targetFolder)
	}
}
func createDefaultFolder(targetFolder string) {
	defaultFolders := []string{"images", "videos", "musics", "docs", "others"}

	for _, folder := range defaultFolders {
		_, err := os.Stat(folder)
		if os.IsNotExist(err) {
			os.Mkdir(filepath.Join(targetFolder, folder), 0755)
		}
	}
}
func organizeFolder(targetFolder string) {
	//read dir
	filesAndFolders, err := os.ReadDir(targetFolder)
	check(err)

	//to track hw many files moved
	noOfFiles := 0

	for _, filesAndFolder := range filesAndFolders {
		//check files
		if !filesAndFolder.IsDir() {
			filesInfo, err := filesAndFolder.Info()
			check(err)

			//get the file full path
			oldPath := filepath.Join(targetFolder, filesInfo.Name())
			fileExt := filepath.Ext(oldPath)

			//switch casea to move files based on theri ext
			switch fileExt {
			case ".png", ".jpg", ".jpeg":
				newPath := filepath.Join(targetFolder, "images", filesInfo.Name())
				err := os.Rename(oldPath, newPath)
				check(err)
				noOfFiles++

			case ".mp4", ".mov", ".avi":
				newPath := filepath.Join(targetFolder, "videos", filesInfo.Name())
				err := os.Rename(oldPath, newPath)
				check(err)
				noOfFiles++

			case ".mp3", ".aac", ".wav":
				newPath := filepath.Join(targetFolder, "musics", filesInfo.Name())
				err := os.Rename(oldPath, newPath)
				check(err)
				noOfFiles++

			case ".csv", ".docx", ".pdf", ".xlsx":
				newPath := filepath.Join(targetFolder, "docs", filesInfo.Name())
				err := os.Rename(oldPath, newPath)
				check(err)
				noOfFiles++

			default:
				newPath := filepath.Join(targetFolder, "images", filesInfo.Name())
				err := os.Rename(oldPath, newPath)
				check(err)
				noOfFiles++
			}

		}
	}
	//print hw many files moved
	if noOfFiles > 0 {
		fmt.Printf("%v numbers of files moved\n", noOfFiles)
	} else {
		fmt.Print("No files moved")
	}
}
