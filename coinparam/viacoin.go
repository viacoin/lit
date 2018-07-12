package coinparam

import (
	"github.com/mit-dci/lit/btcutil/chaincfg/chainhash"
	"github.com/mit-dci/lit/wire"
	"time"
)

var ViacoinParams = Params{
	Name:          "via",
	NetMagicBytes: 0xcbc6680f,
	DefaultPort:   "5223",
	DNSSeeds: []string{
		"seed.viacoin.net",
		"viaseeder.barbatos.fr",
		"mainnet.viacoin.net",
	},

	// chain params
	StartHeader:      [80]byte{},
	StartHeight:      0x02,        // TODO
	AssumeDiffBefore: 0,           // TODO
	DiffCalcFunction: diffBitcoin, // TODO
	MinHeaders:       4032,
	FeePerByte:       100,
	GenesisBlock:     &ViacoinGenesisBlock,
	GenesisHash:      &ViacoinGenesisHash,
	PowLimit:         liteCoinTestNet4PowLimit, // TODO

	PowLimitBits:             0x1e01ffff,
	CoinbaseMaturity:         100,
	SubsidyReductionInterval: 657000,
	TargetTimespan:           time.Hour * 24 * 14,
	TargetTimePerBlock:       time.Second * 24,
	RetargetAdjustmentFactor: 4,
	ReduceMinDifficulty:      false,
	MinDiffReductionTime:     0,
	GenerateSupported:        false,

	// Checkpoints ordered from oldest to newest.
	Checkpoints: []Checkpoint{
		{3901, newHashFromStr("39c94020b653871bbcb29c9489bb12f84c5e89adef75f2e5c5f59e88869129d9")},
		{40821, newHashFromStr("e0471675f9c98aa5ed321ed951d140d4603d96254a4ca9fbca07b8da5ee11954")},
		{41433, newHashFromStr("627e18cc08a276282781705bac09508992dc8b665391edd7bde8a601f011954c")},
		{44606, newHashFromStr("5ceeec38564a36ee3e1e5404970f5715efe0420e92c8e92bedfdfef782c49320")},
		{200000, newHashFromStr("ccd02ffc90ee744cc90fa95aafb75e793e306dac533abf6162c4c5be3c7b80ad")},
		{400000, newHashFromStr("7884c4cc5564a2414d6b9eaebd6cb2e4df47bcce7eb7f5c24b19f2d3bf961fd3")},
		{600000, newHashFromStr("2385e13879304c4223b479267b3a8aca709b3556d0a00b431efc5b4174080d91")},
		{800000, newHashFromStr("964a15c07915e3f1935e28d66f869fba7d63862ce1b494b3ebbb38bb16092f93")},
		{1000000, newHashFromStr("1b187fb9e32e2ce80fc7a151254ef331b225a22dd2f2407d8e90b6276b252047")},
		{1200000, newHashFromStr("db284e629a1ceb3cbf563293e23adecbe57bc0f3a6999b53598c5cdd5f6a74db")},
		{1400000, newHashFromStr("a45e9c776a911978e22cc9c9d01f92c30836bffb6162cffdf96df75e21e1e97e")},
		{1600000, newHashFromStr("26976af79d69b40c23d22a10c9b0440e7ab450dc390128996abdde5704caef81")},
		{1800000, newHashFromStr("a6d494764b8dc454d1d5178fb4adef06784cc5713c91709f90d42d7515bcf2f6")},
		{2000000, newHashFromStr("f363858fb9ef0ecb0c0ac1c690856129863ce65563581adef000c95733ca7fd9")},
		{2200000, newHashFromStr("edb79185486f009823e3eb5a9e8ba877fd923ce278d2b235586cbde47534b1aa")},
		{2400000, newHashFromStr("34f603a7dee67c2c298d607bb40b79a9641b0618671bcacf6c33c77a8f39b0ac")},
		{2600000, newHashFromStr("162c711903353c479d2c4db85eaaa8e02335b337b3065076edf6315e54ecad90")},
		{2800000, newHashFromStr("ab66f0ead8af2420ff28d23836fc3e6aaefb5b04d56ea8ea1e37dc289eb5f5ce")},
		{3000000, newHashFromStr("6382616446e421105760a83ac69ddffd2b13678b3f12895fea2a4a830cbd5e11")},
		{3200000, newHashFromStr("754d508c62c4a8d8c0abc7249f801b454aa016d628af0b18838ca31ea8c24799")},
		{3470255, newHashFromStr("c3f75474ed8171bc303b4cd7ff269592534c0ca8d76fceed3fdbf8abf40d147a")},
		{3600000, newHashFromStr("8bc0e14a2fc7eea4fd560fbdaa5e4162d352cc61766a2bb1b49218766646317e")},
		{3800000, newHashFromStr("36a845e155c750762f16e5a95c91dbfce58e0dcf9ee2b5e5e4ffe38a0d609111")},
		{4000000, newHashFromStr("9b1fa0fa4183c12480ff04aeaadc03051b9fe94d07d444fc454d58ed0b846534")},
		{4200000, newHashFromStr("4ad4e7a68d2510fc28ede94a13bdd73cf321135b85168298b5c67bd63c700d92")},
		{4400000, newHashFromStr("8a46b52e986f04b07e732caefdfb4c9083a9b88240cd7b3c0e48556c07b62f1a")},
		{4600000, newHashFromStr("6ae5f0132613bc0a61167ff78ba3309f0f56e06960099b1a4c99658671abd8c2")},
	},

	// Enforce current block version once majority of the network has
	// upgraded.
	// 51% (51 / 100)
	// Reject previous block versions once a majority of the network has
	// upgraded.
	// 75% (75 / 100)
	BlockEnforceNumRequired: 51,
	BlockRejectNumRequired:  75,
	BlockUpgradeNumToCheck:  100,

	// Mempool parameters
	RelayNonStdTxs: true,

	// Address encoding magics
	PubKeyHashAddrID: 0x47, // starts with V
	ScriptHashAddrID: 0x21,
	Bech32Prefix:     "via",
	PrivateKeyID:     0xc7,

	// BIP44 coin type used in the hierarchical deterministic path for
	// address generation.
	HDCoinType: 14,
}

var ViacoinGenesisHash = chainhash.Hash([chainhash.HashSize]byte{ // Make go vet happy.
	0xa1, 0x0b, 0x40, 0x9a, 0x53, 0xa1, 0x1d, 0x97,
	0x79, 0x0d, 0x97, 0x70, 0x5a, 0xe3, 0xab, 0xea,
	0x15, 0x15, 0x33, 0xc0, 0x8e, 0x12, 0x30, 0x98,
	0x04, 0x76, 0x99, 0x1f, 0x00, 0x54, 0x0b, 0x4e,
})

var ViacoinMerkleRoot = chainhash.Hash([chainhash.HashSize]byte{ // Make go vet happy.
	0x00, 0x52, 0x27, 0x53, 0x00, 0x29, 0x39, 0xc7,
	0x86, 0x59, 0xb4, 0xfd, 0xc6, 0xed, 0x56, 0xc6,
	0xb6, 0xaa, 0xcd, 0xc7, 0x58, 0x6f, 0xac, 0xf2,
	0xf6, 0xad, 0xa2, 0x01, 0x2e, 0xd3, 0x17, 0x03,
})

var ViacoinGenesisBlock = wire.MsgBlock{
	Header: wire.BlockHeader{
		Version:    1,
		PrevBlock:  chainhash.Hash{},
		MerkleRoot: ViacoinMerkleRoot,
		Timestamp:  time.Unix(1405164774, 0),
		Bits:       0x1e01ffff,
		Nonce:      4016033,
	},
}
