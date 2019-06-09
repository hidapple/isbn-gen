package isbn

import "testing"

func TestSearchIdentifier(t *testing.T) {
	tests := []struct {
		name            string
		expected_hit    bool
		expected_prefix string
	}{
		{"Japan", true, "9784"},
		{"jp", true, "9784"},
		{"not_exist", false, ""},
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
	}
}
