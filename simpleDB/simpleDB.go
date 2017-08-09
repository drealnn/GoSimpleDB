package simpleDB

import (
	"path/filepath"
	"os"
	"time"
	"fmt"
	"strings"
	"io/ioutil"
	"encoding/json"
	"strconv"
)

type SimpleDB struct {
	dir string
	typePath string
	finaldir string
}

func NewDB(typeName string) *SimpleDB {
	basedir := filepath.FromSlash("/json/")
	typedir := typeName+filepath.FromSlash("/")
	finaldir := "."+basedir+typedir
	var err error
	if pathNotExists(finaldir) {
		err = os.MkdirAll(finaldir, 0777)
	}

	if (err != nil){
		panic(err)
	}

	return &SimpleDB{
		basedir,
		typedir,
		finaldir}
}

func (db *SimpleDB) Check(file string) string{
	if (pathNotExists(db.finaldir+file)){
		return file
	}
	return db.Check(Uniqid())
}

func (db *SimpleDB) Delete(id string) error {
	path := db.finaldir+id
	file := jsonExtension(path)
	err := os.Remove(file)

	return err
}

func (db *SimpleDB) Get(id string) ([]byte, error) {
	path := jsonExtension(db.finaldir+id)
	if (!pathNotExists(path)){
		if output, err := ioutil.ReadFile(path); err == nil {
			return output, nil
		} else {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("id doesn't exist: "+path)
	}
}

//TODO: make concurrent
func (db *SimpleDB) GetAll() ([][]byte, error) {
	files, err := ioutil.ReadDir(db.finaldir)
	if (err == nil){
		if numFiles := len(files); numFiles > 0 {
			output := make([][]byte, numFiles)
			for i, v := range files {
				jsonString, err := db.Get(v.Name())
				if (err != nil){
					return nil, err
				}
				output[i] = make([]byte, len(jsonString))
				output[i] = jsonString
			}

			return output, nil
		}
		return nil, fmt.Errorf("No tables found")
	}
	
	return nil, err
}

func (db *SimpleDB) Post(data interface{}) (string, error) {
	file := db.Check(Uniqid())
	_, err := db.Put(file, data)
	if err == nil {
		return file, nil
	}

	return "", err
}

func (db *SimpleDB) Put(id string, data interface{}) (interface{}, error) {
	content, err := json.Marshal(data)
	if (err == nil){
		if err := ioutil.WriteFile(db.finaldir+id, content, 0777); err == nil {
			return content, nil
		} else {
			return nil, err
		}
	}

	return nil, err
}

func Uniqid() (string){
	microseconds := time.Now().UnixNano() / 1e3
	seconds := microseconds / 1e6
	secondString := strconv.FormatInt(seconds, 16)
	microsecondsString := strconv.FormatInt(microseconds, 16)
	id := fmt.Sprintf("%v%v",secondString[:8], microsecondsString[len(microsecondsString) - 5:]);
	
	return id
}

func jsonExtension (id string) string {
	if (strings.HasSuffix(id, ".json")) {
		return id
	} else {
		return id + ".json"
	}
}

func pathNotExists(path string) bool {
	_, err := os.Stat(path)
	return (err != nil && os.IsNotExist(err))
}
