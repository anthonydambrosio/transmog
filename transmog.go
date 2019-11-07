package transmog

// TODO: Not thread safe.

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"

	"github.com/clbanning/mxj"
	"github.com/ghodss/yaml"
)

// Transmog an instance of data that can be transmogrified.
type Transmog struct {
	data interface{}
}

func (c *Transmog) parse(data []byte) error {
	err := yaml.Unmarshal(data, &c.data)
	if err != nil {
		return err
	}
	return nil
}

func (c *Transmog) parseXML(data []byte) error {
	xmldata, err := mxj.NewMapXml(data)
	if err != nil {
		return err
	}
	c.data = xmldata
	return nil
}

func traverse(data interface{}, path []string, value *string, write bool) error {
	// No path, don't know what property we are working with.
	if len(path) < 1 {
		return errors.New("path is zero length")
	}
	// We have at least one element in the path, get a reference to the data.
	var node map[string]interface{}
	node = data.(map[string]interface{})
	// Only one element in the path, get or set the value of the node.
	if len(path) == 1 {
		switch node[path[0]].(type) {
		case string:
			if write {
				node[path[0]] = *value
			} else {
				*value = node[path[0]].(string)
			}
			return nil
		case float64:
			if write {
				f, err := strconv.ParseFloat(*value, 64)
				if err != nil {
					return fmt.Errorf("Value type is not a number (%s): %v", *value, err.Error())
				}
				node[path[0]] = f
			} else {
				*value = strconv.FormatFloat(node[path[0]].(float64), 'f', -1, 64)
			}
			return nil
		case bool:
			if write {
				b, err := strconv.ParseBool(*value)
				if err != nil {
					return fmt.Errorf("Value type is not a bool (%s): %v", *value, err.Error())
				}
				node[path[0]] = b
			} else {
				*value = strconv.FormatBool(node[path[0]].(bool))
			}
			return nil
		default:
			valueType := strings.Replace(reflect.ValueOf(node[path[0]]).Kind().String(), "map", "object", -1)
			valueType = strings.Replace(valueType, "slice", "array", -1)
			return fmt.Errorf("value is %s", valueType)
		}
	} else {
		// Traverse
		switch node[path[0]].(type) {
		case map[string]interface{}:
			return traverse(node[path[0]], path[1:], value, write)
		case []interface{}:
			if len(path) < 2 {
				return errors.New("Path too long")
			}
			// The next path value must be an index into this array.
			index, err := strconv.Atoi(path[1])
			if err != nil {
				return fmt.Errorf("path element '%v' is not an index: %v", path[1], err.Error())
			}
			tarray := node[path[0]].([]interface{})
			if index > len(tarray)-1 {
				return fmt.Errorf("index %d out of range [%d]", index, len(tarray))
			}
			switch tarray[index].(type) {
			case string:
				if write {
					tarray[index] = *value
				} else {
					*value = tarray[index].(string)
				}
				return nil
			case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
				if write {
					v, err := strconv.Atoi(*value)
					if err != nil {
						return err
					}
					tarray[index] = v
				} else {
					*value = strconv.Itoa(tarray[index].(int))
				}
				return nil
			case map[string]interface{}:
				return traverse(tarray[index], path[2:], value, write)
			default:
				return nil // Should be some sort of error.
			}
		case nil:
			return fmt.Errorf("nil:unknown value type %v", reflect.ValueOf(node[path[0]]).Kind())
		default:
			return fmt.Errorf("unexpected error traversing")
		}
	}
}

// Load a JSON or YAML file.
func (c *Transmog) Load(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	if filepath.Ext(path) == ".xml" {
		return c.parseXML(data)
	}
	return c.parse(data)
}

// ToYaml transmogrify data to YAML.
func (c *Transmog) ToYaml() ([]byte, error) {
	return yaml.Marshal(c.data)
}

// ToJSON transmogrify data to JSON.
func (c *Transmog) ToJSON() ([]byte, error) {
	return json.MarshalIndent(c.data, "", "  ")
}

// ToXML transmogrify data to XML.
func (t *Transmog) ToXML(indent string) ([]byte, error) {
	j, err := json.Marshal(t.data)
	if err != nil {
		return []byte{}, err
	}
	mv, err := mxj.NewMapJson(j)
	if err != nil {
		return []byte{}, err
	}
	xml, err := mv.XmlIndent("", indent)
	if err != nil {
		return []byte{}, err
	}
	return xml, nil
}

// Get value from data give path string.
func (c *Transmog) Get(path []string) (string, error) {
	value := ""
	err := traverse(c.data, path, &value, false)
	return value, err
}

// Set value in data found at path string.
func (c *Transmog) Set(path []string, value string) error {
	err := traverse(c.data, path, &value, true)
	return err
}
