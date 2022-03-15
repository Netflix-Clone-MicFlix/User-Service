package security

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"
)

// Define salt size
const saltSize = 16

// Generate 16 bytes randomly and securely using the
// Cryptographically secure pseudorandom number generator (CSPRNG)
// in the crypto.rand package
func GenerateRandomSalt() []byte {
	var salt = make([]byte, saltSize)

	_, err := rand.Read(salt[:])

	if err != nil {
		panic(err)
	}

	return salt
}

// Combine password and salt then hash them using the SHA-512
// hashing algorithm and then return the hashed password
// as a base64 encoded string
func HashPassword(password string, salt []byte) string {
	// Convert password string to byte slice
	var passwordBytes = []byte(password)

	// Create sha-512 hasher
	var sha512Hasher = sha512.New()

	// Append salt to password
	passwordBytes = append(passwordBytes, salt...)

	// Write password bytes to the hasher
	sha512Hasher.Write(passwordBytes)

	// Get the SHA-512 hashed password
	var hashedPasswordBytes = sha512Hasher.Sum(nil)

	// Convert the hashed password to a base64 encoded string
	var base64EncodedPasswordHash = base64.URLEncoding.EncodeToString(hashedPasswordBytes)

	return base64EncodedPasswordHash
}

// Check if two passwords match
func CheckPasswordsMatch(hashedPassword, currPassword string,
	salt []byte) bool {
	var currPasswordHash = HashPassword(currPassword, salt)

	return hashedPassword == currPasswordHash
}

// func test() {
// 	// First generate random 16 byte salt
// 	var salt = generateRandomSalt(saltSize)

// 	// Hash password using the salt
// 	var hashedPassword = hashPassword("hello", salt)

// 	fmt.Println("Password Hash:", hashedPassword)
// 	fmt.Println("Salt:", salt)

// 	// Check if passed password matches the original password by hashing it
// 	// with the original password's salt and check if the hashes match
// 	fmt.Println("Password Match:",
// 		doPasswordsMatch(hashedPassword, "hello", salt))
// }
