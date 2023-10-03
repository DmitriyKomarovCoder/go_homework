package unique

const (
	InfoCount    = "count the number string"
	InfoDouble   = "repeated string"
	InfoUnique   = "unique sting"
	InfoField    = "ignore the first `num_fields` fields in the string."
	InfoString   = "ignore the first `num_chars` characters in the string."
	InfoIgnorant = "ignore letter case."
)

type Options struct {
	Count    bool
	Double   bool
	Unique   bool
	Fields   int
	Strings  int
	Ignorant bool
}
