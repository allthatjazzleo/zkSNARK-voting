pragma solidity ^0.8.0;

interface IVerifier {
    struct Proof {
        G1Point a;
        G2Point b;
        G1Point c;
    }

		struct G1Point {
				uint X;
				uint Y;
		}
		// Encoding of field elements is: X[0] * z + X[1]
		struct G2Point {
				uint[2] X;
				uint[2] Y;
		}
    
    function verifyTx(Proof memory, uint[1] memory) external view returns (bool);
}


contract Election{

	mapping(uint=>uint8) public votesReceived;
	mapping(bytes32=>bool) public voted;
	mapping(address=>bool) public verifierAddr;
	event Voted(
			uint indexed _candidateid
	);
	event whitelistedVerifierAddr(address indexed theaddr);
	uint[] public candidateList=[1,2,3];

	constructor(address[] memory _verifierAddr) public {
    for (uint8 i = 0; i < _verifierAddr.length; i++){
			verifierAddr[_verifierAddr[i]] = true;
      emit whitelistedVerifierAddr(_verifierAddr[i]);
		}
	}


	function totalVotesFor(uint candidate) view public returns (uint8){
		require(validCandidate(candidate));
		return votesReceived[candidate];
	}

	function voteForCandidate(address _verifierAddr, uint candidate, IVerifier.Proof memory proof) public {
		uint[1] memory input = [uint(1)];
		bytes32 hash = keccak256(abi.encodePacked(proof.a.X, proof.a.Y, proof.b.X, proof.b.Y, proof.c.X, proof.c.Y));
		require(voted[hash] == false);
		require(verifierAddr[_verifierAddr] == true);
		require(IVerifier(_verifierAddr).verifyTx(proof, input) == true);
		require(validCandidate(candidate));
		votesReceived[candidate] +=	1;
		voted[hash] = true;
		emit Voted(candidate);
	}

	function validCandidate(uint candidate) view public returns (bool){
		for(uint i = 0; i<candidateList.length; i++){
			if(candidateList[i] == candidate){
				return true;
			}

		}
		return false;
	}
}
