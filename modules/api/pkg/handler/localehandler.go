package handler

import (
	"encoding/json"
	"os"

	"github.com/golang/glog"
	"golang.org/x/text/language"
)

const defaultLocaleDir = "en"
const assetsDir = "public"

type Localization struct {
	Translations []string `json:"translations"`
}

func getSupportedLocales(configFile string) ([]language.Tag, error) {
	// read config file
	localesFile, err := os.ReadFile(configFile)
	if err != nil {
		return []language.Tag{}, err
	}

	// unmarshall
	localization := Localization{}
	err = json.Unmarshal(localesFile, &localization)
	if err != nil {
		glog.Warningf("%s %s", string(localesFile), err)
	}

	// filter locale keys
	result := []language.Tag{}
	for _, translation := range localization.Translations {
		result = append(result, language.Make(translation))
	}
	return result, nil
}