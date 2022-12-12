package status

type RepositoryInternalStatus string

const (
	Success                        RepositoryInternalStatus = "OK"
	InternalServerError            RepositoryInternalStatus = "INVALID_CONTENT_TYPE"
	EntityWithSameKeyAlreadyExists RepositoryInternalStatus = "ENTITY_WITH_SAME_KEY_ALREADY_EXISTS"
)
