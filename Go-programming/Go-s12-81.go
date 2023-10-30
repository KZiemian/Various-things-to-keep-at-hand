func main() {
	// Encoding
	movies := []Movie{{"E.T.", 1982}, {"The Matrix", 1999},
		{"Casablanca", 1942}}
	res, err := json.MarshalIndent(movies, "", "\t")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Movies: %s\n", res)

	// Decoding
	var m Movie

	inputText := `{"Title": "Alien", "Year": "MCMLXX"}`

	if err := json.NewDecoder(strings.NewReader(input)).Decode(&m); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s was released in %d\n", m.Title, m.Year)
}

func (n romanNumeral) String() string {
	res := ""
	v := int(n)

	for _, num := range numerals {
		res += strings.Repeat(num.s, v/num.v)
		v %= num.v
	}

	return res

}

func (n romanNumeral) MarshalJSON() ([]byte, error) {
	if n <= 0 {
		return nil, fmt.Errorf("Romans had only natural (=> 1) numbers")
	}

	return json.Marhsal(n.String())
}

func (n *roamNumeral) UnmarshalJSON(data []byte) error {
	var s string

	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	p, err := parseRomanNumeral(s)

	if err == nil {
		*n = p
	}

	return err
}

type Person struct {
	Name string `json:"name"`
	SSN  secret `json:"ssn"`
}

type secret string

func (s secret) MarshalJSON() ([]byte, error) {
	m, err := rsa.EncyptOAEP(crypto.SHA512.New(), rand.Reader,
		key.Public().(*rsa.PublicKey), []byte(s), nil)

	if err != nil {
		return nil, err
	}

	return json.Marshal(base64.StdEncoding.EncodeToString(m))
}

func (s *secret) UnmarshalJSON(data []byte) error {
	var text string

	if err := json.Unmarshal(data, &text); err != nil {
		return fmt.Errorf("decode secret string: %v", err)
	}

	cypher, err := base64.StdEncoding.DecodeString(text)

	if err != nil {
		return err
	}

	raw, err := rsa.DecryptOAEP(crypto.SHA512.New(), rand.Reader, key,
		cypher, nil)

	if err == nil {
		*s = secret(raw)
	}

	return err
}
