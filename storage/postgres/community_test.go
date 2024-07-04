package postgres

import (
	pb "community/generated/community"
	"community/storage"
	"database/sql"
	"errors"
	"reflect"
	"testing"
)

func CommunnityRepo(t *testing.T) *NewCommunity {
	db, err := storage.Connect()
	if err != nil {
		return nil
	}
	com := NewCommunityRepo(db)
	return com
}

func TestGetGroup(t *testing.T) {
	com := CommunnityRepo(t)
	resp, err := com.GetGroup(&pb.GroupId{GroupId: "4437b581-81aa-4bf6-a59a-b36ee053215f"})
	if err != nil {
		panic(err)
	}
	case1 := pb.Group{
		GroupId:     "4437b581-81aa-4bf6-a59a-b36ee053215f",
		Name:        "Tech Enthusiasts",
		Description: "A group for people who love technology.",
		CreatedAt:   "2024-07-01T16:16:56.730211+05:00",
		CreatedBy:   "651bc75e-e040-4000-9c1d-8f77819148ef",
	}

	if !reflect.DeepEqual(resp, &case1) {
		t.Errorf("have %v, \nwant %v", resp, &case1)
	}

	_, err = com.GetGroup(&pb.GroupId{GroupId: "4437b581-81aa-4bf6-a59a-b36ee053215f"})
	if errors.Is(err, sql.ErrNoRows) {
		t.Log("SUCCESS")
	} else {
		t.Errorf("have %s, \nwant %s", err, sql.ErrNoRows)
	}
}

func TestUpdateGroup(t *testing.T) {
	com := CommunnityRepo(t)
	group := pb.Group{
		GroupId:     "",
		Name:        "",
		Description: "",
		CreatedBy:   "",
		CreatedAt:   "",
	}
	resp, err := com.UpdateGroup(&group)
	if err != nil {
		panic(err)
	}
	case1 := pb.Status{
		Status: true,
	}

	if !reflect.DeepEqual(resp, &case1) {
		t.Errorf("have %v, \nwant %v", resp, &case1)
	}

	_, err = com.UpdateGroup(&group)
	if errors.Is(err, sql.ErrNoRows) {
		t.Log("SUCCESS")
	} else {
		t.Errorf("have %s, \nwant %s", err, sql.ErrNoRows)
	}

}

func TestDeleteGroup(t *testing.T) {
	com := CommunnityRepo(t)

	resp, err := com.DeleteGroup(&pb.GroupId{GroupId: "4437b581-81aa-4bf6-a59a-b36ee053215f"})
	if err != nil {
		panic(err)
	}
	case1 := pb.Status{
		Status: true,
	}

	if !reflect.DeepEqual(resp, &case1) {
		t.Errorf("have %v, \nwant %v", resp, &case1)
	}

	_, err = com.DeleteGroup(&pb.GroupId{GroupId: ""})
	if errors.Is(err, sql.ErrNoRows) {
		t.Log("SUCCES")
	} else {
		t.Errorf("have %s, \nwant %s", err, sql.ErrNoRows)
	}

}

func TestGetAllGroups(t *testing.T) {
	com := CommunnityRepo(t)

	resp, err := com.GetAllGroups(&pb.Req{})
	if err != nil {
		panic(err)
	}

	case1 := pb.Status{}
}

func TestJoinGroupUser(t *testing.T) {
	com := CommunnityRepo(t)

	user := pb.JoinLeave{
		GroupId:  "",
		UserId:   "",
		Role:     "",
		JoinedAt: "",
	}

	resp, err := com.JoinGroupUser(&user)
	if err != nil {
		panic(err)
	}

	case1 := pb.Status{
		Status: true,
	}

	if !reflect.DeepEqual(resp, &case1) {
		t.Errorf("have %v, \nwant %v", resp, &case1)
	}

	_, err = com.JoinGroupUser(&user)
	if errors.Is(err, sql.ErrNoRows) {
		t.Log("SUCCESS")
	} else {
		t.Errorf("have %s, \nwant %s", err, sql.ErrNoRows)
	}

}

func TestLeaveGroupUser(t *testing.T) {
	com := CommunnityRepo(t)

	user := pb.JoinLeave{
		GroupId:  "",
		UserId:   "",
		Role:     "",
		JoinedAt: "",
	}

	resp, err := com.LeaveGroupUser(&user)
	if err != nil {
		panic(err)
	}

	case1 := pb.Status{
		Status: true,
	}

	if !reflect.DeepEqual(resp, &case1) {
		t.Errorf("have %v, \nwant %v", resp, &case1)
	}

	_, err = com.JoinGroupUser(&user)
	if errors.Is(err, sql.ErrNoRows) {
		t.Log("SUCCESS")
	} else {
		t.Errorf("have %s, \nwant %s", err, sql.ErrNoRows)
	}
}

func TestUpdateGroupMeber(t *testing.T) {
	com := CommunnityRepo(t)

	userRole := pb.UserRole{
		GroupId: "",
		UserId:  "",
		Role:    "",
	}

	resp, err := com.UpdateGroupMeber(&userRole)
	if err != nil {
		panic(err)
	}

	case1 := pb.Status{
		Status: true,
	}

	if !reflect.DeepEqual(resp, &case1) {
		t.Errorf("have %v, \nwant %v", resp, &case1)
	}

	_, err = com.UpdateGroupMeber(&userRole)
	if errors.Is(err, sql.ErrNoRows) {
		t.Log("SUCCESS")
	} else {
		t.Errorf("have %s, \nwant %s", err, sql.ErrNoRows)
	}
}
