package config

import (
	"bytes"
	"reflect"
	"testing"
)

func TestJson(t *testing.T) {
	//Testing if FromJson(…ToJson()) works:
	config1 := EmptyConfig()
	config2, err := FromJson(config1.ToJson())
	if err != nil {
		t.Errorf("TestJson case 1.1 had an error:\n%s\n", err)
	}
	if !reflect.DeepEqual(config1, config2) {
		t.Errorf("TestJson case 1.2 had an error:\n\t%v\n\t%v\n", config1, config2)
	}
	//Testing if ToJson(FromJson(…)) works:
	data1 := []byte(`{
  "HostCertMap": {
    "das.ohren.gift": ""
  },
  "HostEndpointMap": {
    "das.ohren.gift": {
      "Addr": "[::1]",
      "Port": "27374"
    }
  }
}`)
	config1, err = FromJson(data1)
	if err != nil {
		t.Errorf("TestJson case 2.1 had an error:\n%s\n", err)
	}
	data2 := config1.ToJson()
	if !bytes.Equal(data1, data2) {
		t.Errorf("TestJson case 2.2 had an error:\n\t%v\n\t%v\n", data1, data2)
	}
}
