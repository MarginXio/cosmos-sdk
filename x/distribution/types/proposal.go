package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

const (
	// ProposalTypeCommunityPoolSpend defines the type for a CommunityPoolSpendProposal
	ProposalTypeCommunityPoolSpend = "CommunityPoolSpend"
	ProposalTypeFeeSpendProposal = "FeeSpendProposal"
)

// Assert CommunityPoolSpendProposal implements govtypes.Content at compile-time
var _ govtypes.Content = &CommunityPoolSpendProposal{}
var _ govtypes.Content = &FeeSpendProposal{}

func init() {
	govtypes.RegisterProposalType(ProposalTypeCommunityPoolSpend)
	govtypes.RegisterProposalTypeCodec(&CommunityPoolSpendProposal{}, "cosmos-sdk/CommunityPoolSpendProposal")
	govtypes.RegisterProposalType(ProposalTypeFeeSpendProposal)
	govtypes.RegisterProposalTypeCodec(&FeeSpendProposal{}, "cosmos-sdk/FeeSpendProposal")
}

// NewCommunityPoolSpendProposal creates a new community pool spned proposal.
//
//nolint:interfacer
func NewCommunityPoolSpendProposal(title, description string, recipient sdk.AccAddress, amount sdk.Coins) *CommunityPoolSpendProposal {
	return &CommunityPoolSpendProposal{title, description, recipient.String(), amount}
}

// GetTitle returns the title of a community pool spend proposal.
func (csp *CommunityPoolSpendProposal) GetTitle() string { return csp.Title }

// GetDescription returns the description of a community pool spend proposal.
func (csp *CommunityPoolSpendProposal) GetDescription() string { return csp.Description }

// GetDescription returns the routing key of a community pool spend proposal.
func (csp *CommunityPoolSpendProposal) ProposalRoute() string { return RouterKey }

// ProposalType returns the type of a community pool spend proposal.
func (csp *CommunityPoolSpendProposal) ProposalType() string { return ProposalTypeCommunityPoolSpend }

// ValidateBasic runs basic stateless validity checks
func (csp *CommunityPoolSpendProposal) ValidateBasic() error {
	err := govtypes.ValidateAbstract(csp)
	if err != nil {
		return err
	}
	if !csp.Amount.IsValid() {
		return ErrInvalidProposalAmount
	}
	if csp.Recipient == "" {
		return ErrEmptyProposalRecipient
	}

	return nil
}

// String implements the Stringer interface.
func (csp CommunityPoolSpendProposal) String() string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf(`Community Pool Spend Proposal:
  Title:       %s
  Description: %s
  Recipient:   %s
  Amount:      %s
`, csp.Title, csp.Description, csp.Recipient, csp.Amount))
	return b.String()
}


func NewFeeSpendProposal(title, description string, recipient sdk.AccAddress, amount sdk.Coins) *FeeSpendProposal {
	return &FeeSpendProposal{title, description, recipient.String(), amount}
}

// GetTitle returns the title of a community pool spend proposal.
func (csp *FeeSpendProposal) GetTitle() string { return csp.Title }

// GetDescription returns the description of a community pool spend proposal.
func (csp *FeeSpendProposal) GetDescription() string { return csp.Description }

// GetDescription returns the routing key of a community pool spend proposal.
func (csp *FeeSpendProposal) ProposalRoute() string { return RouterKey }

// ProposalType returns the type of a community pool spend proposal.
func (csp *FeeSpendProposal) ProposalType() string { return ProposalTypeCommunityPoolSpend }

// ValidateBasic runs basic stateless validity checks
func (csp *FeeSpendProposal) ValidateBasic() error {
	err := govtypes.ValidateAbstract(csp)
	if err != nil {
		return err
	}
	if !csp.Amount.IsValid() {
		return ErrInvalidProposalAmount
	}
	if csp.Recipient == "" {
		return ErrEmptyProposalRecipient
	}

	return nil
}

// String implements the Stringer interface.
func (csp FeeSpendProposal) String() string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf(`Community Pool Spend Proposal:
  Title:       %s
  Description: %s
  Recipient:   %s
  Amount:      %s
`, csp.Title, csp.Description, csp.Recipient, csp.Amount))
	return b.String()
}