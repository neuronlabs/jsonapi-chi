package jsonapichi

type EndpointType int

const (
	Create EndpointType = iota
	Get
	GetRelatedFields
	GetRelationships
	List
	Patch
	PatchRelationships
	Delete
)

// AllEndpoints returns all endpoint types.
func AllEndpoints() []EndpointType {
	return allEndpoints()
}

func allEndpoints() []EndpointType {
	return []EndpointType{
		Create,
		Get,
		GetRelatedFields,
		GetRelationships,
		List,
		Patch,
		PatchRelationships,
		Delete,
	}
}
