package service

import "fmt"

type dummyService struct {
}

type DummyService interface {
	NeverBeenTest(x, y int) int
	NothingToSeeHere(x string) string
}

func NewDummyService() DummyService {
	return &dummyService{}
}

func (ds *dummyService) NeverBeenTest(x, y int) int {
	return x + y
}

func (ds *dummyService) NothingToSeeHere(x string) string {
	return fmt.Sprintf("%s %s", "Halo", x)
}
