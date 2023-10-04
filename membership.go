package itswizard_m_userprovisioning

import (
	"encoding/xml"
	"errors"
)

type membership struct {
	XMLName   xml.Name `xml:"membership"`
	Sourcedid struct {
		Source string `xml:"source"` // univention
		ID     string `xml:"id"`     // root, schoolname or groupname
	} `xml:"sourcedid"`
	Member []member `xml:"member"`
}

/*
func Membership (institution string, groupID int, teilnehmer []member) (*membership){
	a := new (membership)
	a.Sourcedid.Source = institution
	a.Sourcedid.ID = strconv.Itoa(groupID)
	a.Member = teilnehmer
	return a
}
*/

func Membership(institution string, groupID string) *membership {
	a := new(membership)
	a.Sourcedid.Source = institution
	a.Sourcedid.ID = groupID
	//a.Member = teilnehmer
	return a
}

func (a *membership) AddMemeber(member member) {
	a.Member = append(a.Member, member)
}

func (a *membership) Parse2XML() ([]byte, string, error) {
	x, err := xml.Marshal(a)
	if err != nil {
		err := errors.New(err.Error() + " Membership " + a.Sourcedid.ID + " was not (!!!!) created \n")
		return nil, "", err
	}
	log := " Membership " + a.Sourcedid.ID + " was created \n "
	return x, log, err
}
