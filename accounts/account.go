package accounts

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"math/big"
	"os"
	"sync"

	"github.com/scroll-tech/go-ethereum/accounts/abi/bind"
	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/core/types"
	"github.com/scroll-tech/go-ethereum/crypto"
	"github.com/scroll-tech/go-ethereum/ethclient"
	"github.com/scroll-tech/go-ethereum/log"

	"tool/utils"
)

var (
	minBalance = new(big.Int).SetBytes(common.FromHex("152d02c7e14af6000000"))
	privs      = []string{
		"7ecb571e8858355e10d1e73e8446366fbb1b13750a8e22f33f9319f3bbac5919",
		"3849cc3ddb917286e226f9ded4ff6fa8694c91f44013e865824bd7fea293d293",
		"2552e6aa58c31ff7bbb1afd1e3803a612acca87ef88b9d0b49f8d21a29a10c67",
		"37e15d57a506fa5436e461fd5786d0c3d56624d048709f35151e5bc56bd9efc0",
		"f9c3284f49bd22029e3407f8a76326913dd7d0106b2124b987badaa3b91cb9ca",
		"4a1b05c6a5dfcf0b8e168ce843bb28c3a8702de13ad0fbeda18cde67c0636280",
		"c1b01806683199188b46d73da4dca723c987e5394b427a5ee55eef0a189f97dd",
		"b92b5c57b48a1cf1269b6b0ed17a3db90470be8565dd32067e79d84c4e874bb7",
		"e110409d0a8bf83975d186775833d2a3996455eab91998e07b44d93100fcbc1a",
		"aac114801a23e0b398e43ad96e8b949fefe974eed31c713297d7dbf9f5bc7809",
		"a04a2ac7fb2cc8a4b0f5fc605a6475bca146f6bfd7738e6a34248fc995c33b24",
		"5af445dc34b8de6cab0d0b342eb675e924febdfb2b54f14107b9bed124c540e9",
		"1bc29fb427301c172e80739bdae296c77342d7d6c1ae69338eb6bffb53a96a45",
		"4b108b32048cf3507e074f064cdf1c3058180f1b74bc2344d65b47c4a84f37be",
		"c8fc6185ebf214f07d0bc15cedf802507f64c8c55b3c0860e43bc956244c3bdb",
		"7b2d0234994da228cca75b7affbf608421fe5de87891fd6fd80b51d7653581ba",
		"26e74d41ee5c1cbd357bd957aaeb9df46d429a05d72bdbc6765139d9ccb48803",
		"4c33d4249d36a1ed3e39b76cac8898a1ced2b647274d8df2a0884225fd1c0ac4",
		"3b0a60fd2d6f7110ccaeddc60089a01a0a10b7e355553c5886962c3ff06d86f7",
		"51807028941be9360c281a83dd7abfa426f12a5130b6f2e9b86c34733453179d",
		"ccd549c6a9f540bba86e516bdc24b5c26f799b600518ee9fde82882c0a69aa59",
		"a8472bf30b649acd7df45b1f2854b7bffb7532c8c88eb97242addc67af4eefb5",
		"f8608a4f8f4e9e5dca4e67b56ae3db8a872b02c3076e604b2a2cb78a47f2368d",
		"80d6dcb58097701f9870c26f32d029ee648d38aa8374debb05c368fd1a93f224",
		"68aeb40245b2c7bee41c5b03f170ef980ad0045004e47e1341fd4f283c39e15d",
		"b6509a20df42c8f451580c794f9a45368720d71dec5474ede361d308473616f7",
		"2d7a912565592f85f2e278c3a5038a45ed1923e43361fcc4404ea15ff69d92c3",
		"f03f942fc7aedd28626d0cfccc74526bae4d9e6c708287d8842e1a999c51918c",
		"7a0906b4edf843db61d3c36bf90cda215e975e9ec7ddcd07bad0b7b64f500fcd",
		"5a4fbae45f7872522ced3dd60961fb05db980fc6abf3287b6e08b949e0546092",
		"cb11252aaada57dd172b3f5e9850b3a7ccfc93e3f950f757a4b4669371b2497e",
		"a120a23552df789bec9ba5b38f212bc90970249c280bdd32ac87de3b998c2a52",
		"78b20d8bf75868560b3e53648c3672f0b8c270bf17f28a130846194ba1188eac",
		"0141509614186d89a5c253f9510473206f81a743131eb4b1bc1db53b91fe1c20",
		"b73e2578ea748654299dcd2a28bc6e738ea162699d8c3a8b3827f1200601308a",
		"3b58c810a186aed93d373c2a1c9a1d8e1b568c81fa277c5ef7fe7ae1f01d2bc8",
		"497af9b488fbf8404778b6d2b54f5468960bd9b5e29be92a58bef4bc3d89ee51",
		"3e2fb7b1222bde46d0304a727c4a00ba5cbf54fa0925fdbe0796d252f8a86333",
		"767a86d8f1d51bb703e092390416fac87a82631576f3afb1cb835224779e1a8a",
		"ee8af34c10162db27d8e86c77c5a5c463753f9f7a372bbbf14b2f0d6b36a7ce8",
	}

	once      sync.Once
	AddrPrivs []*struct {
		Address common.Address
		*ecdsa.PrivateKey
	}
)

func init() {
	once.Do(func() {
		for _, privStr := range privs {
			priv, _ := crypto.HexToECDSA(privStr)
			addr := crypto.PubkeyToAddress(priv.PublicKey)
			AddrPrivs = append(AddrPrivs, &struct {
				Address common.Address
				*ecdsa.PrivateKey
			}{Address: addr, PrivateKey: priv})
		}
	})
}

type Accounts struct {
	client *ethclient.Client

	Root     *bind.TransactOpts
	Accounts []*bind.TransactOpts

	accsCh chan *bind.TransactOpts
}

func NewAccounts(ctx context.Context, limit int, client *ethclient.Client, keystore, password string) (*Accounts, error) {
	// Set useful internal accounts.
	if len(AddrPrivs) > limit {
		AddrPrivs = AddrPrivs[:limit]
	}

	data, err := os.ReadFile(keystore)
	if err != nil {
		return nil, err
	}
	// get chainID from l2geth handler
	chainid, err := client.ChainID(ctx)
	if err != nil {
		return nil, err
	}

	auth, err := bind.NewTransactorWithChainID(bytes.NewBuffer(data), password, chainid)
	if err != nil {
		return nil, err
	}
	//setGasAndGaslimit(auth)

	var accounts []*bind.TransactOpts
	for _, acc := range AddrPrivs {
		acc, err := bind.NewKeyedTransactorWithChainID(acc.PrivateKey, chainid)
		if err != nil {
			return nil, err
		}
		//setGasAndGaslimit(acc)
		accounts = append(accounts, acc)
	}

	accs := &Accounts{
		client:   client,
		Root:     auth,
		Accounts: accounts,
		accsCh:   make(chan *bind.TransactOpts, len(privs)*2),
	}
	accs.checkAccounts(ctx)

	return accs, nil
}

func (a *Accounts) GetAccount() *bind.TransactOpts {
	select {
	case acc := <-a.accsCh:
		return acc
	}
}

func (a *Accounts) SetAccount(acc *bind.TransactOpts) {
	select {
	case a.accsCh <- acc:
	default:
		log.Warn("account chan is full")
	}
}

func (a *Accounts) checkAccounts(ctx context.Context) {
	var tx *types.Transaction
	for _, acc := range a.Accounts {
		balence, err := a.client.BalanceAt(ctx, acc.From, nil)
		if err != nil {
			continue
		}
		if balence.Cmp(minBalance) < 0 {
			var (
				value = minBalance
				err   error
			)
			tx, err = CreateSignedTx(a.client, a.Root, &acc.From, value, nil)
			if err != nil {
				log.Error("failed to create tx", "err", err)
				continue
			}
			err = a.client.SendTransaction(ctx, tx)
			if err != nil {
				log.Error("Failed to send balance to account", "err", err)
			} else {
				log.Debug("send balance to account", "account", acc.From.String(), "balance", balence.Add(balence, value).String())
			}
		}

		a.accsCh <- acc
	}
	// wait util it's on block
	if tx != nil {
		utils.WaitPendingTx(ctx, a.client, tx.Hash())
	}
}

func CreateSignedTx(client *ethclient.Client, auth *bind.TransactOpts, to *common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	nonce, err := client.PendingNonceAt(context.Background(), auth.From)
	if err != nil {
		return nil, err
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}
	tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		To:       to,
		Data:     data,
		Value:    value,
		Gas:      500000,
		GasPrice: gasPrice,
	})
	signedTx, err := auth.Signer(auth.From, tx)
	if err != nil {
		return nil, err
	}
	return signedTx, nil
}

// Set gas price and gas limit.
func setGasAndGaslimit(auth *bind.TransactOpts) {
	auth.GasPrice = big.NewInt(1108583800)
	auth.GasLimit = 11529940
}
