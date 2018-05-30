// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Ethereum network.
var MainnetBootnodes = []string{
	// Ethereum Foundation Go Bootnodes
	"enode://a979fb575495b8d6db44f750317d0f4622bf4c2aa3365d6af7c284339968eef29b69ad0dce72a4d8db5ebb4968de0e3bec910127f134779fbcb0cb6d3331163c@52.16.188.185:30303", // IE
	"enode://3f1d12044546b76342d59d4a05532c14b85aa669704bfe1f864fe079415aa2c02d743e03218e57a33fb94523adb54032871a6c51b2cc5514cb7c7e35b3ed0a99@13.93.211.84:30303",  // US-WEST
	"enode://78de8a0916848093c73790ead81d1928bec737d565119932b98c6b100d944b7a95e94f847f689fc723399d2e31129d182f7ef3863f2b4c820abbf3ab2722344d@191.235.84.50:30303", // BR
	"enode://158f8aab45f6d19c6cbf4a089c2670541a8da11978a2f90dbf6a502a4a3bab80d288afdbeb7ec0ef6d92de563767f3b1ea9e8e334ca711e9f8e2df5a0385e8e6@13.75.154.138:30303", // AU
	"enode://1118980bf48b0a3640bdba04e0fe78b1add18e1cd99bf22d53daac1fd9972ad650df52176e7c7d89d1114cfef2bc23a2959aa54998a46afcf7d91809f0855082@52.74.57.123:30303",  // SG

	// Ethereum Foundation C++ Bootnodes
	"enode://979b7fa28feeb35a4741660a16076f1943202cb72b6af70d327f053e248bab9ba81760f39d0701ef1d8f89cc1fbd2cacba0710a12cd5314d5e0c9021aa3637f9@5.1.83.226:30303", // DE
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Ropsten test network.
var TestnetBootnodes = []string{
	"enode://30b7ab30a01c124a6cceca36863ece12c4f5fa68e3ba9b0b51407ccc002eeed3b3102d20a88f1c1d3c3154e2449317b8ef95090e77b312d5cc39354f86d5d606@52.176.7.10:30303",    // US-Azure geth
	"enode://865a63255b3bb68023b6bffd5095118fcc13e79dcf014fe4e47e065c350c7cc72af2e53eff895f11ba1bbb6a2b33271c1116ee870f266618eadfc2e78aa7349c@52.176.100.77:30303",  // US-Azure parity
	"enode://6332792c4a00e3e4ee0926ed89e0d27ef985424d97b6a45bf0f23e51f0dcb5e66b875777506458aea7af6f9e4ffb69f43f3778ee73c81ed9d34c51c4b16b0b0f@52.232.243.152:30303", // Parity
	"enode://94c15d1b9e2fe7ce56e458b9a3b672ef11894ddedd0c6f247e0f1d3487f52b66208fb4aeb8179fce6e3a749ea93ed147c37976d67af557508d199d9594c35f09@192.81.208.223:30303", // @gpip
}

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network.
var RinkebyBootnodes = []string{
	"enode://a24ac7c5484ef4ed0c5eb2d36620ba4e4aa13b8c84684e1b4aab0cebea2ae45cb4d375b77eab56516d34bfbd3c1a833fc51296ff084b770b94fb9028c4d25ccf@52.169.42.101:30303", // IE
	"enode://343149e4feefa15d882d9fe4ac7d88f885bd05ebb735e547f12e12080a9fa07c8014ca6fd7f373123488102fe5e34111f8509cf0b7de3f5b44339c9f25e87cb8@52.3.158.184:30303",  // INFURA
	"enode://b6b28890b006743680c52e64e0d16db57f28124885595fa03a562be1d2bf0f3a1da297d56b13da25fb992888fd556d4c1a27b1f39d531bde7de1921c90061cc6@159.89.28.211:30303", // AKASHA
}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{
//"enode://ecaebb39dafa9c1d811c4a2544b21a1bf604325f74c568f36ac7a9a5bb91c1b74d3378b059169fb56586735ff8b03533ac1748bec1edd07ea278bb2c40565e4f@52.161.13.240:30303",
}

var CallistoMainnetBootnodes = []string{
	"enode://ab0d9fe81f23653c3217303d3a4bc035be908b05f8b12d6f155ab4598d7760cfea1e009e414771279c9ffc3499950283afaabe099a5329c5a6013f57d77de0a6@104.236.96.118:30303",
	"enode://9e2d9dc2639e02893aa17c80e6ba8e8803fd3166083b622a841fc852161112720281514436f7605c89041d5efa1738215185c4c4024ff812b0f500c403cc0ab1@206.189.47.198:30303",
	"enode://149ba679e8851c3e0d030e0dc0336984b97c83ef649e68ec113dcf266449364cc1ec8ee27950f71b00c2182ef504894fa7bff19f6741978ced67e9e4b6536d2a@206.189.45.31:30303",
	"enode://2038f895f0870f76456a879b97a276135418295ff9f6e24b6de11225360adc87d3aa6d90302aafb36a4cd9542eaf0f160ae48439da938d7e3031d838ef11b8e8@13.125.216.165:30303", // clopool.net
	"enode://fa62ec65fb5cd2f7e917f2424e44407e85f3683c3439ba2401358159a2ffb0bd6c2c25393965145672da2545a6fec2c6c2c3128fcaf332a82b897b9bac7164e1@13.125.89.37:30303",   // clopool.net
	"enode://094181735965f03a05379f8c9ef258cd8d119e13ff1bfe9aa2a7e113473b76043f8ea069fc34296643d257527a7ebc632c17b27da942a1f1e34a956b2820b9ef@52.79.151.194:30303",  // clopool.net
	"enode://85031e6e5fcb30ab5ed288463445de95260a4078a3fe7facbf35fe472ca02b70d0b4dad64697d73feedab984ff2737bb1bfe520b289c88c5960a1e3843ef5491@13.209.18.37:30303",   // clopool.net
	"enode://9fedfa9fddb8f2edda8ba25f3fa74b0a66c8393e90a3c8dae8fd41e417fddd9229611cac379467aa048e00937b55fda86d4fe841d41216c08a1d1965a77306a8@13.125.219.230:30303", // clopool.net
	"enode://f5e7f848533fc7e0bccb5c81ecf255c04bcc6af3ea802e7304247faaaf39b39c4ff38f16aef13fa11952257fb399ac7027fbb0d97156954edd03755a548bf12e@109.110.37.200:30303", // farmine.net
}

// CallistoTestnetBootnodes are the enode URLs of the P2P bootstrap nodes for the
// Callisto test network.
var CallistoTestnetBootnodes = []string{
	"enode://1e1d25d9eafd92319ee85f502c2f5274a62029b3f9ec003cba259a600ae0def256a09ec505a569aabc4a447d2fd4df9c682243091730f83556693dab17206bd0@167.99.96.186:30303",
}
