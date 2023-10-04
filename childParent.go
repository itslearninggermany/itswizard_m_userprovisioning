package itswizard_m_userprovisioning

type childParent struct {
	Relationshiptype string `xml:"relationshiptype"`
	Sourcedid        struct {
		Source string `xml:"source"`
		ID     string `xml:"id"`
	} `xml:"sourcedid"`
}

// Creates a new Child.
func Child(institution, childOrParentID string, parent bool) *childParent {
	a := new(childParent)
	if parent {
		a.Relationshiptype = "parent"
	} else {
		a.Relationshiptype = "child"
	}
	a.Sourcedid.Source = institution
	a.Sourcedid.ID = childOrParentID
	return a
}

// GetChild gibt die Institution und die ChildID zurück
func (a *childParent) GetChildOrParent() (institution, childoParentID string, parent bool) {
	if a.Relationshiptype == "parent" {
		parent = true
	} else {
		parent = false
	}
	return a.Sourcedid.Source, a.Sourcedid.ID, parent
}

// MakeAChildSlice erstellt ein Array mit Childs.
func MakeAChildSlice(institution string, childParentIDs []string, parent bool) (res []childParent) {
	for i := 0; i < len(childParentIDs); i++ {
		a := Child(institution, childParentIDs[0], parent)
		res = append(res, *a)
	}
	return
}
