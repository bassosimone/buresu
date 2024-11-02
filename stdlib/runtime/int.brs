;; SPDX-License-Identifier: GPL-3.0-or-later

;; Num typeclass

(define + (lambda (a b)
	"Add two integer numbers.

	:: (Callable (Int Int) Int)"
	...))

(define * (lambda (a b)
	"Multiply two integer numbers.

	:: (Callable (Int Int) Int)"
	...))

;; Ord typeclass

(define < (lambda (a b)
	"Check if a is less than b.

	:: (Callable (Int Int) Bool)"
	...))

(define > (lambda (a b)
	"Check if a is greater than b.

	:: (Callable (Int Int) Bool)"
	...))
