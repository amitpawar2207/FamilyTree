package familytree

import "strings"

//addPartner adds the member as partner
func addPartner(partnerName, name, gender string, members MemberList) (MemberList, string) {
	partner := members.findNode(partnerName)
	if partner == nil {
		return members, "PERSON_NOT_FOUND"
	}
	if (partner.Gender == strings.ToLower(gender)) || (partner.Partner != nil) {
		return members, "PARTNER_ADDITION_FAILED"
	}

	newMember := Member{}
	newMember.Name = name
	newMember.Gender = strings.ToLower(gender)
	newMember.Partner = partner
	partner.Partner = &newMember

	if strings.ToLower(gender) == "male" {
		if partner.Childrens != nil {
			newMember.Childrens = partner.Childrens
		}
	}

	members[strings.ToLower(name)] = &newMember

	return members, "PARTNER_ADDITION_SUCCEEDED"
}

func (familyMembers MemberList) addChildren(parentName, childName, gender string) string {
	node := familyMembers.findNode(parentName)
	if node == nil {
		return "PERSON_NOT_FOUND"
	}

	node1 := familyMembers.findNode(childName)
	if node1 != nil {
		return "CHILD_ADDITION_FAILED"
	}

	if node.Gender == "female" {

		var newChild Member
		newChild.Name = childName
		newChild.Gender = strings.ToLower(gender)
		newChild.Mother = node

		if node.Partner != nil {
			newChild.Father = node.Partner
		}

		if node.Childrens == nil {
			node.Childrens = &newChild

			if node.Partner != nil {
				node.Partner.Childrens = &newChild
			}
		} else {
			siblingNode := node.Childrens
			for {
				if siblingNode.Siblings == nil {
					siblingNode.Siblings = &newChild
					break
				}
				siblingNode = siblingNode.returnSiblings()
			}
		}

		familyMembers[strings.ToLower(childName)] = &newChild

		return "CHILD_ADDITION_SUCCEEDED"
	}
	return "CHILD_ADDITION_FAILED"
}

func (node *Member) getRelationship(name, relationship string) string {
	result, ans := node.findRelationship(name, relationship)

	if len(result) == 0 && ans != "" {
		return ans
	} else if len(result) == 0 && ans == "" {
		return "NONE"
	}
	return strings.Join(result, " ")
}
