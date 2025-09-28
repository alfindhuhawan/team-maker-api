package domerror

const (
	FailUnmarshalResponseBodyError ErrorType = "ER1000 Fail to unmarshal response body"   // used by controller
	FailUnmarshalRequestBodyError  ErrorType = "ER1000 Fail to unmarshal request body"    // used by controller
	FailUnmarshalTokenError        ErrorType = "ER1000 Fail to unmarshal token into user" // used by controller
	AnyInvalidParameters           ErrorType = "ER1000 there are some invalid parameters" // used by controller
	ApiKeyIsNeeded                 ErrorType = "ER1000 Api Key is needed on the header"   // used by controller
	EntityNotFound                 ErrorType = "ER1001 Entity %s with id %s is not found" // used by injected repo in interactor
	UnrecognizedEnum               ErrorType = "ER1002 %s is not recognized %s enum"      // used by enum
	DatabaseNotFoundInContextError ErrorType = "ER1003 Database is not found in context"  // used by repoimpl
)

const (
	ZipcodeOriginMustNotEmptyOnGroupDaysValidation         ErrorType = "ER1000 zipcode origin must not empty on group days validation"
	ZipcodeDestinationMustNotEmptyOnGroupDaysValidation    ErrorType = "ER1000 zipcode destination must not empty on group days validation"
	ServiceTypeMustNotEmptyOnGroupDaysValidation           ErrorType = "ER1000 service type must not empty on group days validation"
	EstimatedPickupDateMustNotEmptyOnGroupDaysValidation   ErrorType = "ER1000 estimated pickup date must not empty on group days validation"
	EstimatedDeliveryDateMustNotEmptyOnGroupDaysValidation ErrorType = "ER1000 estimated delivery date must not empty on group days validation"
)
