package model

import (
	"errors"
	"fmt"
	"regexp"
	"unicode/utf8"
)

type Password string

func (p Password) String() string {
	return string(p)
}

const (
	PasswordMinLength = 8
	PasswordMaxLength = 70
)

type PasswordValidateError error

var (
	ErrPasswordLength    PasswordValidateError = fmt.Errorf("パスワードは%d～%d文字以内にしてください", PasswordMinLength, PasswordMaxLength)
	ErrPasswordUppercase PasswordValidateError = errors.New("パスワードは大文字の英字が1文字以上必要です")
	ErrPasswordLowercase PasswordValidateError = errors.New("パスワードは小文字の英字が1文字以上必要です")
	ErrPasswordNumber    PasswordValidateError = errors.New("パスワードは数字が1文字以上必要です")
	ErrPasswordSymbol    PasswordValidateError = errors.New("パスワードは記号が1文字以上必要です")
)

func (p Password) Validate() error {
	// パスワードの長さが規定の範囲以内である必要がある
	if utf8.RuneCountInString(p.String()) < PasswordMinLength || utf8.RuneCountInString(p.String()) > PasswordMaxLength {
		return ErrPasswordLength
	}

	// 大文字の英字が1文字以上含まれている必要がある
	if !regexp.MustCompile(`[A-Z]`).MatchString(p.String()) {
		return ErrPasswordUppercase
	}

	// 小文字の英字が1文字以上含まれている必要がある
	if !regexp.MustCompile(`[a-z]`).MatchString(p.String()) {
		return ErrPasswordLowercase
	}

	// 数字が1文字以上含まれている必要がある
	if !regexp.MustCompile(`\d`).MatchString(p.String()) {
		return ErrPasswordNumber
	}

	// 記号が1文字以上含まれている必要がある
	if !regexp.MustCompile("[!\"#$%&'()*+,\\-./:;<=>?@[\\\\\\]^_`{|}~]").MatchString(p.String()) {
		return ErrPasswordSymbol
	}

	return nil
}
