;; SPDX-License-Identifier: GPL-3.0-or-later

;; Num typeclass

(define + (lambda (a b)
	"Add two float64 numbers.

	:: (Callable (Float64 Float64) Float64)"
	...))

(define * (lambda (a b)
	"Multiply two float64 numbers.

	:: (Callable (Float64 Float64) Float64)"
	...))

;; Ord typeclass

(define < (lambda (a b)
	"Check if a is less than b.

	:: (Callable (Float64 Float64) Bool)"
	...))

(define > (lambda (a b)
	"Check if a is greater than b.

	:: (Callable (Float64 Float64) Bool)"
	...))
