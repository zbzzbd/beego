package notify

type EmailNotifier interface {
	GetTemplateName() EmailTemplate
	GetTo() []string
	GetVariables() []map[string]interface{}
}
