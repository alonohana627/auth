package routes

import (
	"fmt"
	"testing"
)

func TestBcrypt(t *testing.T) {
	passwords := []string{"123456", "789456", "asdfefgasjdiow", "GEEEpR0gjgm3I7A", "UhNhE0Bnju8KUMBfvErKvYK9sMm", "0#Fv-nF=#$h6&dODF)86z(sQ)qV", "P;XAD]FzJ4f-9+"}
	var hashedPasswords []string

	for _, password := range passwords {
		hashedPassword, err := bcryptHashPassword(password)
		if err != nil {
			t.Fatal("Hash failed for ", password)
		}
		fmt.Println(hashedPassword)
		hashedPasswords = append(hashedPasswords, hashedPassword)
	}

	for i := range passwords {
		t.Run("Checking "+passwords[i]+" vs "+hashedPasswords[i], func(t *testing.T) {
			if !bcryptCheckPasswordHash(passwords[i], hashedPasswords[i]) {
				t.Errorf("bcryptCheckPasswordHash() Failed for %v,  %v", passwords[i], hashedPasswords[i])
			}
		})
	}
}
