package isbn

import "testing"

func TestSearchIdentifier(t *testing.T) {
	tests := []struct {
		name           string
		expectedHit    bool
		expectedPrefix string
		expectedID     string
	}{
		{"Japan", true, "978", "4"},
		{"jp", true, "978", "4"},
		{"not_exist", false, "", ""},
	}

	for i, tt := range tests {
		actual := SearchIdentifier(tt.name)

		if !tt.expectedHit {
			if actual != nil {
				t.Fatalf("case %d: SearchIdentifier should not return any Identifier by %s", i, tt.name)
			}
			continue
		}

		if tt.expectedHit && actual == nil {
			t.Fatalf("case %d: SearchIdentifier should find Identifier by %s", i, tt.name)
		}
		if tt.expectedPrefix != actual.Prefix {
			t.Fatalf("case %d: wrong prefix was found by %s. want=%s, got=%s",
				i, tt.name, tt.expectedPrefix, actual.Prefix)
		}
		if tt.expectedID != actual.Identifier {
			t.Fatalf("case %d: wrong identidier was found by %s. want=%s, got=%s",
				i, tt.name, tt.expectedID, actual.Identifier)
		}
	}
}
