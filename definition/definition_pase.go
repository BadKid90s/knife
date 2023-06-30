package definition

import (
	"errors"
	"gateway/util"
)

func ParseConfig(configFile string) error {

	buffer, err := util.ReadConfigFile(configFile)
	if err != nil {
		return err
	}

	err = ParseServerConfig(buffer)
	if err != nil {
		return errors.New("an error occurred while parsing the server configuration")
	}

	err = ParseRouteConfig(buffer)
	if err != nil {
		return errors.New("an error occurred while parsing the routes configuration")
	}

	return nil

}
