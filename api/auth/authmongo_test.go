package auth_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"errors"
	. "github.com/globocom/huskyCI/api/auth"
	"github.com/globocom/huskyCI/api/types"
	"hash"
)

type FakeGen struct {
	expectedHash            string
	expectedPbkdf2          types.User
	expectedGetCredsError   error
	expectedDecodedSalt     []byte
	expectedDecodeSaltError error
}

func (fG *FakeGen) GenHashValue(value, salt []byte, iter, keyLen int, h hash.Hash) string {
	return fG.expectedHash
}

func (fG *FakeGen) DecodeSaltValue(salt string) ([]byte, error) {
	return fG.expectedDecodedSalt, fG.expectedDecodeSaltError
}

func (fG *FakeGen) GetCredsFromDB(username string) (types.User, error) {
	return fG.expectedPbkdf2, fG.expectedGetCredsError
}

var _ = Describe("Authmongo", func() {
	Context("When hash algorithm chosen is not valid", func() {
		It("Should return the expected error and a nil string", func() {
			fakeGen := FakeGen{
				expectedHash: "nothing",
			}
			pbkdf2Client := ClientPbkdf2{
				HashGen:      &fakeGen,
				Salt:         "mystrongsalt",
				Iterations:   1,
				KeyLen:       12,
				HashFunction: "sha1",
			}
			hashVal, err := pbkdf2Client.GetHashedPass("mypass")
			Expect(hashVal).To(Equal(""))
			Expect(err).To(Equal(errors.New("Failed to generate a hash! It doesn't meet all criteria")))
		})
	})
	Context("When hash algorithm chosen is valid", func() {
		It("Should return a nil error and the expected string", func() {
			fakeGen := FakeGen{
				expectedHash:            "MyHashedString",
				expectedDecodeSaltError: nil,
				expectedDecodedSalt:     []byte("EncodedSalt"),
			}
			pbkdf2Client := ClientPbkdf2{
				HashGen:      &fakeGen,
				Salt:         "mystrongsalt",
				Iterations:   1,
				KeyLen:       12,
				HashFunction: "sha256",
			}
			hashVal, err := pbkdf2Client.GetHashedPass("mypass")
			Expect(hashVal).To(Equal("MyHashedString"))
			Expect(err).To(BeNil())
		})
		It("Should return a nil error and the expected string for sha512", func() {
			fakeGen := FakeGen{
				expectedHash:            "MyHashedString",
				expectedDecodeSaltError: nil,
				expectedDecodedSalt:     []byte("EncodedSalt"),
			}
			pbkdf2Client := ClientPbkdf2{
				HashGen:      &fakeGen,
				Salt:         "mystrongsalt",
				Iterations:   1,
				KeyLen:       12,
				HashFunction: "sha512",
			}
			hashVal, err := pbkdf2Client.GetHashedPass("mypass")
			Expect(hashVal).To(Equal("MyHashedString"))
			Expect(err).To(BeNil())
		})
		It("Should return a nil error and the expected string for sha3_224", func() {
			fakeGen := FakeGen{
				expectedHash:            "MyHashedString",
				expectedDecodeSaltError: nil,
				expectedDecodedSalt:     []byte("EncodedSalt"),
			}
			pbkdf2Client := ClientPbkdf2{
				HashGen:      &fakeGen,
				Salt:         "mystrongsalt",
				Iterations:   1,
				KeyLen:       12,
				HashFunction: "Sha3_224",
			}
			hashVal, err := pbkdf2Client.GetHashedPass("mypass")
			Expect(hashVal).To(Equal("MyHashedString"))
			Expect(err).To(BeNil())
		})
		It("Should return a nil error and the expected string", func() {
			fakeGen := FakeGen{
				expectedHash:            "MyHashedString",
				expectedDecodeSaltError: nil,
				expectedDecodedSalt:     []byte("EncodedSalt"),
			}
			pbkdf2Client := ClientPbkdf2{
				HashGen:      &fakeGen,
				Salt:         "mystrongsalt",
				Iterations:   1,
				KeyLen:       12,
				HashFunction: "SHA3_256",
			}
			hashVal, err := pbkdf2Client.GetHashedPass("mypass")
			Expect(hashVal).To(Equal("MyHashedString"))
			Expect(err).To(BeNil())
		})
		It("Should return a nil error and the expected string", func() {
			fakeGen := FakeGen{
				expectedHash:            "MyHashedString",
				expectedDecodeSaltError: nil,
				expectedDecodedSalt:     []byte("EncodedSalt"),
			}
			pbkdf2Client := ClientPbkdf2{
				HashGen:      &fakeGen,
				Salt:         "mystrongsalt",
				Iterations:   1,
				KeyLen:       12,
				HashFunction: "Sha3_384",
			}
			hashVal, err := pbkdf2Client.GetHashedPass("mypass")
			Expect(hashVal).To(Equal("MyHashedString"))
			Expect(err).To(BeNil())
		})
		It("Should return a nil error and the expected string", func() {
			fakeGen := FakeGen{
				expectedHash:            "MyHashedString",
				expectedDecodeSaltError: nil,
				expectedDecodedSalt:     []byte("EncodedSalt"),
			}
			pbkdf2Client := ClientPbkdf2{
				HashGen:      &fakeGen,
				Salt:         "mystrongsalt",
				Iterations:   1,
				KeyLen:       12,
				HashFunction: "sha3_512",
			}
			hashVal, err := pbkdf2Client.GetHashedPass("mypass")
			Expect(hashVal).To(Equal("MyHashedString"))
			Expect(err).To(BeNil())
		})
	})
	Context("When one of the required fields for PBKDF2 is not valid", func() {
		It("Should return an the expected error and an empty hash for an empty salt", func() {
			fakeGen := FakeGen{
				expectedHash:            "MyHashedString",
				expectedDecodeSaltError: nil,
				expectedDecodedSalt:     []byte("EncodedSalt"),
			}
			pbkdf2Client := ClientPbkdf2{
				HashGen:      &fakeGen,
				Iterations:   1,
				KeyLen:       12,
				HashFunction: "sha256",
			}
			hashVal, err := pbkdf2Client.GetHashedPass("mypass")
			Expect(hashVal).To(Equal(""))
			Expect(err).To(Equal(errors.New("Failed to generate a hash! It doesn't meet all criteria")))
		})
		It("Should return an the expected error and an empty hash for a 0 iteration", func() {
			fakeGen := FakeGen{
				expectedHash: "MyHashedString",
			}
			pbkdf2Client := ClientPbkdf2{
				HashGen:      &fakeGen,
				Salt:         "ValidSalt",
				KeyLen:       12,
				HashFunction: "SHA224",
			}
			hashVal, err := pbkdf2Client.GetHashedPass("mypass")
			Expect(hashVal).To(Equal(""))
			Expect(err).To(Equal(errors.New("Failed to generate a hash! It doesn't meet all criteria")))
		})
		It("Should return an the expected error and an empty hash for a 0 keyLength", func() {
			fakeGen := FakeGen{
				expectedHash: "MyHashedString",
			}
			pbkdf2Client := ClientPbkdf2{
				HashGen:      &fakeGen,
				Salt:         "ValidSalt",
				Iterations:   1,
				KeyLen:       0,
				HashFunction: "Sha384",
			}
			hashVal, err := pbkdf2Client.GetHashedPass("mypass")
			Expect(hashVal).To(Equal(""))
			Expect(err).To(Equal(errors.New("Failed to generate a hash! It doesn't meet all criteria")))
		})
	})
	Context("When GetCredsFromDB return an error for GetPassFromDB", func() {
		It("Should return the expected error and a nil string", func() {
			fakeGen := FakeGen{
				expectedPbkdf2:        types.User{},
				expectedGetCredsError: errors.New("Could not return credentials from DB"),
			}
			pbkdf2Client := ClientPbkdf2{
				HashGen: &fakeGen,
			}
			pass, err := pbkdf2Client.GetPassFromDB("husky")
			Expect(pass).To(Equal(""))
			Expect(err).To(Equal(errors.New("Could not return credentials from DB")))
		})
	})
	Context("When GetCredsFromDB return nil for GetPassFromDB", func() {
		It("Should return a nil error with the expected PBKDF2 parameters", func() {
			fakeGen := FakeGen{
				expectedPbkdf2: types.User{
					HashFunction: "sha512",
					Iterations:   500,
					KeyLen:       1024,
					Salt:         "MyComplexSalt",
					Password:     "MyHashedPassword",
				},
				expectedGetCredsError: nil,
			}
			pbkdf2Client := ClientPbkdf2{
				HashGen: &fakeGen,
			}
			pass, err := pbkdf2Client.GetPassFromDB("husky")
			Expect(pass).To(Equal("MyHashedPassword"))
			Expect(pbkdf2Client.HashFunction).To(Equal("sha512"))
			Expect(pbkdf2Client.Iterations).To(Equal(500))
			Expect(pbkdf2Client.KeyLen).To(Equal(1024))
			Expect(err).To(BeNil())
		})
	})
	Context("When GetPassFromDB is called with and returns all PBKDF2", func() {
		It("Should return the expected hashed pass and PBKDF2 params well validated", func() {
			FakeGen := FakeGen{
				expectedPbkdf2: types.User{
					HashFunction: "sha256",
					Iterations:   500,
					KeyLen:       1024,
					Salt:         "MyComplexSalt",
					Password:     "MyHashedPassword",
				},
				expectedGetCredsError:   nil,
				expectedHash:            "MyHashedPassword",
				expectedDecodeSaltError: nil,
				expectedDecodedSalt:     []byte("MyComplexSalt"),
			}
			pbkdf2Client := &ClientPbkdf2{
				HashGen: &FakeGen,
			}
			pass, errGetPass := pbkdf2Client.GetPassFromDB("husky")
			hashedPass, errGetHashed := pbkdf2Client.GetHashedPass("notHashedPass")
			Expect(pbkdf2Client.HashFunction).To(Equal("sha256"))
			Expect(pbkdf2Client.Iterations).To(Equal(500))
			Expect(pbkdf2Client.KeyLen).To(Equal(1024))
			Expect(pbkdf2Client.Salt).To(Equal("MyComplexSalt"))
			Expect(pass).To(Equal("MyHashedPassword"))
			Expect(hashedPass).To(Equal("MyHashedPassword"))
			Expect(errGetHashed).To(BeNil())
			Expect(errGetPass).To(BeNil())
		})
	})
})
