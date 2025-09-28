package enum

type AssignmentCourierTaskTypeEnum string

const (
	FirstmileAssignmentCourierTaskTypeEnum       = AssignmentCourierTaskTypeEnum("firstmile")
	LastmilePickupAssignmentCourierTaskTypeEnum  = AssignmentCourierTaskTypeEnum("lastmilepickup")
	LastmileDropoffAssignmentCourierTaskTypeEnum = AssignmentCourierTaskTypeEnum("lastmiledropoff")
)
