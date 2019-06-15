package isbn

import "testing"

func TestSearchIdentifier(t *testing.T) {
	tests := []struct {
		name            string
		expected_hit    bool
		expected_prefix string
		expected_id     string
	}{
		{"Japan", true, "978", "4"},
		{"jp", true, "978", "4"},
		{"not_exist", false, "", ""},
	}

	for i, tt := range tests {
		actual := SearchIdentifier(tt.name)

		if !tt.expected_hit {
			if actual != nil {
				t.Fatalf("case %d: SearchIdentifier should not return any Identifier by %s", i, tt.name)
			}
			continue
		}

		if tt.expected_hit && actual == nil {
			t.Fatalf("case %d: SearchIdentifier should find Identifier by %s", i, tt.name)
		}
		if tt.expected_prefix != actual.Prefix {
			t.Fatalf("case %d: wrong prefix was found by %s. want=%s, got=%s",
				i, tt.name, tt.expected_prefix, actual.Prefix)
		}
		if tt.expected_id != actual.Identifier {
			t.Fatalf("case %d: wrong identidier was found by %s. want=%s, got=%s",
				i, tt.name, tt.expected_id, actual.Identifier)
		}
	}
}
