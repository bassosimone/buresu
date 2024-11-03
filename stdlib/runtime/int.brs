;; SPDX-License-Identifier: GPL-3.0-or-later

;; Num typeclass

(declare + (lambda (a b)
	"Add two integer numbers.

	:: (Callable (Int Int) Int)"
	...))

(declare * (lambda (a b)
	"Multiply two integer numbers.

	:: (Callable (Int Int) Int)"
	...))

;; Ord typeclass

(declare < (lambda (a b)
	"Check if a is less than b.

	:: (Callable (Int Int) Bool)"
	...))

(declare > (lambda (a b)
	"Check if a is greater than b.

	:: (Callable (Int Int) Bool)"
	...))
