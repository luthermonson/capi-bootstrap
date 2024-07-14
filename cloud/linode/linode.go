package linode

import "capi-bootstrap/cloud"

type Linode struct {
	InstanceID     string
	NodeBalancerID string
}

func NewProvider() cloud.Provider {
	return &Linode{}
}

func (l *Linode) Initialize() error {
	return nil
}

func (l *Linode) PreCreate() error {
	return nil
}

func (l *Linode) Create() error {
	return nil
}

func (l *Linode) PostCreate() error {
	return nil
}

func (l *Linode) PreDelete() error {
	return nil
}

func (l *Linode) Delete() error {
	return nil
}

func (l *Linode) PostDelete() error {
	return nil
}
