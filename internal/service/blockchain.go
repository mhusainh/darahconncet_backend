package service

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"log"
	"math/big"

	"github.com/mhusainh/DarahConnect/DarahConnectAPI/configs"
	"github.com/mhusainh/DarahConnect/DarahConnectAPI/internal/contracts" // Impor package kontrak hasil generate abigen
	"github.com/mhusainh/DarahConnect/DarahConnectAPI/utils"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// BlockchainService mendefinisikan fungsi-fungsi untuk interaksi dengan blockchain.
type BlockchainService interface {
	// Mengembalikan: Transaction Hash (string), Error
	CreateCertificate(donorAddress, donorName, donorAlamat string) (string, string, error)
}

// Struct implementasi dari interface di atas.
type blockchainService struct {
	client           *ethclient.Client
	contractInstance *contracts.SertifikatDonasi
	privateKey       *ecdsa.PrivateKey
	chainID          *big.Int
}

// NewBlockchainService adalah constructor yang membuat instance baru dari BlockchainService.
// Ia menerima konfigurasi sebagai parameter (Dependency Injection).
func NewBlockchainService(cfg configs.BlockchainConfig) (BlockchainService, error) {
	// Ganti semua log.Fatalf dengan return nil, error

	client, err := ethclient.Dial(cfg.RPCURL)
	if err != nil {
		// return error, bukan mematikan program
		return nil, fmt.Errorf("gagal terhubung ke Ethereum client: %w", err)
	}

	contractAddress := common.HexToAddress(cfg.ContractAddress)
	instance, err := contracts.NewSertifikatDonasi(contractAddress, client)
	if err != nil {
		return nil, fmt.Errorf("gagal membuat instance kontrak: %w", err)
	}

	privateKey, err := crypto.HexToECDSA(cfg.PrivateKey)
	if err != nil {
		return nil, fmt.Errorf("gagal memuat private key: %w", err)
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("gagal mendapatkan ChainID: %w", err)
	}

	fmt.Println("✅ Blockchain service successfully initialized.")

	// Kembalikan service dan error yang nil (tidak ada error)
	return &blockchainService{
		client,
		instance,
		privateKey,
		chainID,
	}, nil
}

// CreateCertificate memanggil fungsi mintSertifikat di smart contract.
func (s *blockchainService) CreateCertificate(donorAddress, donorName, donorAlamat string) (string, string, error) {
	// Buat "transactor" (penanda tangan transaksi) dari private key backend
	auth, err := bind.NewKeyedTransactorWithChainID(s.privateKey, s.chainID)
	if err != nil {
		return "", "", errors.New("gagal membuat transactor")
	}

	certificateNumber, err := utils.GenerateUniqueCertificateNumber()
	if err != nil {
		return "", "", errors.New("Gagal membuat nomor sertifikat")
	}

	// Konversi tipe data Go ke tipe data yang dimengerti Solidity
	pendonorAddress := common.HexToAddress(donorAddress)
	nomorSertifikat, success := new(big.Int).SetString(certificateNumber, 10)
	if !success {
		return "", "", errors.New("failed to convert certificate number to big.Int: " + certificateNumber)
	}

	log.Println("Mengirim transaksi MintSertifikat ke blockchain...")

	// Panggil fungsi dari smart contract (dari file hasil generate abigen)
	tx, err := s.contractInstance.MintSertifikat(auth, pendonorAddress, donorName, nomorSertifikat, donorAlamat)
	if err != nil {
		return "", "", errors.New("gagal memanggil fungsi MintSertifikat: " + err.Error())
	}

	// Ambil Transaction Hash dari objek transaksi!
	txHash := tx.Hash().Hex()
	log.Printf("Transaksi berhasil dikirim! Tx Hash: %s", txHash)
	// Kembalikan hash tersebut untuk disimpan ke database
	return txHash, certificateNumber, nil
}
