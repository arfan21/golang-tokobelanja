package constant

import "errors"

var ErrorInvalidLogin = errors.New("invalid email or password")
var ErrorEmailAlreadyExists = errors.New("email already exists")
var ErrorInvalidRole = errors.New("invalid role")
var ErrorCategoryDoesNotExists = errors.New("category doesn't exist")
var ErrorOwnership = errors.New("invalid ownership")
var ErrorOutOfStock = errors.New("out of stock")
var ErrorBalance = errors.New("balance is not sufficient")
