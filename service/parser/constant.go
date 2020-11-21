package parser

const (
	defaultAssignmentKeyword    = "is"
	defaultCreditsKeyword       = "Credits"
	defaultQueryTransnumKeyword = "how much"
	defaultQueryPriceKeyword    = "how many " + defaultCreditsKeyword
	defaultQuestionMarkKeyword  = "?"
	defaultUnexpectedInput      = "I have no idea what you are talking about"
)

func space(s string) string {
	return " " + s + " "
}
