package acme

// A Challenge object represents a server's offer to validate a
// client's possession of an identifier in a specific way.  Unlike the
// other objects listed above, there is not a single standard structure
// for a challenge object.  The contents of a challenge object depend on
// the validation method being used.  The general structure of challenge
// objects and an initial set of validation methods are described in [§8]
//
// [§8]: https://www.rfc-editor.org/rfc/rfc8555#section-8
type Challenge any

// ChallengeState indicates the status of a Challenge ([§7.1.6])
//
// [§7.1.6]: https://www.rfc-editor.org/rfc/rfc8555#section-7.1.6
type ChallengeState string

const (
	// ChallengePending is the initial state of a Challenge
	// once it's created
	ChallengePending ChallengeState = "pending"
	// ChallengeProcessing is the state after ChallengePending
	// when the client responds to the challenge ([§7.1.5])
	//
	// [§7.1.5]: https://www.rfc-editor.org/rfc/rfc8555#section-7.1.5
	ChallengeProcessing ChallengeState = "processing"
	// ChallengeValid indicates the validation was successful
	ChallengeValid ChallengeState = "valid"
	// ChallengeInvalid indicates there was an error on the validation
	ChallengeInvalid ChallengeState = "invalid"
)
