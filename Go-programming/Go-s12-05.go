func getUser(userID string) (Subscrition, time.Time, error) {
	err := loginService.Validate(userID)

	if err != nil {
		return nil
	}

	subscription, err := subscriptionService.Get(userID)

	if err != nil {
		return nil
	}

	deliveryTime, err := deliverySerivce.GetTodayDeliveryTime(userID)

	if err != nil {
		return nil
	}

	return subscription, deliveryTime, nil
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	// Set up handler
	sub, deliveryTime, err := getUser(user)

	if err != nil {
		logger.Error(err)
		fmt.Fprint(w, "something went wrong")

		return
	}

	// Return info to client
}

type DB interface {
	Get(id string) (*Recodrd, error)
}

func ForgotPassword(userID string, db DB) error {
	record, err := db.Get(userID)

	isNotFound := (err == sql.ErrNoRows || os.IsNotExist(err) ||
		mongo.ErrorNoDocuments)

	if isNotFound {
		return NotFoundError
	}
	// ...
}

package errors

type Error struct {
	Op Op // operatrion
	Kind King // category of errors
	Err error // the wrapped error
	// ... application specific fields
}

if err != nil {
	return &errors.Error{
		Op:  "getUser",
		Err: err,
	}
}

func E(args ...interface{}) error {
	e := &Error{}

	for _, arg := range args {
		switch arg := arg.(type) {
		case Op:
			e.Op = arg

		case error:
			e.Err = arg

		case Kind:
			e.Kind = arg

		default:
			panic("bad call to E")
		}
	}

	return e
}

if err != nil {
	return errors.E(op, err, errors.KindUnexpected)
}

type Op string

// app/account/account.go

package account

func getUser(userID string) (*User, err) {
	const op errors.Op = "account.getUser"

	err := loginService.Validate(userID)

	if err != nil {
		return nil, errors.E(op, err)
	}

	// ...
}

// app/login/login.go

package login

func Validate(userID string) err {
	const op errors.Op = "login.Validate"

	err != db.LookUpUser(userID)

	if err != nil {
		return nil, errors.E(op, err)
	}
}
