package errcodes
var (
	codeMessages = [...]string{
		0: "Внутренняя ошибка",
		1: "Неверный хешкод конвертации",
		2: "Невалидный запрос",
		3: "Размер файла больше 100мб",
		4: "Пустой файл",
		5: "Пустая презентация",
		6: "Не получилось спарсить pdf",
		7: "Количество слайдов больше 100",
		8: "Ошибка при конвертации",
		9: "Ошибка создания презентации",
		10: "Презентация не найдена",
		11: "Версия презентации не найдена",
		12: "Пользователь не найден",
		13: "Интерактив не найден",
		14: "Слайды не найдены",
		15: "Не удалось запустить сессию - пустая презентация",
		16: "Не удалось получить (или отсутсвует) refresh token",
		17: "...",
		18: "...",
		19: "...",
		20: "...",
		21: "...",
		22: "...",
		23: "...",
		24: "...",
		25: "Не получилось создать интерактив",
		26: "Ответы не найдены",

	}
	codesHTTP = [...]int{
		0: 500,
		1: 400,
		2: 400,
		3: 400,
		4: 400,
		5: 400,
		6: 400,
		7: 400,
		8: 400,
		9: 500,
		10: 400,
		11: 400,
		12: 400,
		13: 400,
		14: 400,
		15: 400,
		16: 400,
		17: 400,
		18: 400,
		19: 400,
		20: 400,
		21: 400,
		22: 400,
		23: 400,
		24: 400,
		25: 500,
		26: 400,

	}
)

const (
	startCode = 10000

	ErrUnknown CodeError = iota - 2
	ErrNil
)

const (
	ErrInternal CodeError = iota + 10000
	ErrInvalidHashcode
	ErrInvalidRequest
	ErrFileToBig
	ErrEmptyFile
	ErrEmptyPresentation
	ErrParsePDF
	ErrSlidesLimit
	ErrConvertation
	ErrPresentationCreate
	ErrPresentationNotFound
	ErrPresentationVersionNotFound
	ErrUserNotFound
	ErrInteractiveNotFound
	ErrSlidesNotFound
	ErrSessionEmptyPresentation
	ErrNoRefreshToken
	ErrNoAuthorizationHeader
	ErrInvalidAuthorizationHeader
	ErrTokenEmpty
	ErrInvalidToken
	ErrTokenExpired
	ErrTokensCountExceeded
	ErrEmptyVKIDStatePassed
	ErrInvalidPriveleges
	ErrInteractiveCreate
	ErrAnswersNotFound
)

type CodeError int

func (e CodeError) Error() string {
	code := e - startCode
	if code < 0 || int(code) >= len(codeMessages) {
		return "Неизвестная ошибка"
	}

	return codeMessages[code]
}

func GetHTTPCode(e CodeError) int {
	code := e - startCode
	if code < 0 || int(code) >= len(codesHTTP) {
		return 500
	}

	return codesHTTP[code]
}
	