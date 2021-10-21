package main

import (
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
	"math/big"
	"math/rand"
	"os"
	"time"
)

func getPublicKey(env_var string) *rsa.PublicKey {

	N := new(big.Int)

	public_key_str := os.Getenv(env_var)
	N.SetString(public_key_str, 16)

	public_key := new(rsa.PublicKey)
	public_key.N = N
	public_key.E = 65537

	return public_key
}

func getPrivateKey(private_key_env string, public_key_env string) *rsa.PrivateKey {
	private_exponent := new(big.Int)

	private_key_str := os.Getenv(private_key_env)
	private_exponent.SetString(private_key_str, 16)

	public_modulus := getPublicKey(public_key_env)

	private_key := new(rsa.PrivateKey)
	private_key.PublicKey = *public_modulus
	private_key.D = private_exponent

	return private_key
}

func encrypt(text string,
	public_key_env string) string {

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	public_key := getPublicKey(public_key_env)
	label := []byte("plain")

	cipher_text, _ := rsa.EncryptOAEP(sha256.New(), rng, public_key, []byte(text), label)

	return string(cipher_text)
}

func decrypt(text string,
	private_key_env string,
	public_key_env string) string {

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	private_key := getPrivateKey(private_key_env, public_key_env)
	label := []byte("plain")

	plain_text, _ := rsa.DecryptOAEP(sha256.New(), rng, private_key, []byte(text), label)

	return string(plain_text)

}

func main() {
	plain := "salmon is a very cool fish"

	os.Setenv("RSA_PUB_KEY", "a8b3b284af8eb50b387034a860f146c4919f318763cd6c5598c8ae4811a1e0abc4c7e0b082d693a5e7fced675cf4668512772c0cbc64a742c6c630f533c8cc72f62ae833c40bf25842e984bb78bdbf97c0107d55bdb662f5c4e0fab9845cb5148ef7392dd3aaff93ae1e6b667bb3d4247616d4f5ba10d4cfd226de88d39f16fb")
	os.Setenv("RSA_PRIV_KEY", "53339cfdb79fc8466a655c7316aca85c55fd8f6dd898fdaf119517ef4f52e8fd8e258df93fee180fa0e4ab29693cd83b152a553d4ac4d1812b8b9fa5af0e7f55fe7304df41570926f3311f15c4d65a732c483116ee3d3d2d0af3549ad9bf7cbfb78ad884f84d5beb04724dc7369b31def37d0cf539e9cfcdd3de653729ead5d1")

	// Set env is not thread safe
	time.Sleep(2 * time.Second)

	encrypted := encrypt(plain, "RSA_PUB_KEY")
	decrypted := decrypt(encrypted, "RSA_PRIV_KEY", "RSA_PUB_KEY")

	fmt.Println("Encrypted", encrypted, "Original", plain)
	fmt.Println("Decrypted", decrypted)
}
