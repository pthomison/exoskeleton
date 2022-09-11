package cmd

import (
	"bytes"
	"errors"
	"fmt"
	"os"

	"github.com/pthomison/errcheck"
	"gopkg.in/yaml.v3"
)

type UnstructureYamlData map[interface{}]interface{}

func readYamlFile(filepath string) UnstructureYamlData {
	yamlBytes, err := os.ReadFile(filepath)
	errcheck.Check(err)

	yamlData := make(UnstructureYamlData)

	err = yaml.Unmarshal(yamlBytes, &yamlData)
	errcheck.Check(err)

	return yamlData
}

func recurseDir(dir *os.File, callback func(leafNode *os.File, outputLocation string, variableData UnstructureYamlData)) {

	subNodes, err := dir.ReadDir(0)
	errcheck.Check(err)

	for _, subNode := range subNodes {
		fmt.Printf("%+v\n", subNode)
	}
}

func ls(directoryPath string) {
	dir := loadFilepath(directoryPath)

	subNodes, err := dir.ReadDir(0)
	errcheck.Check(err)

	for _, subNode := range subNodes {
		fmt.Printf("%+v\n", subNode)
	}
}

func loadFilepath(filepath string) *os.File {
	file, err := os.Open(filepath)
	errcheck.Check(err)

	return file
}

func readFile(file *os.File) []byte {
	var fileBytes []byte

	n, err := file.Read(fileBytes)
	errcheck.Check(err)

	if n != 0 {
		errcheck.Check(errors.New("Unable to read entire file"))
	}

	return fileBytes
}

func readFilepath(filepath string) []byte {
	return readFile(loadFilepath(filepath))
}

func indexDirEntrys(content []os.DirEntry) map[string]os.DirEntry {
	index := make(map[string]os.DirEntry)

	for _, v := range content {
		index[v.Name()] = v
	}

	return index
}

func CompareFolders(folderAlpha string, folderBeta string) bool {
	dirAlpha := loadFilepath(folderAlpha)
	dirBeta := loadFilepath(folderAlpha)

	return compareFolders(dirAlpha, dirBeta, "/")
}

func compareFolders(alphaNode *os.File, betaNode *os.File, filepath string) bool {
	if alphaNode == nil || betaNode == nil {
		return false
	}

	alphaStat, err := alphaNode.Stat()
	errcheck.Check(err)
	alphaIsDir := alphaStat.IsDir()

	betaStat, err := betaNode.Stat()
	errcheck.Check(err)
	betaIsDir := betaStat.IsDir()

	if alphaIsDir && betaIsDir {
		alphaFiles, err := alphaNode.ReadDir(0)
		errcheck.Check(err)
		alphaIndex := indexDirEntrys(alphaFiles)

		betaFiles, err := betaNode.ReadDir(0)
		errcheck.Check(err)
		betaIndex := indexDirEntrys(betaFiles)

		result := true

		for name, alphaSubnode := range alphaIndex {
			fmt.Printf("%+v\n", alphaSubnode)
			_, _ = name, betaIndex
			// result = result && compareFolders(alphaSubnode, betaIndex[name], filepath+"name/")
		}

		return result

	} else if (alphaIsDir && !betaIsDir) || (!alphaIsDir && betaIsDir) {
		return false
	} else if !alphaIsDir && !betaIsDir {
		alphaBytes := readFile(alphaNode)
		betaBytes := readFile(betaNode)

		return bytes.Equal(alphaBytes, betaBytes)
	}

	return false
}
