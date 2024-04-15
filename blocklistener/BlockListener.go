package blocklistener

import (
	"context"
	"fmt"

	"github.com/bitcoin-sv/ubsv/ulogger"
	"github.com/bitcoin-sv/ubsv/util/p2p"

	"github.com/ordishs/gocore"
)

var blockTopicName string

type blocklistener struct {
	logger  ulogger.Logger
	p2pNode *p2p.P2PNode
}

func NewBlockListener() *blocklistener {
	logger := ulogger.New("blocklistener")

	logger.Infof("Creating P2P service")

	p2pIp, ok := gocore.Config().Get("p2p_ip")
	if !ok {
		panic("p2p_ip_prefix not set in config")
	}
	p2pPort, ok := gocore.Config().GetInt("p2p_port")
	if !ok {
		panic("p2p_port not set in config")
	}

	topicPrefix, ok := gocore.Config().Get("p2p_topic_prefix")
	if !ok {
		panic("p2p_topic_prefix not set in config")
	}

	btn, ok := gocore.Config().Get("p2p_block_topic")
	if !ok {
		panic("p2p_block_topic not set in config")
	}

	sharedKey, ok := gocore.Config().Get("p2p_shared_key")
	if !ok {
		panic(fmt.Errorf("error getting p2p_shared_key"))
	}

	usePrivateDht := gocore.Config().GetBool("p2p_dht_use_private", false)
	optimiseRetries := gocore.Config().GetBool("p2p_optimise_retries", false)

	blockTopicName = fmt.Sprintf("%s-%s", topicPrefix, btn)

	staticPeers, _ := gocore.Config().GetMulti("p2p_static_peers", "|")
	privateKey, _ := gocore.Config().Get("p2p_private_key")

	config := p2p.P2PConfig{
		ProcessName:     "peer",
		IP:              p2pIp,
		Port:            p2pPort,
		PrivateKey:      privateKey,
		SharedKey:       sharedKey,
		UsePrivateDHT:   usePrivateDht,
		OptimiseRetries: optimiseRetries,
		Advertise:       true,
		StaticPeers:     staticPeers,
	}

	p2pNode := p2p.NewP2PNode(logger, config)

	return &blocklistener{
		logger:  logger,
		p2pNode: p2pNode,
	}

}

func (bl *blocklistener) Start(ctx context.Context) error {
	bl.logger.Infof("Starting block listener")

	if err := bl.p2pNode.Start(ctx, blockTopicName); err != nil {
		return fmt.Errorf("Error starting P2P node: %v", err)
	}

	err := bl.p2pNode.SetTopicHandler(ctx, blockTopicName, func(ctx context.Context, data []byte, peerID string) {
		bl.logger.Infof("Node received data from peer %s: %s", peerID, string(data))
	})
	if err != nil {
		return fmt.Errorf("Error setting topic handler: %v", err)
	}

	<-ctx.Done()

	return nil
}
