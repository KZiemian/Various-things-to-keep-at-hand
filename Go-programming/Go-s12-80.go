func (ss *ShirtSize) UnmarshalJSON(data []byte) error {
	// Extract the string from data.
	var s string

	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("shirt-size should be a string, got %s",
			data)
	}

	// The rest is equivalent to ParseShirtSize.
	got, ok := map[string]ShirtSize{"XS": XS, "S": S, "M": M, "L": L,
		"XL": XL}[s]

	if !ok {
		return fmt.Errorf("invalid ShirtSize %q", s)
	}

	*ss = got

	return nil
}

func (p *Person) UnmarshalJSON(data []byte) error {
	var aux struct {
		Name string
		Born string    `json:"birthdate"`
		Size ShirtSize `json:"shirt-size"`
	}

	dec := json.NewDecoder(bytes.NewReader(data))

	if err := dec.Decode(&aux); err != nil {
		return fmt.Errorf("decode person: %v", err)
	}

	p.Name = aux.Name
	p.Size = aux.Size
	// ... rest of function omitted ...
}

type Date struct {
	time.Time
}

func (d *Date) UnmarshalJSON(data []byte) error {
	var s string

	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("birthdate should be a string, got %s",
			data)
	}

	t, err := time.Parse("2006/01/02", s)

	if err != nil {
		return fmt.Errorf("invalid date: %v", err)
	}

	d.Time = t

	return nil
}

func (p *Person) UnmarshalJSON(data []byte) error {
	r := bytes.NewReader(data)

	var aux struct {
		Name string
		Born Date      `json:"birthdate"`
		Size ShirtSize `json:"shirt-size"`
	}

	if err := json.NewDecoder(r).Decode(&aux); err != nil {
		return fmt.Errorf("decode person: %v", err)
	}

	p.Name = aux.Name
	p.Size = aux.Size
	p.Born = aux.Born

	return nil
}

func main() {
	var p Person

	dec := json.NewDecoder(strings.NewREader(input))

	if err := dec.Decode(&p); err != nil {
		log.Fatalf("parse person: %v", err)
	}

	fmt.Println(p)
}

type romanNumeral int

type Movie struct {
	Title string
	Year  romanNumeral
}
