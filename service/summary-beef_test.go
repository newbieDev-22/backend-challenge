package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSummaryBeef(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected map[string]int
		wantErr  bool
	}{
		{
			name:  "Simple case with repeated words",
			input: "beef beef pork chicken beef",
			expected: map[string]int{
				"beef":    3,
				"pork":    1,
				"chicken": 1,
			},
			wantErr: false,
		},
		{
			name:  "Case with punctuation",
			input: "beef, pork. chicken, beef.",
			expected: map[string]int{
				"beef":    2,
				"pork":    1,
				"chicken": 1,
			},
			wantErr: false,
		},
		{
			name:     "Empty string case",
			input:    "",
			expected: map[string]int(nil),
			wantErr:  true,
		},
		{
			name:  "Multiple lines case",
			input: "beef pork\nchicken beef\nturkey pork",
			expected: map[string]int{
				"beef":    2,
				"pork":    2,
				"chicken": 1,
				"turkey":  1,
			},
			wantErr: false,
		},
		{
			name: "Long text case 1",
			input: "Buffalo consectetur officia bacon commodo pastrami laboris ullamco picanha alcatra labore aliquip laborum. " +
				"Turkey biltong strip steak venison labore ad, excepteur tempor short ribs consectetur cow. " +
				"Beef kielbasa nisi dolore, labore ball tip do laborum cillum. " +
				"Do frankfurter cupidatat, laboris rump in doner tongue reprehenderit ground round ut veniam. " +
				"Beef ribs laborum buffalo bacon leberkas.",
			expected: map[string]int{
				"Beef":          2,
				"Buffalo":       1,
				"Do":            1,
				"Turkey":        1,
				"ad":            1,
				"alcatra":       1,
				"aliquip":       1,
				"bacon":         2,
				"ball":          1,
				"biltong":       1,
				"buffalo":       1,
				"cillum":        1,
				"commodo":       1,
				"consectetur":   2,
				"cow":           1,
				"cupidatat":     1,
				"do":            1,
				"dolore":        1,
				"doner":         1,
				"excepteur":     1,
				"frankfurter":   1,
				"ground":        1,
				"in":            1,
				"kielbasa":      1,
				"labore":        3,
				"laboris":       2,
				"laborum":       3,
				"leberkas":      1,
				"nisi":          1,
				"officia":       1,
				"pastrami":      1,
				"picanha":       1,
				"reprehenderit": 1,
				"ribs":          2,
				"round":         1,
				"rump":          1,
				"short":         1,
				"steak":         1,
				"strip":         1,
				"tempor":        1,
				"tip":           1,
				"tongue":        1,
				"ullamco":       1,
				"ut":            1,
				"veniam":        1,
				"venison":       1,
			},
			wantErr: false,
		},
		{
			name: "Long text case 2",
			input: "Short loin sirloin in nisi beef ribs beef anim occaecat dolore fatback incididunt laboris. " +
				"Beef ribs laborum buffalo bacon leberkas, id officia velit porchetta t-bone do kevin cillum labore anim. " +
				"Sunt beef ribs minim pork belly anim, ribeye dolor capicola adipisicing incididunt.",
			expected: map[string]int{
				"Beef":        1,
				"Short":       1,
				"Sunt":        1,
				"adipisicing": 1,
				"anim":        3,
				"bacon":       1,
				"beef":        3,
				"belly":       1,
				"buffalo":     1,
				"capicola":    1,
				"cillum":      1,
				"do":          1,
				"dolor":       1,
				"dolore":      1,
				"fatback":     1,
				"id":          1,
				"in":          1,
				"incididunt":  2,
				"kevin":       1,
				"labore":      1,
				"laboris":     1,
				"laborum":     1,
				"leberkas":    1,
				"loin":        1,
				"minim":       1,
				"nisi":        1,
				"occaecat":    1,
				"officia":     1,
				"porchetta":   1,
				"pork":        1,
				"ribeye":      1,
				"ribs":        3,
				"sirloin":     1,
				"t-bone":      1,
				"velit":       1,
			},
			wantErr: false,
		},
		{
			name: "Long text case 3",
			input: "Kevin cupim swine reprehenderit beef ribs eu magna. " +
				"Chuck ut shankle in jowl laboris. " +
				"Dolor flank pariatur dolore cupidatat. " +
				"Venison reprehenderit ut, turducken commodo tri-tip pork chop culpa lorem beef laborum mollit.",
			expected: map[string]int{
				"Chuck":         1,
				"Dolor":         1,
				"Kevin":         1,
				"Venison":       1,
				"beef":          2,
				"chop":          1,
				"commodo":       1,
				"culpa":         1,
				"cupidatat":     1,
				"cupim":         1,
				"dolore":        1,
				"eu":            1,
				"flank":         1,
				"in":            1,
				"jowl":          1,
				"laboris":       1,
				"laborum":       1,
				"lorem":         1,
				"magna":         1,
				"mollit":        1,
				"pariatur":      1,
				"pork":          1,
				"reprehenderit": 2,
				"ribs":          1,
				"shankle":       1,
				"swine":         1,
				"tri-tip":       1,
				"turducken":     1,
				"ut":            2,
			},
			wantErr: false,
		},
	}

	service := NewSummaryBeefService()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := service.GetSummaryBeef(tt.input)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.expected, result.Beef, "Word counts should match expected values")
		})
	}
}
