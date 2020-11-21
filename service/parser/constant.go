package parser

const (
	defaultAssignmentKeyword = "is"
	defaultCreditsKeyword    = "Credits"
	defaultUnexpectedInput   = "I have no idea what you are talking about"
)

func space(s string) string {
	return " " + s + " "
}
