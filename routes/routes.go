package routes

import proto "auth/pb"

type ErrorWithRoute struct {
	ErrorDescription string
}

func (r *ErrorWithRoute) Error() string {
	return r.ErrorDescription
}

type AuthRoutes struct {
	proto.UnimplementedAuthRoutesServer
}
