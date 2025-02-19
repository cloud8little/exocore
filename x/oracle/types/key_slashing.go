package types

var (
	SlashingPrefix            = []byte("Slashing/")
	ValidatorReportInfoPrefix = append(SlashingPrefix, []byte("validator/value/")...)
	MissedBitArrayPrefix      = append(SlashingPrefix, []byte("missed/value/")...)
)

func SlashingValidatorReportInfoKey(validator string) []byte {
	return append(ValidatorReportInfoPrefix, []byte(validator)...)
}

func SlashingMissedBitArrayPrefix() []byte {
	return MissedBitArrayPrefix
}

func SlashingMissedBitArrayValidatorPrefix(validator string) []byte {
	key := append([]byte(validator), DelimiterForCombinedKey)
	return append(MissedBitArrayPrefix, key...)
}

func SlashingMissedBitArrayKey(validator string, index uint64) []byte {
	return append(SlashingMissedBitArrayValidatorPrefix(validator), Uint64Bytes(index)...)
}
