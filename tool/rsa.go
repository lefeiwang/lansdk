package tool

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"log"
)

//Rsa 密钥对
func GetRsaKeys() (pubKey, priKey []byte) {
	privateKey, _ := rsa.GenerateKey(rand.Reader, 2048) // sise=256
	x509_Privatekey := x509.MarshalPKCS1PrivateKey(privateKey)
	log.Println("私钥: ", x509_Privatekey)

	publicKey := privateKey.PublicKey
	x509_PublicKey, _ := x509.MarshalPKIXPublicKey(&publicKey)
	log.Println("公钥: ", x509_PublicKey)

	return x509_PublicKey, x509_Privatekey
}

// 公钥加密函数-分段
func RsaPubEncryptBlock(plaintext, publicKeyByte []byte) (ciphertext []byte, err error) {
	publicKey, err := x509.ParsePKIXPublicKey(publicKeyByte)
	if err != nil {
		return
	}

	keySize, textSize := publicKey.(*rsa.PublicKey).Size(), len(plaintext)
	pub := publicKey.(*rsa.PublicKey)

	//单次加密的长度需要减掉padding的长度，PKCS1为11
	offSet, once := 0, keySize-11
	buffer := bytes.Buffer{}

	for offSet < textSize {
		endIndex := offSet + once
		if endIndex > textSize {
			endIndex = textSize
		}
		// 加密一部分
		bytesOnce, err := rsa.EncryptPKCS1v15(rand.Reader, pub, plaintext[offSet:endIndex])
		if err != nil {
			return nil, err
		}
		buffer.Write(bytesOnce)
		offSet = endIndex
	}
	ciphertext = buffer.Bytes()
	return
}

// 私钥解密函数-分段
func RsaPrivDecryptBlock(ciphertext, privateKeyByte []byte) (plaintext []byte, err error) {
	privateKey, err := x509.ParsePKCS1PrivateKey(privateKeyByte)
	if err != nil {
		return
	}

	keySize, textSize := privateKey.Size(), len(ciphertext)
	var offSet = 0
	var buffer = bytes.Buffer{}

	for offSet < textSize {
		endIndex := offSet + keySize
		if endIndex > textSize {
			endIndex = textSize
		}
		bytesOnce, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, ciphertext[offSet:endIndex])
		if err != nil {
			return nil, err
		}
		buffer.Write(bytesOnce)
		offSet = endIndex
	}
	plaintext = buffer.Bytes()
	return
}

// 私钥加密函数-分段
func RsaPrivEncryptBlock(plaintext, privateKeyByte []byte) (ciphertext []byte, err error) {
	privateKey, err := x509.ParsePKCS1PrivateKey(privateKeyByte)
	if err != nil {
		return
	}

	keySize, textSize := privateKey.Size(), len(plaintext)
	//单次加密的长度需要减掉padding的长度，PKCS1为11
	offSet, once := 0, keySize-11
	buffer := bytes.Buffer{}

	for offSet < textSize {
		endIndex := offSet + once
		if endIndex > textSize {
			endIndex = textSize
		}
		// 加密一部分
		bytesOnce, err := rsa.SignPKCS1v15(nil, privateKey, crypto.Hash(0), plaintext[offSet:endIndex])
		if err != nil {
			return nil, err
		}
		buffer.Write(bytesOnce)
		offSet = endIndex
	}
	ciphertext = buffer.Bytes()
	return
}

// 公钥解密函数-分段
func RsaPubDecryptBlock(ciphertext, publicKeyByte []byte) (plaintext []byte, err error) {
	publicKey, err := x509.ParsePKIXPublicKey(publicKeyByte)
	if err != nil {
		return
	}

	pub := publicKey.(*rsa.PublicKey)
	keySize, textSize := pub.Size(), len(ciphertext)
	var offSet = 0
	var buffer = bytes.Buffer{}

	for offSet < textSize {
		endIndex := offSet + keySize
		if endIndex > textSize {
			endIndex = textSize
		}
		bytesOnce, err := publicDecrypt(pub, crypto.Hash(0), nil, ciphertext[offSet:endIndex])
		if err != nil {
			return nil, err
		}
		buffer.Write(bytesOnce)
		offSet = endIndex
	}
	plaintext = buffer.Bytes()
	return
}
