package cmd

import (
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
	dir, err := os.Open(directoryPath)
	errcheck.Check(err)

	subNodes, err := dir.ReadDir(0)
	errcheck.Check(err)

	for _, subNode := range subNodes {
		fmt.Printf("%+v\n", subNode)
	}
}
