package main

const (
	ALIAS = "alias"
	ID    = "registration_id"
)

type Audience struct {
	Object   interface{}
	audience map[string][]string
}

func (this *Audience) All() {
	this.Object = "all"
}

func (this *Audience) SetID(ids []string) {
	this.set(ID, ids)
}

func (this *Audience) SetAlias(alias []string) {
	this.set(ALIAS, alias)
}

func (this *Audience) set(key string, v []string) {
	if this.audience == nil {
		this.audience = make(map[string][]string)
		this.Object = this.audience
	}
	this.audience[key] = v
}
