package user

import "user_service/proto"

type UserHandler struct {
	proto.UnimplementedUserServer
}
