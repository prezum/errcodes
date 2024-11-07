package errcodes

import (
	"net/http"
)

var codeMessages = [...]string{ //nolint:gochecknoglobals
	0:  "Внутренняя ошибка",
	1:  "Невалидный запрос",
	2:  "Невалидный токен",
	3:  "Размер файла больше 100мб",
	4:  "Пустой файл",
	5:  "Пустая презентация",
	6:  "Не получилось спарсить pdf",
	7:  "Количество слайдов больше 100",
	8:  "Ошибка при конвертации",
	9:  "Ошибка создания презентации",
	10: "Презентация не найдена",
	11: "Версия презентации не найдена",
	12: "Пользователь не найден",
	13: "Интерактив не найден",
	14: "Слайды не найдены",
	15: "Не удалось запустить сессию - пустая презентация",
	16: "Не удалось получить (или отсутсвует) refresh token",
	17: "Пустая авторизация",
	18: "Неавалидная авторизация",
	19: "Пустой токен авторизации",
	20: "Невалидный токен авторизации",
	21: "Токен авторизации истек",
	22: "Количество сессий превышает лимит",
	23: "Пустое состояние авторизации",
	24: "Недостаточно прав",
	25: "Не получилось создать интерактив",
	26: "Ответы не найдены",
	27: "Не найдена сессия",
	28: "Невалидный id слайда",
	29: "Не удалось уведомить",
	30: "Плохой запрос",
	31: "Плохой запрос - невалидный массив ответов",
	32: "Не найден статус интерактива",
	33: "Пользователь уже ответил",
	34: "Не удалось создать сессию",
	35: "Пользователь не реальный",
	36: "Превышено максимальное количество человек в сессии",
	37: "Пользователь не ответил на интерактив",
	38: "Интерактив не начался",
	39: "Интерактив не закончился",
	40: "Время вышло",
	41: "Время старта интерактива не найдено",
}

var httpCodes = [...]int{ //nolint:gochecknoglobals
	0:  http.StatusInternalServerError,
	1:  http.StatusBadRequest,
	2:  http.StatusBadRequest,
	3:  http.StatusBadRequest,
	4:  http.StatusBadRequest,
	5:  http.StatusBadRequest,
	6:  http.StatusBadRequest,
	7:  http.StatusBadRequest,
	8:  http.StatusBadRequest,
	9:  http.StatusInternalServerError,
	10: http.StatusNotFound,
	11: http.StatusNotFound,
	12: http.StatusNotFound,
	13: http.StatusNotFound,
	14: http.StatusNotFound,
	15: http.StatusBadRequest,
	16: http.StatusUnauthorized,
	17: http.StatusUnauthorized,
	18: http.StatusUnauthorized,
	19: http.StatusUnauthorized,
	20: http.StatusUnauthorized,
	21: http.StatusUnauthorized,
	22: http.StatusBadRequest,
	23: http.StatusBadRequest,
	24: http.StatusForbidden,
	25: http.StatusInternalServerError,
	26: http.StatusNotFound,
	27: http.StatusNotFound,
	28: http.StatusBadRequest,
	29: http.StatusInternalServerError,
	30: http.StatusBadRequest,
	31: http.StatusBadRequest,
	32: http.StatusBadRequest,
	33: http.StatusBadRequest,
	34: http.StatusInternalServerError,
	35: http.StatusForbidden,
	36: http.StatusBadRequest,
	37: http.StatusBadRequest,
	38: http.StatusBadRequest,
	39: http.StatusBadRequest,
	40: http.StatusBadRequest,
	41: http.StatusBadRequest,
}

type CodeError int

func (e CodeError) Error() string {
	code := e - startCode
	if code < 0 || int(code) >= len(codeMessages) {
		return "Неизвестная ошибка"
	}

	return codeMessages[e-startCode]
}

func (e CodeError) GetHTTPCode() int {
	code := e - startCode
	if code < 0 || int(code) >= len(codeMessages) {
		return http.StatusInternalServerError
	}

	return httpCodes[e-startCode]
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
	ErrPresentationNotFound
	ErrPresentationVersionNotFound
	ErrUserNotFound
	ErrInteractiveNotFound
	ErrSlidesNotFound
	InvalidSessionEmptyPresentation
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
	ErrSessionNotFound
	ErrInvalidSlideID
	ErrNotificationFailed
	ErrBadRequest
	ErrInvalidAnswers
	ErrStatusNotFound
	ErrAlreadyVoted
	ErrSessionCreate
	ErrUserIsNotReal
	ErrLimitSession
	ErrUserNotVoteInQuiz
	ErrInteractiveNotStarted
	ErrInteractiveNotFinished
	ErrTimeIsOver
	ErrTimestampNotFound
)
