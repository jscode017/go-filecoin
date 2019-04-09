package proofs

import (
	"github.com/filecoin-project/go-filecoin/types"
)

// VerifySealRequest represents a request to verify the output of a Seal() operation.
type VerifySealRequest struct {
	CommD      types.CommD     // returned from seal
	CommR      types.CommR     // returned from seal
	CommRStar  types.CommRStar // returned from seal
	Proof      types.SealProof // returned from seal
	ProverID   [31]byte        // uniquely identifies miner
	SectorID   [31]byte        // uniquely identifies sector
	SectorSize types.SectorSize
}

// VerifyPoSTRequest represents a request to generate verify a proof-of-spacetime.
type VerifyPoSTRequest struct {
	ChallengeSeed types.PoStChallengeSeed
	CommRs        []types.CommR
	Faults        []uint64
	Proofs        []types.PoStProof
	SectorSize    types.SectorSize
}

// VerifyPoSTResponse communicates the validity of a provided proof-of-spacetime.
type VerifyPoSTResponse struct {
	IsValid bool
}

// VerifySealResponse communicates the validity of a provided proof-of-replication.
type VerifySealResponse struct {
	IsValid bool
}

// Verifier provides an interface to the proving subsystem.
type Verifier interface {
	VerifyPoST(VerifyPoSTRequest) (VerifyPoSTResponse, error)
	VerifySeal(VerifySealRequest) (VerifySealResponse, error)
}

// Mode configures sealing, sector packing, PoSt generation and other behaviors
// of libfilecoin_proofs. Use Test mode to seal and generate PoSts quickly over
// tiny sectors. Use Live when operating a real Filecoin node.
type Mode int

const (
	// TestMode configures the SectorBuilder to be used with large sectors, in tests.
	TestMode = Mode(iota)
	// LiveMode configures the SectorBuilder to be used by someone operating a real
	// Filecoin node.
	LiveMode
)
