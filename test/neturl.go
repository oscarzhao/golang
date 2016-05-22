package main

import (
	"fmt"
	"net/url"
)

func main() {

	registryUrls := []string{"gcr.io", "http://gcr.io/", "https://gcr.io", "v1/repositories/google-containers/hyperkube/tags"}
	for _, registryUrl := range registryUrls {
		baseUrl, err := url.Parse(registryUrl)
		if err != nil {
			fmt.Printf("error parse %s: %s\n", registryUrl, err)
		} else {
			fmt.Printf("parse %s success, result: %s\n", registryUrl, baseUrl)
		}
		relnum1, _ := url.Parse("/helpme")
		relnum2, _ := url.Parse("helpme")
		fmt.Printf("resolve /helpme: %v\n", baseUrl.ResolveReference(relnum1))
		fmt.Printf("resolve helpme: %v\n", baseUrl.ResolveReference(relnum2))

	}
}
