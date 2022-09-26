package args

var builder = &holderBuilder{holder: Holder}

type holderBuilder struct {
	holder *holder
}

// GetHolderBuilder returns singleton instance of argument holder builder.
func GetHolderBuilder() *holderBuilder {
	return builder
}
