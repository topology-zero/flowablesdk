package process

type BpmnModel struct {
	definitionsAttributes map[string][]ExtensionAttribute
}

type ExtensionAttribute struct {
	name, value, namespacePrefix, namespace string
}
