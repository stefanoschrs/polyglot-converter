package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"

	"github.com/k0kubun/pp"
	"gopkg.in/yaml.v2"
)

const (
	TypeString   = "string"
	TypeStringPy = "str"
	TypeInt      = "int"
	TypeIntTs    = "number"
)

type ErrUnsupported struct{}

func (e *ErrUnsupported) Error() string {
	return "unsupported type"
}

func processEntryToGo(key string, items map[string]interface{}) (entry string, err error) {
	var entryType string
	for _, value := range items {
		if _, ok := value.(int); ok {
			entryType = TypeInt
		}
		if _, ok := value.(string); ok {
			entryType = TypeString
		}
		break
	}
	if entryType == "" {
		err = new(ErrUnsupported)
		return
	}

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("type %s struct {}\n", key))

	var getKeySb strings.Builder
	getKeySb.WriteString(fmt.Sprintf("func (a %s) GetKey(value %s) string {\n", key, entryType))

	var getValueSb strings.Builder
	getValueSb.WriteString(fmt.Sprintf("func (a %s) GetValue(key string) %s {\n", key, entryType))

	for itemKey, itemValue := range items {
		var itemValueParsed string
		if entryType == TypeString {
			itemValueParsed = fmt.Sprintf(`"%s"`, itemValue.(string))
		} else if entryType == TypeInt {
			itemValueParsed = fmt.Sprint(itemValue.(int))
		}

		sb.WriteString(fmt.Sprintf("func (a %s) %s() %s { return %s }\n", key, itemKey, entryType, itemValueParsed))

		sb.WriteString(fmt.Sprintf("func (a %s) %sKey() string { return \"%s\" }\n", key, itemKey, itemKey))

		getKeySb.WriteString(fmt.Sprintf("\tif value == a.%s() { return a.%sKey() }\n", itemKey, itemKey))

		getValueSb.WriteString(fmt.Sprintf("\tif key == a.%sKey() { return a.%s() }\n", itemKey, itemKey))
	}

	getKeySb.WriteString("\treturn \"\"\n")
	getKeySb.WriteString("}\n")
	sb.WriteString(getKeySb.String())

	if entryType == TypeString {
		getValueSb.WriteString("\treturn \"\"\n")
	} else if entryType == TypeInt {
		getValueSb.WriteString("\treturn -1\n")
	}
	getValueSb.WriteString("}\n")
	sb.WriteString(getValueSb.String())

	return sb.String(), nil
}

func processEntryToTs(key string, items map[string]interface{}) (entry string, err error) {
	var entryType string
	for _, value := range items {
		if _, ok := value.(int); ok {
			entryType = TypeIntTs
		}
		if _, ok := value.(string); ok {
			entryType = TypeString
		}
		break
	}
	if entryType == "" {
		err = new(ErrUnsupported)
		return
	}

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("class %s {\n", key))

	var getKeySb strings.Builder
	getKeySb.WriteString(fmt.Sprintf("\tstatic GetKey (value: %s): string {\n", entryType))

	var getValueSb strings.Builder
	getValueSb.WriteString(fmt.Sprintf("\tstatic GetValue (key: string): %s {\n", entryType))

	for itemKey, itemValue := range items {
		var itemValueParsed string
		if entryType == TypeString {
			itemValueParsed = fmt.Sprintf(`"%s"`, itemValue.(string))
		} else if entryType == TypeIntTs {
			itemValueParsed = fmt.Sprint(itemValue.(int))
		}

		sb.WriteString(fmt.Sprintf("\tstatic %s = %s\n", itemKey, itemValueParsed))

		sb.WriteString(fmt.Sprintf("\tstatic %sKey = \"%s\"\n", itemKey, itemKey))

		getKeySb.WriteString(fmt.Sprintf("\t\tif (value === this.%s) { return this.%sKey }\n", itemKey, itemKey))

		getValueSb.WriteString(fmt.Sprintf("\t\tif (key === this.%sKey) { return this.%s }\n", itemKey, itemKey))
	}

	getKeySb.WriteString("\t\treturn \"\"\n")
	getKeySb.WriteString("\t}\n")
	sb.WriteString(getKeySb.String())

	if entryType == TypeString {
		getValueSb.WriteString("\t\treturn \"\"\n")
	} else if entryType == TypeIntTs {
		getValueSb.WriteString("\t\treturn -1\n")
	}
	getValueSb.WriteString("\t}\n")
	sb.WriteString(getValueSb.String())

	sb.WriteString("}\n")

	return sb.String(), nil
}

func processEntryToPy(key string, items map[string]interface{}) (entry string, err error) {
	var entryType string
	for _, value := range items {
		if _, ok := value.(int); ok {
			entryType = TypeInt
		}
		if _, ok := value.(string); ok {
			entryType = TypeStringPy
		}
		break
	}
	if entryType == "" {
		err = new(ErrUnsupported)
		return
	}

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("class %s:\n", key))

	var getKeySb strings.Builder
	getKeySb.WriteString(fmt.Sprintf("\t@staticmethod\n\tdef GetKey(value: %s) -> str:\n", entryType))

	var getValueSb strings.Builder
	getValueSb.WriteString(fmt.Sprintf("\t@staticmethod\n\tdef GetValue(key: str) -> %s:\n", entryType))

	for itemKey, itemValue := range items {
		var itemValueParsed string
		if entryType == TypeStringPy {
			itemValueParsed = fmt.Sprintf(`"%s"`, itemValue.(string))
		} else if entryType == TypeInt {
			itemValueParsed = fmt.Sprint(itemValue.(int))
		}

		sb.WriteString(fmt.Sprintf("\t%s = %s\n", itemKey, itemValueParsed))

		sb.WriteString(fmt.Sprintf("\t%sKey = \"%s\"\n", itemKey, itemKey))

		getKeySb.WriteString(fmt.Sprintf("\t\tif value == %s.%s: return %s.%sKey\n", key, itemKey, key, itemKey))

		getValueSb.WriteString(fmt.Sprintf("\t\tif key == %s.%sKey: return %s.%s\n", key, itemKey, key, itemKey))
	}

	getKeySb.WriteString("\t\treturn \"\"\n")
	sb.WriteString(getKeySb.String())

	if entryType == TypeStringPy {
		getValueSb.WriteString("\t\treturn \"\"\n")
	} else if entryType == TypeInt {
		getValueSb.WriteString("\t\treturn -1\n")
	}
	sb.WriteString(getValueSb.String())

	return sb.String(), nil
}

func main() {
	inputFilename := os.Args[1]
	file, err := ioutil.ReadFile(inputFilename)
	if err != nil {
		log.Fatal(fmt.Errorf("ioutil.ReadFile: %w", err))
	}

	var data map[string]map[string]interface{}
	err = yaml.Unmarshal(file, &data)
	if err != nil {
		log.Fatal(fmt.Errorf("yaml.Unmarshal: %w", err))
	}

	if os.Getenv("DEBUG") == "true" {
		pp.Println(data)
	}

	var entryGo strings.Builder
	var entryTs strings.Builder
	var entryPy strings.Builder

	for key, items := range data {
		var entry string
		entry, err = processEntryToGo(key, items)
		if err != nil {
			log.Fatal(fmt.Errorf("processEntryToGo: %w", err))
		}
		entryGo.WriteString(entry + "\n")

		entry, err = processEntryToTs(key, items)
		if err != nil {
			log.Fatal(fmt.Errorf("processEntryToTs: %w", err))
		}
		entryTs.WriteString(entry + "\n")

		entry, err = processEntryToPy(key, items)
		if err != nil {
			log.Fatal(fmt.Errorf("processEntryToPy: %w", err))
		}
		entryPy.WriteString(entry + "\n")
	}

	getOutName := func(s string) string {
		return strings.Replace(
			path.Base(inputFilename),
			path.Ext(inputFilename),
			"."+s,
			1)
	}

	// TODO: change the package with a flag
	err = ioutil.WriteFile(path.Join(os.Args[2], getOutName("go")), []byte("package main\n"+entryGo.String()), 0644)
	if err != nil {
		log.Fatal(fmt.Errorf("ioutil.WriteFile(go): %w", err))
	}

	err = ioutil.WriteFile(path.Join(os.Args[2], getOutName("ts")), []byte(entryTs.String()), 0644)
	if err != nil {
		log.Fatal(fmt.Errorf("ioutil.WriteFile(ts): %w", err))
	}

	err = ioutil.WriteFile(path.Join(os.Args[2], getOutName("py")), []byte(entryPy.String()), 0644)
	if err != nil {
		log.Fatal(fmt.Errorf("ioutil.WriteFile(py): %w", err))
	}
}
