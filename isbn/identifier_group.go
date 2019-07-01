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
		GroupName:    "English",
		Abbreviation: "en",
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
		GroupName:    "Brazil",
		Abbreviation: "br",
		Prefix:       "978",
		Identifier:   "65",
	},
	{
		GroupName:    "Czech",
		Abbreviation: "cz",
		Prefix:       "978",
		Identifier:   "80",
	},
	{
		GroupName:    "India",
		Abbreviation: "in",
		Prefix:       "978",
		Identifier:   "81",
	},
	{
		GroupName:    "Norge",
		Abbreviation: "no",
		Prefix:       "978",
		Identifier:   "82",
	},
	{
		GroupName:    "Poland",
		Abbreviation: "pl",
		Prefix:       "978",
		Identifier:   "83",
	},
	{
		GroupName:    "Spain",
		Abbreviation: "es",
		Prefix:       "978",
		Identifier:   "84",
	},
	{
		GroupName:    "Brazil2",
		Abbreviation: "br2",
		Prefix:       "978",
		Identifier:   "85",
	},
	{
		GroupName:    "Serbia",
		Abbreviation: "rs",
		Prefix:       "978",
		Identifier:   "86",
	},
	{
		GroupName:    "Denmark",
		Abbreviation: "dk",
		Prefix:       "978",
		Identifier:   "87",
	},
	{
		GroupName:    "Italy",
		Abbreviation: "it",
		Prefix:       "978",
		Identifier:   "88",
	},
	{
		GroupName:    "SouthKorea",
		Abbreviation: "kr",
		Prefix:       "978",
		Identifier:   "89",
	},
	{
		GroupName:    "Netherlands",
		Abbreviation: "nl",
		Prefix:       "978",
		Identifier:   "90",
	},
	{
		GroupName:    "Sweden",
		Abbreviation: "se",
		Prefix:       "978",
		Identifier:   "91",
	},
	{
		GroupName:    "NGO",
		Abbreviation: "ngo",
		Prefix:       "978",
		Identifier:   "92",
	},
	{
		GroupName:    "India2",
		Abbreviation: "in2",
		Prefix:       "978",
		Identifier:   "93",
	},
	{
		GroupName:    "Netherlands2",
		Abbreviation: "nl2",
		Prefix:       "978",
		Identifier:   "94",
	},
	{
		GroupName:    "French2",
		Abbreviation: "fr2",
		Prefix:       "979",
		Identifier:   "10",
	},
	{
		GroupName:    "SouthKorea2",
		Abbreviation: "kr2",
		Prefix:       "979",
		Identifier:   "11",
	},
	{
		GroupName:    "Italy2",
		Abbreviation: "it2",
		Prefix:       "979",
		Identifier:   "12",
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
