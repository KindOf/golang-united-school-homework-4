package string_sum

import "testing"

func TestStringSum(t *testing.T) {

	tests := []struct {
		name       string
		input      string
		wantOutput string
		wantErr    error
	}{
		{
			name:       "should sum \"3+5\"",
			input:      "3+5",
			wantOutput: "8",
			wantErr:    nil,
		}, {
			name:       "should sum \"-3+5\"",
			input:      "-3+5",
			wantOutput: "2",
			wantErr:    nil,
		}, {
			name:       "should sum \"-3-5\"",
			input:      "-3-5",
			wantOutput: "-8",
			wantErr:    nil,
		}, {
			name:       "should sum \" -3 + 5\"",
			input:      " -3 + 5",
			wantOutput: "2",
			wantErr:    nil,
		}, {
			name:       "should return errorEmptyInput if input is empty",
			input:      "",
			wantOutput: "",
			wantErr:    errorEmptyInput,
		}, {
			name:       "should return errorNotTwoOperands if input has less then two operands",
			input:      "3",
			wantOutput: "",
			wantErr:    errorNotTwoOperands,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOutput, err := StringSum(tt.input)
			if err != tt.wantErr {
				t.Errorf("StringSum() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotOutput != tt.wantOutput {
				t.Errorf("StringSum() gotOutput = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}
