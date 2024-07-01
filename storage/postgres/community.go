package postgres

import (
	pb "community/generated"
	"database/sql"
	"fmt"
	"time"
	"strconv"
	"strings"
	"github.com/google/uuid"
)

type NewCommunity struct {
	Db *sql.DB
}

func NewCommunityRepo(db *sql.DB) *NewCommunity {
	return &NewCommunity{Db: db}
}

func (C *NewCommunity) CreateGroup(group *pb.Group) (*pb.Status, error) {
	_, err := C.Db.Exec(`INSERT INTO
							Groups(name, description, created_by)
						VALUES($1, $2, $3)`,
		group.Name, group.Description, group.CreatedBy)
	if err != nil {
		return &pb.Status{Status: false}, err
	}
	return &pb.Status{Status: true}, nil
}


func (C *NewCommunity) GetGroup(groupId *pb.GroupId) (*pb.Group, error) {
	id, err := uuid.Parse(groupId.GroupId)
	if err != nil {
		return nil, fmt.Errorf("invalid group ID: %w", err)
	}

	group := pb.Group{}
	var userIdBytes []byte
	err = C.Db.QueryRow(`SELECT group_id, name, description, created_by, created_at
       FROM groups
       WHERE group_id = $1 AND deleted_at IS NULL`, id).Scan(&group.GroupId, group.Name, group.Description, group.CreatedBy, group.CreatedAt)
	if err != nil {
		return nil, err
	}

	group.GroupId = &pb.GroupId{GroupId: uuid.UUID(userIdBytes).String()}
	return &group, nil

}


func (C *NewCommunity) UpdateGroup(group *pb.Group) (*pb.Status, error) {
	groupId, err := uuid.Parse(group.GroupId.GroupId)
	if err != nil {
		return &pb.Status{Status: false}, err
	}
	arr := []interface{}{groupId}
	var param []string
	query := "UPDATE Groups SET group_id = $1"
	if len(group.Name) > 0 {
		arr = append(arr, group.Name)
		query += ", name = :name"
		param = append(param, ":name")
	}
	if len(group.Description) > 0 {
		arr = append(arr, group.Description)
		query += ", description = :description"
		param = append(param, ":description")
	}
	if len(group.CreatedBy) > 0 {
		arr = append(arr, group.CreatedBy)
		query += ", created_by = :created_by"
		param = append(param, ":created_by")
	}
	n := 1
	for _, j := range param {
		query = strings.Replace(query, j, "$"+strconv.Itoa(n), 1)
		n++
	}
	query += fmt.Sprintf(" WHERE group_id = $%d", n)
	arr = append(arr, groupId)

	_, err = C.Db.Exec(query, arr...)
	if err != nil {
		return &pb.Status{Status: false}, err
	}
	return &pb.Status{Status: true}, nil
}


func (C *NewCommunity) DeleteGroup(groupId *pb.GroupId) (*pb.Status, error) {
	id, err := uuid.Parse(groupId.GroupId)
	if err != nil {
		return nil, fmt.Errorf("invalid group ID: %w", err)
	}

	_, err = C.Db.Exec(`UPDATE groups
		SET deleted_at = $1
		Where group_id = $2`, time.Now(), id)
	if err != nil {
		return &pb.Status{Status: false}, err
	}

	return &pb.Status{Status: true}, nil

}

func(C *NewCommunity) GetAllGroups(req *pb.Req)(*pb.Groups, error){
	groupId := []byte{}
	groups := []*pb.Group{}
	rows, err := C.Db.Query(`SELECT * FROM Groups`)
	if err != nil{
		return &pb.Groups{Groups: groups}, err
	}
	for rows.Next(){
		var group = &pb.Group{}
		err = rows.Scan(&groupId, &group.Name, &group.Description, &group.CreatedBy, &group.CreatedAt)
		if err != nil{
			return &pb.Groups{Groups: groups}, err

		}
		group.GroupId = &pb.GroupId{GroupId: uuid.UUID(groupId).String()}
		
		groups = append(groups, group)
	}
	return &pb.Groups{Groups: groups}, nil
}


func (C *NewCommunity) JoinGroupUser(IDs *pb.JoinLeave) (*pb.Status, error) {

	rows, err := C.Db.Exec(`
    INSERT 
    INTO 
    group_members(
      group_id,
      user_id,
      role,
      joined_at
    )
    VALUES(
      $1,
      $2,
      $3,
      $4)
    `, IDs.GroupId, IDs.UserId, IDs.Role, IDs.JoinedAt)

  if err != nil {
    return &pb.Status{Status: false}, err
  }

  rowsAffected, err := rows.RowsAffected()

  if err != nil || rowsAffected == 0 {
    return &pb.Status{Status: false}, err
  }

  return &pb.Status{Status: true}, err

}


