package user

type PreferenceEdit struct {
}
type Preference struct {
	Pairs []PreferencePair `json:"pairs"`
}
type PreferencePair struct {
	Value string `json:"value"`
	Id    uint   `json:"id"`
}
