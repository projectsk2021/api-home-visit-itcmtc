package services

import (
	"fmt"

	"github.com/kamchai-n/api-student-home-visit/domain"
	repoStatic "github.com/kamchai-n/api-student-home-visit/repository/static"
)

type staticService struct {
	staticRepo repoStatic.StaticRepository
}

func NewStaticService(staticRepo repoStatic.StaticRepository) StaticService {
	return staticService{staticRepo: staticRepo}
}
func (s staticService) ListEducation() (education domain.ResponseGetEducation, err error) {

	degree, err := s.staticRepo.GetDegress()
	if err != nil {
		return education, fmt.Errorf("1325")
	}

	faculty, err := s.staticRepo.GetFaculty()
	if err != nil {
		return education, fmt.Errorf("1326")
	}

	majors, err := s.staticRepo.GetMajors()
	if err != nil {
		return education, fmt.Errorf("1327")
	}

	education = domain.ResponseGetEducation{
		Faculty: faculty,
		Degree:  degree,
		Major:   majors,
	}

	return education, nil
}
