package main

var data = `
{
  "user": "Scrooge McDuck",
  "type": "deposit",
  "amount": 1000000.3
  }
  `
type Request struct {
  Login string `json:"user"`,
  Type strign `json:"type"`,
  Amount float64 `json:"amount"`
}

func main() {
  rdr := bytes.NewBufferString(data)

  dec := json.NewDecoder(rdr)
  req := &Request{}
  if err := dec.Decode(req); err != nil {
    log.Fatalf("error: cant`t decode - %s", err)
  }
  fmt.Printf("got: %+v\n", req)
}
