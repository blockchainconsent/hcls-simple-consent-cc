# hcls-simple-consent-cc

### Prerequisites

#### GVM Setup

zsh < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)

echo "source $HOME/.gvm/scripts/gvm" >> ~/.bash_profile && source ~/.bash_profile

gvm install go1.14

gvm use go1.14

#### Clone Simple Consent chaincode

cd $GOPATH

mkdir src

git clone https://github.ibm.com/HCLS-Consent-Manager/hcls-simple-consent-cc

### Run go tests

To run an individual test in verbose mode:

`go test -v -run TestCreateConsent`

To run all tests:

`go test`

### Package Fabric 2 chaincode

Note: At the time of this writing there is a 2.3 version of Fabric, but IBP's latest software is at 2.2.1-5.

1. Download the Fabric Samples using this command:

    `curl -sSL https://bit.ly/2ysbOFE | bash -s -- 2.2.2 1.4.9`

    To get the latest version instead, run this command:

    `curl -sSL https://bit.ly/2ysbOFE | bash -s`

3. Set your PATH and FABRIC_CFG_PATH environment variables:

    `export PATH=<path to current directory>/fabric-samples/bin:${PATH}`
    
    `export FABRIC_CFG_PATH=<path to current directory>/fabric-samples/config`

4. Navigate to the root directory of this repository and run this command (specify the desired chaincode version) to package the chaincode:

    `peer lifecycle chaincode package consent@<version>.tar.gz --label consent_<version> --path .`


Useful Links:
- https://hyperledger-fabric.readthedocs.io/en/v2.2.1/commands/peerlifecycle.html
- Fabric Samples: https://github.com/hyperledger/fabric-samples/
- https://hyperledger-fabric.readthedocs.io/en/release-2.2/install.html