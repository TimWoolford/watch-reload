package main

import "fmt"

type volumeDirsFlag []string

func (v *volumeDirsFlag) Set(value string) error {
	*v = append(*v, value)
	return nil
}

func (v *volumeDirsFlag) String() string {
	return fmt.Sprint(*v)
}