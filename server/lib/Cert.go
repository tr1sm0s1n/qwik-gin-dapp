// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lib

import (
	"bytes"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = bytes.Equal
	_ = errors.New
	_ = big.NewInt
	_ = common.Big1
	_ = types.BloomLookup
	_ = abi.ConvertType
)

// CertMetaData contains all meta data concerning the Cert contract.
var CertMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"course\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"grade\",\"type\":\"string\"}],\"name\":\"Issued\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Certificates\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"course\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"grade\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"date\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_course\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_grade\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_date\",\"type\":\"string\"}],\"name\":\"issue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	ID:  "Cert",
	Bin: "0x6080604052348015600e575f5ffd5b50335f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550610b828061005b5f395ff3fe608060405234801561000f575f5ffd5b5060043610610034575f3560e01c80639622c83614610038578063d1e7be261461006b575b5f5ffd5b610052600480360381019061004d919061047e565b610087565b6040516100629493929190610519565b60405180910390f35b610085600480360381019061008091906106a4565b6102cb565b005b6001602052805f5260405f205f91509050805f0180546100a6906107b8565b80601f01602080910402602001604051908101604052809291908181526020018280546100d2906107b8565b801561011d5780601f106100f45761010080835404028352916020019161011d565b820191905f5260205f20905b81548152906001019060200180831161010057829003601f168201915b505050505090806001018054610132906107b8565b80601f016020809104026020016040519081016040528092919081815260200182805461015e906107b8565b80156101a95780601f10610180576101008083540402835291602001916101a9565b820191905f5260205f20905b81548152906001019060200180831161018c57829003601f168201915b5050505050908060020180546101be906107b8565b80601f01602080910402602001604051908101604052809291908181526020018280546101ea906107b8565b80156102355780601f1061020c57610100808354040283529160200191610235565b820191905f5260205f20905b81548152906001019060200180831161021857829003601f168201915b50505050509080600301805461024a906107b8565b80601f0160208091040260200160405190810160405280929190818152602001828054610276906107b8565b80156102c15780601f10610298576101008083540402835291602001916102c1565b820191905f5260205f20905b8154815290600101906020018083116102a457829003601f168201915b5050505050905084565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610359576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161035090610832565b60405180910390fd5b60405180608001604052808581526020018481526020018381526020018281525060015f8781526020019081526020015f205f820151815f01908161039e91906109f0565b5060208201518160010190816103b491906109f0565b5060408201518160020190816103ca91906109f0565b5060608201518160030190816103e091906109f0565b50905050826040516103f29190610af9565b60405180910390207fc7fc4858662996e8fe091a9941384ab745bfc51199d7c0a5e45a791ff945993c868460405161042b929190610b1e565b60405180910390a25050505050565b5f604051905090565b5f5ffd5b5f5ffd5b5f819050919050565b61045d8161044b565b8114610467575f5ffd5b50565b5f8135905061047881610454565b92915050565b5f6020828403121561049357610492610443565b5b5f6104a08482850161046a565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f6104eb826104a9565b6104f581856104b3565b93506105058185602086016104c3565b61050e816104d1565b840191505092915050565b5f6080820190508181035f83015261053181876104e1565b9050818103602083015261054581866104e1565b9050818103604083015261055981856104e1565b9050818103606083015261056d81846104e1565b905095945050505050565b5f5ffd5b5f5ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6105b6826104d1565b810181811067ffffffffffffffff821117156105d5576105d4610580565b5b80604052505050565b5f6105e761043a565b90506105f382826105ad565b919050565b5f67ffffffffffffffff82111561061257610611610580565b5b61061b826104d1565b9050602081019050919050565b828183375f83830152505050565b5f610648610643846105f8565b6105de565b9050828152602081018484840111156106645761066361057c565b5b61066f848285610628565b509392505050565b5f82601f83011261068b5761068a610578565b5b813561069b848260208601610636565b91505092915050565b5f5f5f5f5f60a086880312156106bd576106bc610443565b5b5f6106ca8882890161046a565b955050602086013567ffffffffffffffff8111156106eb576106ea610447565b5b6106f788828901610677565b945050604086013567ffffffffffffffff81111561071857610717610447565b5b61072488828901610677565b935050606086013567ffffffffffffffff81111561074557610744610447565b5b61075188828901610677565b925050608086013567ffffffffffffffff81111561077257610771610447565b5b61077e88828901610677565b9150509295509295909350565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806107cf57607f821691505b6020821081036107e2576107e161078b565b5b50919050565b7f4163636573732044656e696564000000000000000000000000000000000000005f82015250565b5f61081c600d836104b3565b9150610827826107e8565b602082019050919050565b5f6020820190508181035f83015261084981610810565b9050919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026108ac7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610871565b6108b68683610871565b95508019841693508086168417925050509392505050565b5f819050919050565b5f6108f16108ec6108e78461044b565b6108ce565b61044b565b9050919050565b5f819050919050565b61090a836108d7565b61091e610916826108f8565b84845461087d565b825550505050565b5f5f905090565b610935610926565b610940818484610901565b505050565b5b81811015610963576109585f8261092d565b600181019050610946565b5050565b601f8211156109a85761097981610850565b61098284610862565b81016020851015610991578190505b6109a561099d85610862565b830182610945565b50505b505050565b5f82821c905092915050565b5f6109c85f19846008026109ad565b1980831691505092915050565b5f6109e083836109b9565b9150826002028217905092915050565b6109f9826104a9565b67ffffffffffffffff811115610a1257610a11610580565b5b610a1c82546107b8565b610a27828285610967565b5f60209050601f831160018114610a58575f8415610a46578287015190505b610a5085826109d5565b865550610ab7565b601f198416610a6686610850565b5f5b82811015610a8d57848901518255600182019150602085019450602081019050610a68565b86831015610aaa5784890151610aa6601f8916826109b9565b8355505b6001600288020188555050505b505050505050565b5f81905092915050565b5f610ad3826104a9565b610add8185610abf565b9350610aed8185602086016104c3565b80840191505092915050565b5f610b048284610ac9565b915081905092915050565b610b188161044b565b82525050565b5f604082019050610b315f830185610b0f565b8181036020830152610b4381846104e1565b9050939250505056fea2646970667358221220377a44915d0bd1f5410614bbe711f138a1806ae26c77a47d074a19511f51277164736f6c634300081d0033",
}

// Cert is an auto generated Go binding around an Ethereum contract.
type Cert struct {
	abi abi.ABI
}

// NewCert creates a new instance of Cert.
func NewCert() *Cert {
	parsed, err := CertMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &Cert{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *Cert) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackCertificates is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9622c836.
//
// Solidity: function Certificates(uint256 ) view returns(string name, string course, string grade, string date)
func (cert *Cert) PackCertificates(arg0 *big.Int) []byte {
	enc, err := cert.abi.Pack("Certificates", arg0)
	if err != nil {
		panic(err)
	}
	return enc
}

// CertificatesOutput serves as a container for the return parameters of contract
// method Certificates.
type CertificatesOutput struct {
	Name   string
	Course string
	Grade  string
	Date   string
}

// UnpackCertificates is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9622c836.
//
// Solidity: function Certificates(uint256 ) view returns(string name, string course, string grade, string date)
func (cert *Cert) UnpackCertificates(data []byte) (CertificatesOutput, error) {
	out, err := cert.abi.Unpack("Certificates", data)
	outstruct := new(CertificatesOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Course = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Grade = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Date = *abi.ConvertType(out[3], new(string)).(*string)
	return *outstruct, err

}

// PackIssue is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd1e7be26.
//
// Solidity: function issue(uint256 _id, string _name, string _course, string _grade, string _date) returns()
func (cert *Cert) PackIssue(id *big.Int, name string, course string, grade string, date string) []byte {
	enc, err := cert.abi.Pack("issue", id, name, course, grade, date)
	if err != nil {
		panic(err)
	}
	return enc
}

// CertIssued represents a Issued event raised by the Cert contract.
type CertIssued struct {
	Course common.Hash
	Id     *big.Int
	Grade  string
	Raw    *types.Log // Blockchain specific contextual infos
}

const CertIssuedEventName = "Issued"

// ContractEventName returns the user-defined event name.
func (CertIssued) ContractEventName() string {
	return CertIssuedEventName
}

// UnpackIssuedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Issued(string indexed course, uint256 id, string grade)
func (cert *Cert) UnpackIssuedEvent(log *types.Log) (*CertIssued, error) {
	event := "Issued"
	if log.Topics[0] != cert.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(CertIssued)
	if len(log.Data) > 0 {
		if err := cert.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range cert.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}
