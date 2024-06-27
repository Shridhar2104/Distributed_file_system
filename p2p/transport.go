package p2p
//represents remote node
type Peer interface{

}
//handles communication
//between nodes in the network
//of the form UCP TCP
type Transport interface{

	ListenAndAccept() error

}