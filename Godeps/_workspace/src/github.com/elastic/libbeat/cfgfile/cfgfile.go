package cfgfile

import (
	"flag"
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Command line flags
var configfile *string
var testConfig *bool

func init() {
	// The default config cannot include the beat name as it is not initialised when this function is called
	configfile = flag.String("c", "/etc/beat/beat.yml", "Configuration file")
	testConfig = flag.Bool("test", false, "Test configuration and exit.")
}

// Read reads the configuration from a yaml file into the given interface structure.
// In case path is not set this method reads from the default configuration file for the beat.
func Read(out interface{}, path string) error {

	if path == "" {
		path = *configfile
	}

	filecontent, err := ioutil.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Failed to read %s: %v. Exiting.", path, err)
	}
	if err = yaml.Unmarshal(filecontent, out); err != nil {
		return fmt.Errorf("YAML config parsing failed on %s: %v. Exiting.", path, err)
	}

	return nil
}

func IsTestConfig() bool {
	return *testConfig
}
