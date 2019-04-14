# fabric-pbnetwork

This repository contains my project in **Project Laboratory** subject.

It is a simple Hyperledger Fabric network implementation featuring:
* 3 organizations with 1 peer each
* own couchdb Docker container to serve as database server for each organization
* chaincode with private data (written in Go)
* own Certificate Authority for each organization
* TLS communication between CA servers and clients
* TLS communication between peers and orderer
* Identity Mixer enrollment for peers and users

## Getting started

Prerequisites:
* [Hyperledger Fabric prerequisites](https://hyperledger-fabric.readthedocs.io/en/release-1.4/prereqs.html)
* [Hyperledger Fabric-Ca prerequisites](https://hyperledger-fabric-ca.readthedocs.io/en/release-1.4/users-guide.html#prerequisites)

After cloning:
1. Make sure that the project folder can be found at `$HOME/mycode/onlab/`. In the future, the use of `$PROJECT_HOME` environment variable is going to make setting the project location for each script less painful.
2. Run `setup-scripts/start-ca-docker` to start the CA containers
3. Run `setup-scripts/certificate-setup` to issue certificates
4. Run `setup-scripts/setup-pbnetwork` to generate genesis block, create channel, start peer and orderer containers and create anchor peers
5. Run `pb-network-chaincode-setup` (from CLI container) to join the peers to the channel, and install and instantiate chaincode on all peers
6. Run `setup-scripts/destroy-pbnetwork` to stop all containers and clear every generated data (certs, MSPs, default config files, channel artifacts, etc.)

## Invoking and querying chaincode

Only members of Org1 and Org2 are allowed to read private data, but reading public data and writing onto the ledger is open for all of the organizations.

* Create a pen instance: `create-pen peerIdentity penID model color length width lineWidth buyingPrice sellingPrice`
* Query pen public data by ID: `query-public peerIdentity penID`
* Query pen private data by ID: `query-private peerIdentity penID`

`peerIdentity` is expected in the form of peerN.orgM
`buyingPrice` and `sellingPrice` are considered to be sensitive, and also private data. After the creation of 5 more blocks, they get purged from the ledger.

To get help with the use of the commands, type them without arguments.
