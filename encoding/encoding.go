package encoding

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Yandex-Practicum/final-project-encoding-go/models"
	"gopkg.in/yaml.v3"
)

// JSONData тип для перекодирования из JSON в YAML
type JSONData struct {
	DockerCompose *models.DockerCompose
	FileInput     string
	FileOutput    string
}

// YAMLData тип для перекодирования из YAML в JSON
type YAMLData struct {
	DockerCompose *models.DockerCompose
	FileInput     string
	FileOutput    string
}

// MyEncoder интерфейс для структур YAMLData и JSONData
type MyEncoder interface {
	Encoding() error
}

// Encoding перекодирует файл из JSON в YAML
func (j *JSONData) Encoding() error {
	jsonData, err := os.ReadFile(j.FileInput)
	if err != nil {
		err = fmt.Errorf("ошибка при чтении JSON файла: %s", err.Error())
		fmt.Println(err)
		return err
	}

	err = json.Unmarshal(jsonData, &j.DockerCompose)
	if err != nil {
		err = fmt.Errorf("ошибка при десериализации из JSON: %s", err.Error())
		fmt.Println(err)
		return err
	}

	yamlData, err := yaml.Marshal(&j.DockerCompose)
	if err != nil {
		err = fmt.Errorf("ошибка при сериализации в YAML: %s", err.Error())
		fmt.Println(err)
		return err
	}

	err = os.WriteFile(j.FileOutput, yamlData, 0644)
	if err != nil {
		err = fmt.Errorf("ошибка при записи YAML в файл: %s", err.Error())
		fmt.Println(err)
		return err
	}

	return nil
}

// Encoding перекодирует файл из YAML в JSON
func (y *YAMLData) Encoding() error {
	yamlData, err := os.ReadFile(y.FileInput)
	if err != nil {
		err = fmt.Errorf("ошибка при чтении YAML файла: %s", err.Error())
		fmt.Println(err)
		return err
	}

	err = yaml.Unmarshal(yamlData, &y.DockerCompose)
	if err != nil {
		err = fmt.Errorf("ошибка при десериализации из YAML: %s", err.Error())
		fmt.Println(err)
		return err
	}

	jsonData, err := json.Marshal(&y.DockerCompose)
	if err != nil {
		err = fmt.Errorf("ошибка при сериализации в JSON: %s", err.Error())
		fmt.Println(err)
		return err
	}

	err = os.WriteFile(y.FileOutput, jsonData, 0644)
	if err != nil {
		err = fmt.Errorf("ошибка при записи JSON в файл: %s", err.Error())
		fmt.Println(err)
		return err
	}

	return nil
}
