package version

import "testing"

func TestGiven_Both_Current_And_Required_Versions_Are_Identical(t *testing.T) {
	isValid, _ := Valid("1.1.1", "1.1.1")

	if isValid == false {
		t.Errorf("Valid was incorrect, got: %v, want: %v.", isValid, true)
	}
}

func TestGiven_Then_Current_Is_Less_Than_The_Required_Version(t *testing.T) {
	isValid, _ := Valid("1.1.1", "2.2.2")

	if isValid == true {
		t.Errorf("Valid was incorrect, got: %v, want: %v.", isValid, false)
	}
}

func TestGiven_I_Have_A_Minor_Current_Version_And_Patch_Required_Version(t *testing.T) {
	isValid, _ := Valid("1.1", "1.1.1")

	if isValid == true {
		t.Errorf("Valid was incorrect, got: %v, want: %v.", isValid, true)
	}
}
