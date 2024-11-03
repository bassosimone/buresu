;; SPDX-License-Identifier: GPL-3.0-or-later

;; Num typeclass

(declare + (lambda (a b)
	"Add two float64 numbers.

	:: (Callable (Float64 Float64) Float64)"
	...))

(declare * (lambda (a b)
	"Multiply two float64 numbers.

	:: (Callable (Float64 Float64) Float64)"
	...))

;; Ord typeclass

(declare < (lambda (a b)
	"Check if a is less than b.

	:: (Callable (Float64 Float64) Bool)"
	...))

(declare > (lambda (a b)
	"Check if a is greater than b.

	:: (Callable (Float64 Float64) Bool)"
	...))
