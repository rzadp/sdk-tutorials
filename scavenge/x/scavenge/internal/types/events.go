package types

// scavenge module event types
const (
	EventTypeCreateScavenge = "CreateScavenge"
	EventTypeCommitSolution = "CommitSolution"
	EventTypeSolveScavenge  = "SolveScavenge"

	AttributeDescription           = "description"
	AttributeSolution              = "solution"
	AttributeSolutionHash          = "solutionHash"
	AttributeCoinReward            = "coinReward"
	AttributeNFTRewardDenom        = "nftRewardDenom"
	AttributeNFTRewardID           = "nftRewardID"
	AttributeScavenger             = "scavenger"
	AttributeSolutionScavengerHash = "solutionScavengerHash"

	AttributeValueCategory = ModuleName
)
