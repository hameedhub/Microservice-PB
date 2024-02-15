package csvlog

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"reflect"
	"time"
)

type Logger interface {
	Log(log Log)
}
type Log struct {
	Priority     string
	Partition    int32
	SentTime     time.Time
	ReceivedTime time.Time
	TimeSpent    int64
	PayloadSize  int
}

type logger struct {
	fileName string
}

func NewLogger(fileName string) (Logger, error) {
	if fileName == "" {
		return nil, errors.New("file name required")
	}
	return &logger{
		fileName: fileName,
	}, nil
}

func (l *logger) Log(log Log) {
	//Time difference in milliseconds
	log.TimeSpent = log.ReceivedTime.Sub(log.SentTime).Milliseconds()

	path := fmt.Sprintf("%s.csv", l.fileName)
	key, values := keyValues(log)
	var csvArray [][]string
	_, err := os.Stat(path)
	if err != nil {
		csvArray = append(csvArray, key)
	}
	csvArray = append(csvArray, values)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(f)

	w := csv.NewWriter(f)
	err = w.WriteAll(csvArray)
	if err != nil {
		fmt.Println(err)
	}
}

func keyValues(log Log) ([]string, []string) {
	s := reflect.ValueOf(log)

	var keys []string
	var values []string
	for i := 0; i < s.NumField(); i++ {
		keys = append(keys, s.Type().Field(i).Name)
		values = append(values, fmt.Sprint(s.Field(i).Interface()))
	}
	return keys, values
}
