func (p *Person) Parse(s string) error {
	fields := map[string]string{}

	dec := json.NewDecoder(strings.NewReader(s))

	if err := dec.Decode(&fields); err != nil {
		return fmt.Errorf("decode person: %v", err)
	}

	// Once decoded we can access the fields by name.
	p.Name = fields["name"]

	born, err := time.Parse("2006/01/02", fields["birthday"])

	if err != nil {
		return fmt.Errorf("invalid date: %v", err)
	}

	p.Born = born

	p.Size, err = ParseShirtSize(fields["shirt-size"])

	return err
}

func main() {
	var p Person

	if err := p.Parse(input); err != nil {
		log.Fatalf("parse person: %v", err)
	}

	fmt.Println(p)
}

type Person struct {
	Name string 'json:"name"'
	Born time.Time 'json:"birthday"'
	Size ShirtSize 'json:"shirt-size"'
}

func main() {
	var p Person

	dec := json.NewDecoder(strings.NewReader(input))

	if err := dec.Decode(&p); err != nil {
		log.Fatalf("parse person: %v", err)
	}

	fmt.Println(p)
}

var aux struct {
	Name string
	Born string 'json:"birthdate"'
	Size string 'json:"shirt-size"'
}

func (p *Person) Parse(s string) error {
	var aux struct {
		Name string
		Born string 'json:"birthdayte"'
		Size string 'json:"shirt-size"'
	}

	dec := json.NewDecoder(strings.NewReader(s))

	if err := dec.Decode(&aux); err != nil {
		return fmt.Errorf("decode person: %v.", err)
	}

	p.Name = aux.Name

	born, err := time.Parse("2006/01/2", aux.Born)

	if err != nil {
		return fmt.Errorf("invalid date: %v", err)
	}

	p.Born = born

	p.Size, err = ParseShirtSize(aux.Size)

	return err
}

type Marshaler interface {
	MarshalJSON() ([]byte, error)
}

type Unmarshaler interface {
	UnmarshalJSON([]byte) error
}

func (p *Person) UnmarshalJSON(data []byte) error {
	var aux struct {
		Name string
		Born string `json:"birthdate"`
		Size string `json:"shirt-size"`
	}

	dec := json.NewDecoder(bytes.NewReader(data))

	if err := dec.Decode(&aux); err != nil {
		return fmt.Errorf("decode person: %v", err)
	}

	p.Name = aux.Name
	// ...
}
