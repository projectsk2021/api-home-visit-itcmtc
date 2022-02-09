package services

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"

	"github.com/google/uuid"
	"github.com/kamchai-n/api-student-home-visit/domain"
	"github.com/kamchai-n/api-student-home-visit/middlewares"
	"github.com/kamchai-n/api-student-home-visit/model"
	repoSetting "github.com/kamchai-n/api-student-home-visit/repository/setting"
	repoStudent "github.com/kamchai-n/api-student-home-visit/repository/student"
	"github.com/xuri/excelize/v2"
)

type studentService struct {
	stdRepo     repoStudent.StudentRepository
	settingRepo repoSetting.SettingRepository
}

func NewStudentService(stdRepo repoStudent.StudentRepository, settingRepo repoSetting.SettingRepository) StudentService {
	return studentService{stdRepo: stdRepo, settingRepo: settingRepo}
}

func (s studentService) NewStudentFromExcel(file *multipart.FileHeader, createdBy uuid.UUID) ([]domain.ResponseCreateStudentFromExcel, error) {
	locationFile := fmt.Sprintf("./static/temp/%s", file.Filename)
	worksheet := "example_form"

	// Move File To Temp
	resFile, err := os.Create(locationFile)
	if err != nil {
		return nil, fmt.Errorf("1303")
	}
	defer resFile.Close()

	multipartFile, err := file.Open()
	if err != nil {
		return nil, fmt.Errorf("1304")
	}

	io.Copy(resFile, multipartFile)
	defer os.Remove(locationFile)

	// Read Data Form File
	f, err := excelize.OpenFile(locationFile)
	if err != nil {
		return nil, fmt.Errorf("1305")

	}

	rows, err := f.GetRows(worksheet)
	if err != nil {
		return nil, fmt.Errorf("1306")
	}

	var dataForm []domain.StudentForm
	for _, row := range rows[4:] {

		var student_level uint

		if row[4] == "ป.ตรี" || row[4] == "ตรี" || row[4] == "ปริญญาตรี" || row[4] == "49" {
			student_level = 49
		} else if row[4] == "ปวส." || row[4] == "ปวส" || row[4] == "ประกาศนียบัตรวิชาชีพชั้นสูง" || row[4] == "39" {
			student_level = 39
		} else if row[4] == "ปวช." || row[4] == "ปวช" || row[4] == "ประกาศนียบัตรวิชาชีพ" || row[4] == "29" {
			student_level = 29
		} else {
			student_level = 0
		}

		dataForm = append(dataForm, domain.StudentForm{
			StudentId:          row[1],
			StudentName:        row[2],
			StudentPhone:       &row[3],
			StudentLevel:       &student_level,
			StudentImage:       &row[5],
			AddressNo:          &row[6],
			AddressLane:        &row[7],
			AddressRoad:        &row[8],
			AddressSubDistrict: row[9],
			AddressDistrict:    row[10],
			AddressProvince:    row[11],
			Latitude:           row[12],
			Longitude:          row[13],
			ParentName:         row[14],
			ParentPhone:        row[15],
			CreatedBy:          createdBy,
		})
	}

	// Send to database
	response, err := s.stdRepo.CreateStudentFromExcel(dataForm)
	if err != nil {
		return nil, fmt.Errorf("1302")
	}
	return response, err
}

func (s studentService) NewStudent(req domain.StudentForm) (domain.ResponseCreateStudentFromExcel, error) {
	students, err := s.stdRepo.CreateStudent(req)
	if err != nil {
		return students, fmt.Errorf("1302")
	}

	return students, nil
}

func (s studentService) ListInfoStudent(req *domain.RequestGetAll) ([]model.StudentCustom, error) {
	students, err := s.stdRepo.GetInfoStudent(req, *middlewares.UserClaims.UserId)
	if err != nil {
		return students, fmt.Errorf("1307")
	}

	return students, nil
}

func (s studentService) ListInfoStudentById(req domain.RequestGetOneStudent) (model.StudentCustom, error) {
	students, err := s.stdRepo.GetInfoStudentById(req)
	if err != nil {
		return students, fmt.Errorf("1308")
	}
	return students, nil
}

func (s studentService) RemoveStudentById(req []string) error {
	if err := s.stdRepo.DeleteStudentById(req); err != nil {
		return fmt.Errorf("1318")
	}
	return nil
}
