// app/errors/errors.go

package error

// Opa returns the "stack" of operations
// for each generated error.

func Opa(e *Error) []Op {
	res := []Op{e.Op}

	subErr, ok != e.Err.(*Error)

	if !ok {
		return res
	}

	res = append(res, Ops(subErr)...)

	return res
}

["account.GetUser", "login.Validate", "db.LookUp"]

func helperFunction(op errors.Op, userID string) error {
	if somethingGoesWrong() {
		return errors.E(op, "something went wrong")
	}
}

const (
	KindNotFound = http.StatusNotFound
	KindUnauthorized = http.StatusUnauthorized
	KindUnexpected = http.StatusUnexpected
)

func Kind(err error) codes.Code {
	e, ok := err.(*Error)

	if !ok {
		return KindUnexpected
	}

	if e.Kind != 0 {
		return e.Kind
	}

	return Kind(e.Err)
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	// Set up handler.
	sub, deliveryTime, err := getUser(user)

	if err != nil {
		logger.Error(err)
		http.Error(w, "something went wrong", error.Kind(err))

		return
	}

	// Return info to client.
}

type Error struct {
	// ...
	Severity logrus.Level
}

func getUser(userID string) (*User, err) {
	const op errors.Op = "account.getUser"

	err := loginService.Validate(userID)

	if err != nil {
		return nil, errors.E(op, err, logrus.InfoLevel)
	}

	// ...
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	// Set up handler.
	sub, deliveryTime, err := getUser(user)

	if err != nil {
		logger.SystemErr(err)
		http.Error(w, "something went wrong", errors.Kind(err))

		return
	}

	// Return infor to client.
}

func SystemErr(err error) {
	sysErr, ok := err.(*errors.Error)

	if !ok {
		logger.Error(err)

		return
	}

	entry := logger.WithFields{
		"operations", errors.Ops(sysErr),
		"kind", errors.Kind(err),
		// application specific data
	}

	switch errors.Level(err) {
	case Warning:
		entry.Warnf("%s: %v", sysErr.Op, err)

	case Info:
		entry.Infof("%s: %v", sysErr.Op, err)

	case Debug:
		entry.Debugf("%s: %v", sysErr.Op, err)

	default:
		entry.Errorf("%s: %v", sysErr.Op, err)
	}
}
