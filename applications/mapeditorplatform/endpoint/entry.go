package endpoint

type (
	EntrySet struct {
		*UserEndPoint
	}
)

func NewEntrySet(userEndPoint *UserEndPoint) *EntrySet {
	return &EntrySet{
		userEndPoint,
	}
}
