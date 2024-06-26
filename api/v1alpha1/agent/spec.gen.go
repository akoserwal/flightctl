// Package v1alpha1 provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.3.0 DO NOT EDIT.
package v1alpha1

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	externalRef0 "github.com/flightctl/flightctl/api/v1alpha1"
	"github.com/getkin/kin-openapi/openapi3"
)

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/+w763PbtpP/CobtTB+jh5PLdVp9c5yk9eRhj+T0Zi72dSByJaEhAQYA5aoZ/e83eJEg",
	"AVqUY/f3If2SyFhgd7HYN8DPScqKklGgUiSzz4lIN1Bg/fO0LHOSYkkYXUgsKz1YclYClwT0XxQXoP7P",
	"QKSclGpqMkt+qwpMEQec4WUOSE1CbIXkBhBucE6SUSJ3JSSzREhO6DrZjxK1aBdivNoAolWxBK4QpYxK",
	"TChwgW43JN0gzEGT2yFCB5IREnOz4zaldzUVNwexpQC+hQytGL8DO6ES1sAVelGL61sOq2SWfDNtpDy1",
	"Ip4G8r1SiPaavU8V4ZAlsw9GxE4wHuc1lZuaA7b8E1KpGIijnn1OgFaFwnrJocRaGqNkoRCan/OKUvPr",
	"JeeMJ6PkPf1I2S1NRskZK8ocJGQeRSvRUfLXWGEebzFX/ApFIuDBpxkAPSYCWMNVAHJsBoCG7wDkbaQt",
	"KrGoigLzXZ+2E7piB7VdTeKFxocykJjkhK612uRYSCR2QkLhqxCSHFNBenX1aGVqbyOqVMNUJ4LIU6Hf",
	"AOdyo3TyBaw5ziCLqM3RqtKm2dDoneIR750T0ZL2hJrdvdJzmhFztt2jrkHOBQl7xEL7BkYBYVFCKp27",
	"SyvOgUqk5G19IBHo9PIczUGwiqegjrytZUpNrmqVuCIxD3vl1EmSAgylmrVGnZTL4qzQfJkTR5IhTJnc",
	"AFeEjaYmsyTDEsYKV0wBCxACrw/7eTsPEZppIdN1LR28ZJW0HN+t7c7Z/goUOI4fg9r9pACJMyzxZF3P",
	"RHKDZUcat1ggARItsYAMVaUhW2+cUPnTs6gP54BFjPj3S05g9QMy8Dom1BS/E4P2Ocyqa4WzLmnvMA1c",
	"FjV+jaHmYBRTuHr7zenHfEWXPc87XPFKoXmFcwFH+4MOXourM+pQd4ajptz1XadlydnWOI00BSHIMofu",
	"H85ELzEXeupiR1P942ILPMdlSeh6ATmkknElyN9xThT4fZlhG8sWJaRu2Pw/TAIvKWd5XgCVc/hUgZAe",
	"x3MomSCS8V2UXcVlLyDYkw+s9/cqB5A9m9Qwt6UXsCUpePs1A/6ur6Aocyzhd+CCMGqFsHdTQwMz44hD",
	"yUEotUYYlZudICnOUaaBodPEJbEEQoSnl+cWhjJYEQpCW+zWjEGGjNnU7rmmbJwKWyFMkVH6CVoo78QF",
	"EhtW5Zky+y1wiTikbE3J3zU27WpN1JcgJFKehVOcoy3OKxghTDNU4B3ioPCiinoY9BQxQW8ZN/nEDG2k",
	"LMVsOl0TOfn4s5gQpuy+qCiRu6kKRpwsK3VC0wy2kE8FWY8xTzdEQiorDlNckrFmluroNymyb7g9ehHz",
	"Tx8JzUJRviY0Q0SdiJlpWG0k5lKd+cvFFXL4jVSNAL1jbWSp5EDoCriZqWOWwgI0Kxmh1qXnREfSalkQ",
	"qQ5Jm4US8wSdYUqZREtAlVJFyCbonKIzXEB+hgU8uiSV9MRYiUzEA6gJVYfc9oUW0VuQWEeIEtJDKxpz",
	"Gx5T7BobUDqxwbMjqwMe+7EQYLC1Eque7NlJAGfGJ+P8sgU/qlRSpNuq+RaXylQj+bURCwgvDjf8C5MG",
	"3ju9DiSot9ng7ZfZGaMrsu6TFgeaAYes16s5l2Yzzcx5TbNMOaYVWUdSjw67XTr9/J6r3IgT2VseDRRl",
	"FJuVaVioHBRjD6IvL95MzloXbsTReZjM7i7mjy3ZDuLyC38sTOB/hUmufzSV8nsqqrJkfHiNH6Vck4hC",
	"a7pRaMNMD9jjsN75xWJhfWXnyIto2cKE5ABIQ21ziqP38zeHjcUg7D+Ci0Vv6yDOSseILxaGqy/npE7r",
	"evhJy2qYhrYRGc0cJRkRH79kfQEFG+opYhg60lC7qZFa7obKpr+t8T+Y27bTGSdSpZ9HFzIxgn7fJIQ2",
	"RGNQj5EY2K99vNQgPH0dHEJ1fEOEtA3WFVnXCZyOnERCYeI4UUsKQrFk3MO9e6ebwRa5O2lG4WKVzD7c",
	"fdK/EmnC4SVnW5IBtynN3ateV0vgFCSIBaQc5FGLz2lOKMSoxjTHDmDO8U793bSgQ+kWWKabSyxVwm+M",
	"34muNIPJLPm/D3j8943652T8y/iPyc2P38ZiSZvsPsIYGxhprI9U8UlnSNkxfBf4rzdA13KTzJ7+90+j",
	"7j5Ox/97Mv5ldn09/mNyfX19/eM9d7PvNdnGjcVqRQP1K8Z4HLc9MFXQuUIS2bWqBJEck9w0+FNZ4bzp",
	"2uE76s4mLxx2EJFU2eiTyYrFfVPkpm0ZpMY1yJOR3qdpthlezD6jTUt/+8EBNb7k8N5bKe9+lNQ51b2y",
	"pSPVv17T1L33iD9H1AxWOdvVgrO/c5uODkDQzN+PElvgDlv63kxuaNvVp1KtH9L97QbaRk1bGxm1DcGX",
	"sX/Ktbbog2s204jUZ7E/gv8D9zS2nnJt84dL+b/ocqYPhZe/XOi4Fr+VmcOSMdupu2S3qu67WK3umdW0",
	"uPCoBjCPkQjU8RYB+exGwK0dROCRbKhletFQUs+wbS7QCRHJxLSqSKa7ehUlnyrId4hkQCVZ7bwCPBIh",
	"vN5R/Crj1JuhHLQub9CyizbQOiWc8xchzueMSXT+4hhUimHdwjX7j/N54SYhM2s4gW6TyRdJvY+Qi34L",
	"aDu2B6/1rfEbV/SQxt/i+37GH6LwjP99ecVeYKmkelHJi5X97TXo72PpLZIeiQjUpxpd3LkpaEM9gw0u",
	"QMLzDKa07wxsh1jf0GF9e4JzZb6gl92Z0/17l/DvXcJXd5cQmNNx1wrh8nvcMFhOY96v50YU56H7x+6u",
	"NNA5B3FvFECg2w3IDZhLfOcyNligJQBFbr7n+JeM5YB1leWgp7Kf0qlUOq6Q66caWNonaz65WyxalIY9",
	"y3Arnu/6qT/fOeqdR3gKyqPhLMdLyO+sQoMlbdoGQSt9skOS6XuZnXNnd1SVHNZR/2vG3abcX9QTqkvc",
	"jU9dgvX3rTPsy0uc3gzSP9daOxCU1DQjDG+iqaqDud8JJDFfg629w9CUCh6STAU3BC5fvh0DTVkGGbp8",
	"fbb45skJStXildZ1JMiaKq9q9S56/Fmn33Lvm0TF6jA59vR2eiYe1+YJkERbOLUfOcrB1Q5oP0o8MUcO",
	"yDuD4KDUoUDmn1P0XNoNok5/J3xlBvGd1x29wZ2koGfX0w6IHrUu7cI+Y997Mj3fPSM7aK31w6T9KNGv",
	"xUhq+0vOMI9qVMc65C69Dk7DvXs+2C1xSOySmJTivW+lmHk+pHUebH0/6kp8TeRcYeiOl1huovvj9ROj",
	"nrfYXjRp5npZGkOVAISNmxY7miIDuabRxrB2enPYEpd1H7q0rtkLFo/Mrm5i3WUfh5VJOE/pU+/lwIMe",
	"CzFNk95QK3kFh7ZhccS3cecFyYNuRWj8USUrWEXlZZ+m9ViSAYgSpwPszD6Rb1aMPKIHFaFhPS5EL5MO",
	"bKGBKcfu0ljzDhXnOSpVpiskZM1dGioqnWFuYWSzFkLTvMpA6BWGslDmw+3clEM0GdCFCj54aXAgXWsm",
	"m3eyppUjgPd3/BU/9qWokLgohzaTFekc7rl0fceD4FMkVIyjaf2VSKuKxEhVFSq6Iu+xcP1wRihNsBkX",
	"umRlpcrsuqllutwTNAecjRnNdwPfD39xHu0eNJni+CPshC70TUFvVCzFVBeqAjLlchlfY1X163kqk1gz",
	"rv78XqSsNKNCv+T8walZ9HzjH/X4Lt/OjTUSbynw2AF5BTyWiN1S4RokZnyksvhrXRBOFanrBBkh932/",
	"o1f192koYiX+VIGTnyZrG7bEdm30S3P+nfAaKs17jKZPE81EAqnN7UOqYffusfzuHhfK/8iFceQpWtgE",
	"Pf5S+T7Xw8e9XNvrK0bTCs5JClRA89FaclridAPo6eQkGSUVz5NZ4vo5t7e3E6zBE8bXU7tWTN+cn718",
	"t3g5fjo5mWxkoTN/SWSu0F2UQJF9u/wWU7wGlXvrbz3GCK/Vb9eDSkbJ1skyqahpJGa2D09xSZJZ8l+T",
	"k8kTm8xocU1xSabbJ1NToInpZ7WN/dRJQad5EGlEqGpSKfuqyvPaDxoP6L5dsKOQ2eK5acUzep4ls+RX",
	"FR8D9VbMcVyA1Er7IfiSznMYNV6iIDpBc36m+cDNHa1Je4w2RoN/79cwun2MuophqX6qQOeLlqyeOw+m",
	"9pO90W6nZEoTFPzpyYm1aQlUdh4DTP+0H480+O4yuYh0tfZ2so3XSkeenjyLfLbIkGNkP0qenTx5MNZM",
	"ERfh5j3FldzoCJMZos8en+g7Jl+xilqCvzw+QffJIF3lxL1RwWt9OWKV+kaN9Vhn0zotKxnrZpU5Tv2O",
	"RdscX8TNcW6WtdovB4zRj94vHtIYb8xkEPI5M1/uPsh5WB73baevmNk/ohn6VGOm9+wBafVq3HOcIXfn",
	"9ZXY8gGjgrqV5W5qtEUxETUp01Nv1phGYI8pnenaKrxsfBytDukMUvAnj81Ap52oZZKZWPPzP0v7NDff",
	"8c/t04KvzOr+swEtsLNDZmjDXG/uqc6yE9IaLYiENZzFLPHOwKZTW0LXwEtOqOztfj9kuHuk6DPIQFwg",
	"+qqCQlQxVdVpLv21WpgKbprsb/b/HwAA//+BFwWh1EQAAA==",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	pathPrefix := path.Dir(pathToFile)

	for rawPath, rawFunc := range externalRef0.PathToRawSpec(path.Join(pathPrefix, "../openapi.yaml")) {
		if _, ok := res[rawPath]; ok {
			// it is not possible to compare functions in golang, so always overwrite the old value
		}
		res[rawPath] = rawFunc
	}
	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
