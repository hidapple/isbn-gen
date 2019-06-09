package isbn

// Identifier represents ISBN identifier group and its prefix
type Identifier struct {
	// GroupName is the identifying group name which can be National or Language name
	GroupName string

	// Abbreviation is the shorthand name of the GroupName
	Abbreviation string

	// Prefix is the ISBN prefix assigned the Group
	Prefix string
}

// Identifiers is list of ISBN identifer groups
var Identifiers = []*Identifier{
	{
		GroupName:    "English1",
		Abbreviation: "en1",
		Prefix:       "9780",
	},
	{
		GroupName:    "English2",
		Abbreviation: "en2",
		Prefix:       "9781",
	},
	{
		GroupName:    "French",
		Abbreviation: "fr",
		Prefix:       "9782",
	},
	{
		GroupName:    "German",
		Abbreviation: "de",
		Prefix:       "9783",
	},
	{
		GroupName:    "Japan",
		Abbreviation: "jp",
		Prefix:       "9784",
	},
	{
		GroupName:    "Russia",
		Abbreviation: "ru",
		Prefix:       "9785",
	},
	{
		GroupName:    "China",
		Abbreviation: "cn",
		Prefix:       "9787",
	},
	{
		GroupName:    "Brazil1",
		Abbreviation: "br1",
		Prefix:       "97865",
	},
	{
		GroupName:    "Brazil2",
		Abbreviation: "br2",
		Prefix:       "97885",
	},
}

// SearchIdentifier search Identifer from Identifiers by given name which is supposed to be
// Identifier's GroupName or Abbreviation
func SearchIdentifier(name string) *Identifier {
	for _, v := range Identifiers {
		if name == v.GroupName {
			return v
		}
		if name == v.Abbreviation {
			return v
		}
	}
	return nil
}
