package utils

import (
	"errors"
	"fmt"
	"log"

	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/natefinch/lumberjack"

	"github.com/go-playground/validator/v10"

	"golang.org/x/crypto/bcrypt"
)

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

func GetRandomHex(length int) string {
	rand.Seed(time.Now().UnixNano())
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ01234566789"
	sb := strings.Builder{}
	sb.Grow(length)
	for i := 0; i < length; i++ {
		sb.WriteByte(charset[rand.Intn(len(charset))])
	}
	res := sb.String()
	return res
}

func GenerateRandomHashedPassword() (string, []byte) {
	pwd := GetRandomHex(12)
	b := []byte(pwd)
	hpwd, err := bcrypt.GenerateFromPassword(b, bcrypt.DefaultCost)
	if err != nil {
		log.Print("Failed to hash password.", err)
		return "", []byte{}
	}
	return pwd, hpwd
}

func ValidateStruct(structVar any) (string, error) {
	var (
		errorResp []*ErrorResponse
	)

	validate := validator.New()

	err := validate.Struct(structVar)
	if err != nil {
		msg := "Validation Failed: FieldNotFound/UnexpectedValue `%s`"

		for _, err := range err.(validator.ValidationErrors) {

			var element ErrorResponse

			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()

			errorResp = append(errorResp, &element)
			fmt.Printf("errorResp: %v\n", errorResp)

			msg = fmt.Sprintf(msg, strings.Split(element.FailedField, ".")[1])
		}
		return msg, errors.New("validation failed")
	} else {
		return "", nil
	}

}

type Logger struct {
	Info    *log.Logger
	Error   *log.Logger
	Debug   *log.Logger
	Request *log.Logger
}

func InitiateLoggers(logFilePath *string) *Logger {

	infoLogger := log.New(os.Stderr, "INFO : ", log.Ldate|log.Lshortfile)
	errorLogger := log.New(os.Stderr, "EROR : ", log.Ldate|log.Lshortfile)
	debugLogger := log.New(os.Stderr, "DEBG : ", log.Ldate|log.Lshortfile)
	requestLogger := log.New(os.Stderr, "RQST : ", log.Ldate|log.Lshortfile)

	if logFilePath != nil && *logFilePath != "" {
		f, err := os.OpenFile(*logFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("error opening file: %v", err)
		}
		defer f.Close()

		infoLogger.SetOutput(&lumberjack.Logger{
			Filename:   *logFilePath,
			MaxSize:    1,  // megabytes after which new file is created
			MaxBackups: 3,  // number of backups
			MaxAge:     28, //days
		})

		errorLogger.SetOutput(&lumberjack.Logger{
			Filename:   *logFilePath,
			MaxSize:    1,  // megabytes after which new file is created
			MaxBackups: 3,  // number of backups
			MaxAge:     28, //days
		})

		debugLogger.SetOutput(&lumberjack.Logger{
			Filename:   *logFilePath,
			MaxSize:    1,  // megabytes after which new file is created
			MaxBackups: 3,  // number of backups
			MaxAge:     28, //days
		})

		requestLogger.SetOutput(&lumberjack.Logger{
			Filename:   *logFilePath,
			MaxSize:    1,  // megabytes after which new file is created
			MaxBackups: 3,  // number of backups
			MaxAge:     28, //days
		})
	}

	return &Logger{
		Info:    infoLogger,
		Error:   errorLogger,
		Debug:   debugLogger,
		Request: requestLogger,
	}
}
