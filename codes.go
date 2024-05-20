package errcodes

import (
	"net/http"
)

var codeMessages = [...]string{
	0: "Внутренняя ошибка",
	1: "Невалидный запрос",
	2: "Невалидный токен",
	3: "Размер файла больше 100мб",
	4: "Пустой файл",
	5: "Пустая презентация",
	6: "Не получилось спарсить pdf",
	7: "Количество слайдов больше 100",
	8: "Ошибка при конвертации",
	9: "Ошибка создания презентации",
}

type CodeError int

func (e CodeError) Error() string {
	if e-startCode < 0 {
		return "Неизвестная ошибка"
	}

	return codeMessages[e-startCode]
}

func GetHTTPCode(errCode CodeError) int {
	switch {
	case errCode >= UnknownEntry && errCode <= ConvertationProblem:
		return http.StatusBadRequest
	case errCode == ErrInternal || errCode == ErrPresentationCreate:
		return http.StatusInternalServerError
	default:
		return http.StatusInternalServerError
	}
}

const (
	startCode = 10000

	ErrUnknown CodeError = iota - 2
	ErrNil
)

const (
	ErrInternal CodeError = iota + startCode
	UnknownEntry
	BadRequest
	FileTooBig
	EmptyFile
	EmptyPresentation
	FailedParsePDF
	SlidesLimit
	ConvertationProblem
	ErrPresentationCreate
)
