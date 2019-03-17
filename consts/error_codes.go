package consts

import "net/http"

const CodeDbError = 1
const CodeWrongProductIdParameter = 2

const CodeSuccess = http.StatusOK
const CodeError = http.StatusBadRequest

const CodeEmptyProductName = 1001
const CodeEmptyProductCategory = 1002
const CodeEmptyProductType = 1003
const CodeCantSaveImage = 1004