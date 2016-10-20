package envconfigfromfile

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/kelseyhightower/envconfig"
)

type TestStruct struct {
	TestField *EnvConfigFromFile `envconfig:"TEST_FIELD_FILE_PATH"`
}

type TestStructPathNotSet struct {
	TestField *EnvConfigFromFile `envconfig:"TEST_FIELD_PATH_NOT_SET"`
}

type TestStructPathNotSetButRequired struct {
	TestField *EnvConfigFromFile `envconfig:"TEST_FIELD_PATH_NOT_SET_BUT_REQUIRED" required:"true"`
}

type TestStructPathDoesNotExist struct {
	TestField *EnvConfigFromFile `envconfig:"TEST_FIELD_PATH_DOES_NOT_EXIST"`
}

func TestEnvConfigFromFile(t *testing.T) {
	testFilePath := "testdata/test_field_contents.txt"
	testFileContents, _ := ioutil.ReadFile(testFilePath)
	os.Clearenv()
	os.Setenv("TEST_FIELD_FILE_PATH", testFilePath)
	testStruct := TestStruct{}
	err := envconfig.Process("", &testStruct)
	if err != nil {
		t.Error(err)
	}
	if testStruct.TestField.Value != string(testFileContents) {
		t.Error(err)
	}
}

func TestEnvConfigFromFilePathNotSet(t *testing.T) {
	os.Clearenv()
	testStruct := TestStructPathNotSet{}
	err := envconfig.Process("", &testStruct)
	if err != nil {
		t.Errorf("envconfigfromfile: this test should not return an error. Got: %s", err.Error())
	}
}

func TestEnvConfigFromFilePathNotSetButRequired(t *testing.T) {
	os.Clearenv()
	testStruct := TestStructPathNotSetButRequired{}
	err := envconfig.Process("", &testStruct)
	if err == nil {
		t.Error("envconfigfromfile: this test should return an error")
	}
	expectedError := "required key TEST_FIELD_PATH_NOT_SET_BUT_REQUIRED missing value"
	if err.Error() != expectedError {
		t.Errorf("envconfigfromfile: the error returned does not match the expected value. Wanted: '%s' Got: '%s'", expectedError, err.Error())
	}
}

func TestEnvConfigFromFilePathDoesNotExist(t *testing.T) {
	testFilePath := "testdata/non_existent_file.txt"
	os.Clearenv()
	os.Setenv("TEST_FIELD_PATH_DOES_NOT_EXIST", testFilePath)
	testStruct := TestStructPathDoesNotExist{}
	err := envconfig.Process("", &testStruct)
	if err == nil {
		t.Error("envconfigfromfile: this test should return an error")
	}
	expectedError := "envconfig.Process: assigning TEST_FIELD_PATH_DOES_NOT_EXIST to TEST_FIELD_PATH_DOES_NOT_EXIST: converting 'testdata/non_existent_file.txt' to type envconfigfromfile.EnvConfigFromFile. details: envconfigfromfile: open testdata/non_existent_file.txt: no such file or directory"
	if err.Error() != expectedError {
		t.Errorf("envconfigfromfile: the error returned does not match the expected value. Wanted: '%s' Got: '%s'", expectedError, err.Error())
	}
}
