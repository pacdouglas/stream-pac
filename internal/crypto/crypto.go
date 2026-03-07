package crypto

// Encrypter encrypts and decrypts data using AES-256-GCM.
// A random 12-byte nonce is prepended to every ciphertext.
type Encrypter struct {
	key []byte // must be 32 bytes (AES-256)
}

// New creates an Encrypter from a 32-byte key.
func New(key []byte) (*Encrypter, error) {
	panic("not implemented")
}

// DeriveKey derives a 32-byte AES key from a secret and salt using HKDF-SHA256.
// Requires golang.org/x/crypto/hkdf.
func DeriveKey(secret, salt []byte) ([]byte, error) {
	panic("not implemented")
}

// Encrypt encrypts plaintext and returns nonce + ciphertext.
func (e *Encrypter) Encrypt(plaintext []byte) ([]byte, error) {
	panic("not implemented")
}

// Decrypt decrypts ciphertext produced by Encrypt (nonce prepended).
func (e *Encrypter) Decrypt(ciphertext []byte) ([]byte, error) {
	panic("not implemented")
}

// GenerateSalt returns n cryptographically random bytes.
func GenerateSalt(n int) ([]byte, error) {
	panic("not implemented")
}

// MasterKeyFromKeyring retrieves or generates the master key from the OS keychain.
// Uses zalando/go-keyring (Windows Credential Manager, macOS Keychain, Linux Secret Service).
// On first run, generates a random 32-byte key and stores it in the keychain.
func MasterKeyFromKeyring(service, user string) ([]byte, error) {
	panic("not implemented")
}
