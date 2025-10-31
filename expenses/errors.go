package expenses
import "errors"
var ErrNOM error = errors.New("not enough money")
var ErrNotExist = errors.New("this id is not exists")
var ErrMonth = errors.New("this Month is not exists")