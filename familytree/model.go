package familytree

//Member is the member of the tree
type Member struct {
	Name      string
	Gender    string
	Partner   *Member
	Mother    *Member
	Father    *Member
	Childrens *Member
	Siblings  *Member
}

//MemberList stores references of all the members
type MemberList map[string]*Member

//FamilyMemberFilePath is the path for the file input
const FamilyMemberFilePath string = "./resources/family_members.txt"
