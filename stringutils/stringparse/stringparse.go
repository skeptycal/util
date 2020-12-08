package stringparse

import (
	"fmt"
	"os"
)

func main() {
	// Execution of CLI tool behavior goes here
	configuration := Configuration{}
	configFileName := ConfigFileName()

	// mock env variable
	os.Setenv("STRING_PARSE_VERSION", "0.2.2")
	os.Setenv("STRING_PARSE_ENV", "dev-environment-variable")

	err := Configure(configFileName, &configuration)
	if err != nil {
		panic(fmt.Errorf("unable to load configuration file %s: %v", configFileName, err))
	}

	fmt.Println(configuration)

	fmt.Printf("usage text is %v", configuration.Help)

	// helpptr := flag.String("example", "defaultValue", " Help text.")

	// println(*helpptr)
}
