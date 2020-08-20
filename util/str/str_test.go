package str

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCamelToSnake(t *testing.T) {
	assert.Equal(t, "member", CamelToSnake("Member"))
	assert.Equal(t, "member", CamelToSnake("member"))

	assert.Equal(t, "member_friend", CamelToSnake("MemberFriend"))
	assert.Equal(t, "member_friend", CamelToSnake("memberFriend"))

	assert.Equal(t, "member_friend", CamelToSnake("MemberFriend"))

}

func TestSnakeToCamel(t *testing.T) {
	assert.Equal(t, "Member", SnakeToCamel("Member"))
	assert.Equal(t, "member", SnakeToCamel("member"))

	assert.Equal(t, "memberFriend", SnakeToCamel("member_friend"))
	assert.Equal(t, "memberFriend", SnakeToCamel("member_Friend"))
	assert.Equal(t, "MemberFriend", SnakeToCamel("Member_friend"))
	assert.Equal(t, "MemberFriend", SnakeToCamel("Member_Friend"))

	assert.Equal(t, "memberFriend", SnakeToCamel("_member_friend"))
	assert.Equal(t, "memberFriend", SnakeToCamel("_member_Friend"))
	assert.Equal(t, "MemberFriend", SnakeToCamel("_Member_friend"))
	assert.Equal(t, "MemberFriend", SnakeToCamel("_Member_Friend"))

	assert.Equal(t, "memberFriend", SnakeToCamel("member_friend_"))
	assert.Equal(t, "memberFriend", SnakeToCamel("member_Friend_"))
	assert.Equal(t, "MemberFriend", SnakeToCamel("Member_friend_"))
	assert.Equal(t, "MemberFriend", SnakeToCamel("Member_Friend_"))

	assert.Equal(t, "memberFriend", SnakeToCamel("_member_friend_"))
	assert.Equal(t, "memberFriend", SnakeToCamel("_member_Friend_"))
	assert.Equal(t, "MemberFriend", SnakeToCamel("_Member_friend_"))
	assert.Equal(t, "MemberFriend", SnakeToCamel("_Member_Friend_"))

	assert.Equal(t, "memberFriend", SnakeToCamel("memberFriend"))
	assert.Equal(t, "MemberFriend", SnakeToCamel("MemberFriend"))
}

func TestSnakeToBigCamel(t *testing.T) {
	assert.Equal(t, "Member", SnakeToBigCamel("Member"))
	assert.Equal(t, "Member", SnakeToBigCamel("member"))

	assert.Equal(t, "MemberFriend", SnakeToBigCamel("member_friend"))
	assert.Equal(t, "MemberFriend", SnakeToBigCamel("member_Friend"))
	assert.Equal(t, "MemberFriend", SnakeToBigCamel("Member_friend"))
	assert.Equal(t, "MemberFriend", SnakeToBigCamel("Member_Friend"))

	assert.Equal(t, "MemberFriend", SnakeToBigCamel("member_friend_"))
	assert.Equal(t, "MemberFriend", SnakeToBigCamel("member_Friend_"))
	assert.Equal(t, "MemberFriend", SnakeToBigCamel("Member_friend_"))
	assert.Equal(t, "MemberFriend", SnakeToBigCamel("Member_Friend_"))

	assert.Equal(t, "MemberFriend", SnakeToBigCamel("_member_friend"))
	assert.Equal(t, "MemberFriend", SnakeToBigCamel("_member_Friend"))
	assert.Equal(t, "MemberFriend", SnakeToBigCamel("_Member_friend"))
	assert.Equal(t, "MemberFriend", SnakeToBigCamel("_Member_Friend"))

	assert.Equal(t, "MemberFriend", SnakeToBigCamel("_member_friend_"))
	assert.Equal(t, "MemberFriend", SnakeToBigCamel("_member_Friend_"))
	assert.Equal(t, "MemberFriend", SnakeToBigCamel("_Member_friend_"))
	assert.Equal(t, "MemberFriend", SnakeToBigCamel("_Member_Friend_"))

	assert.Equal(t, "MemberFriend", SnakeToBigCamel("memberFriend"))
	assert.Equal(t, "MemberFriend", SnakeToBigCamel("MemberFriend"))
}
