package isbn

// Identifier represents ISBN identifier group and its prefix
type Identifier struct {
	// GroupName is the identifying group name which can be National or Language name
	GroupName string

	// Abbreviation is the shorthand name of the GroupName
	Abbreviation string

	// Prefix is the ISBN prefix assigned to the Group
	Prefix string

	// Identifiers is the ISBN identifier assigned to the Group
	Identifier string
}

// Identifiers is list of ISBN identifer groups
var Identifiers = []*Identifier{
	{
		GroupName:    "English1",
		Abbreviation: "en1",
		Prefix:       "978",
		Identifier:   "0",
	},
	{
		GroupName:    "English2",
		Abbreviation: "en2",
		Prefix:       "978",
		Identifier:   "1",
	},
	{
		GroupName:    "French",
		Abbreviation: "fr",
		Prefix:       "978",
		Identifier:   "2",
	},
	{
		GroupName:    "German",
		Abbreviation: "de",
		Prefix:       "978",
		Identifier:   "3",
	},
	{
		GroupName:    "Japan",
		Abbreviation: "jp",
		Prefix:       "978",
		Identifier:   "4",
	},
	{
		GroupName:    "Russia",
		Abbreviation: "ru",
		Prefix:       "978",
		Identifier:   "5",
	},
	{
		GroupName:    "China",
		Abbreviation: "cn",
		Prefix:       "978",
		Identifier:   "7",
	},
	{
		GroupName:    "Brazil1",
		Abbreviation: "br1",
		Prefix:       "978",
		Identifier:   "65",
	},
	{
		GroupName:    "Brazil2",
		Abbreviation: "br2",
		Prefix:       "978",
		Identifier:   "85",
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
